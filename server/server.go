package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/dasio/pcfg-manager/manager"
	pb "github.com/dasio/pcfg-manager/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"os"
	"sync/atomic"
	"time"
)

const (
	chunkThreshold = uint64(100000)
)

type Service struct {
	port            string
	mng             *manager.Manager
	remainingHashes map[string]struct{}
	completedHashes map[string]string
	ch              <-chan *manager.TreeItem
	clients         map[string]ClientInfo
	chunkId         uint32
	endCracking     chan bool
}

type Chunk struct {
	Id             uint32
	Items          []*pb.TreeItem
	TerminalsCount uint64
}

func NewService() *Service {
	return &Service{
		port:    ":50051",
		clients: make(map[string]ClientInfo),
		chunkId: 0,
	}
}

type ClientInfo struct {
	Addr        string
	ActualChunk Chunk
}

func (s *Service) Load(ruleName, hashFile string) error {
	lines, err := readLines(hashFile)
	if err != nil {
		return err
	}
	s.remainingHashes = make(map[string]struct{})
	s.completedHashes = make(map[string]string)
	s.endCracking = make(chan bool)
	for _, l := range lines {
		s.remainingHashes[l] = struct{}{}
	}
	fmt.Println(s.remainingHashes)
	s.mng = manager.NewManager(ruleName)
	if err := s.mng.Load(); err != nil {
		return err
	}
	s.ch = s.mng.Generator.RunForServer(&manager.InputArgs{})
	return nil
}
func (s *Service) Run() error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterPCFGServer(server, s)
	logrus.Infof("Listening on port %s", s.port)
	go func() {
		<-s.endCracking
		server.GracefulStop()
	}()
	if err := server.Serve(lis); err != nil {
		return err
	}
	for hash, pass := range s.completedHashes {
		fmt.Println(hash, " ", pass)
	}
	return nil
}

func (s *Service) Connect(ctx context.Context, req *pb.Empty) (*pb.ConnectResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &pb.ConnectResponse{}, errors.New("no peer")
	}
	client := ClientInfo{
		Addr: p.Addr.String(),
	}
	s.clients[client.Addr] = client
	var hashList []string
	for k := range s.remainingHashes {
		hashList = append(hashList, k)
	}
	logrus.Infof("client %s connected", client.Addr)
	return &pb.ConnectResponse{
		Grammar:  pb.GrammarToProto(s.mng.Generator.Pcfg.Grammar),
		HashList: hashList,
	}, nil
}

func (s *Service) Disconnect(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &pb.Empty{}, errors.New("no peer")
	}
	delete(s.clients, p.Addr.String())
	logrus.Infof("client %s disconnected", p.Addr.String())

	return &pb.Empty{}, nil
}

func (s *Service) GetNextChunk() (Chunk, bool) {
	total := uint64(0)
	endGen := false
	var chunkItems []*pb.TreeItem
loop:
	for total < chunkThreshold {
		select {
		case it := <-s.ch:
			if it == nil {
				endGen = true
				break loop
			}
			chunkItems = append(chunkItems, pb.TreeItemToProto(it))
			guessGeneration := manager.NewGuessGeneration(s.mng.Generator.Pcfg.Grammar, it)
			total += guessGeneration.Count()
		case <-time.After(time.Second * 2):
			break loop
		}
	}
	return Chunk{
		Id:             atomic.AddUint32(&s.chunkId, 1),
		Items:          chunkItems,
		TerminalsCount: total,
	}, endGen
}
func (s *Service) GetNextItems(ctx context.Context, req *pb.Empty) (*pb.TreeItems, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &pb.TreeItems{}, errors.New("no peer")
	}
	chunk, endGen := s.GetNextChunk()
	if endGen && len(chunk.Items) == 0 {
		return &pb.TreeItems{}, manager.ErrPriorirtyQueEmpty
	}
	clientInfo := s.clients[p.Addr.String()]
	clientInfo.ActualChunk = chunk
	s.clients[p.Addr.String()] = clientInfo
	logrus.Infof("sending chunk[%d], preTerminals: %d, terminals: %d to %s", chunk.Id, len(chunk.Items), chunk.TerminalsCount, clientInfo.Addr)
	return &pb.TreeItems{
		Items: chunk.Items,
	}, nil
}

func (s *Service) SendResult(ctx context.Context, in *pb.CrackingResponse) (*pb.Empty, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &pb.Empty{}, errors.New("no peer")
	}
	clientInfo := s.clients[p.Addr.String()]
	for hash, password := range in.Hashes {
		delete(s.remainingHashes, hash)
		s.completedHashes[hash] = password
	}
	logrus.Infof("result from %s: %d", clientInfo.Addr, len(in.Hashes))
	if len(s.remainingHashes) == 0 {
		s.endCracking <- true
	}
	return &pb.Empty{}, nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
