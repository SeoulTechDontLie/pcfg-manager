// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ConnectResponse struct {
	Grammar              *Grammar `protobuf:"bytes,1,opt,name=grammar,proto3" json:"grammar,omitempty"`
	HashList             []string `protobuf:"bytes,2,rep,name=hashList,proto3" json:"hashList,omitempty"`
	HashcatMode          string   `protobuf:"bytes,3,opt,name=hashcatMode,proto3" json:"hashcatMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

func (m *ConnectResponse) GetGrammar() *Grammar {
	if m != nil {
		return m.Grammar
	}
	return nil
}

func (m *ConnectResponse) GetHashList() []string {
	if m != nil {
		return m.HashList
	}
	return nil
}

func (m *ConnectResponse) GetHashcatMode() string {
	if m != nil {
		return m.HashcatMode
	}
	return ""
}

type ResultResponse struct {
	End                  bool     `protobuf:"varint,1,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultResponse) Reset()         { *m = ResultResponse{} }
func (m *ResultResponse) String() string { return proto.CompactTextString(m) }
func (*ResultResponse) ProtoMessage()    {}
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{2}
}

func (m *ResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultResponse.Unmarshal(m, b)
}
func (m *ResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultResponse.Marshal(b, m, deterministic)
}
func (m *ResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultResponse.Merge(m, src)
}
func (m *ResultResponse) XXX_Size() int {
	return xxx_messageInfo_ResultResponse.Size(m)
}
func (m *ResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResultResponse proto.InternalMessageInfo

func (m *ResultResponse) GetEnd() bool {
	if m != nil {
		return m.End
	}
	return false
}

type CrackingResponse struct {
	Hashes               map[string]string `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CrackingResponse) Reset()         { *m = CrackingResponse{} }
func (m *CrackingResponse) String() string { return proto.CompactTextString(m) }
func (*CrackingResponse) ProtoMessage()    {}
func (*CrackingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{3}
}

func (m *CrackingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrackingResponse.Unmarshal(m, b)
}
func (m *CrackingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrackingResponse.Marshal(b, m, deterministic)
}
func (m *CrackingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrackingResponse.Merge(m, src)
}
func (m *CrackingResponse) XXX_Size() int {
	return xxx_messageInfo_CrackingResponse.Size(m)
}
func (m *CrackingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CrackingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CrackingResponse proto.InternalMessageInfo

func (m *CrackingResponse) GetHashes() map[string]string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type Grammar struct {
	RuleName             string             `protobuf:"bytes,1,opt,name=ruleName,proto3" json:"ruleName,omitempty"`
	Sections             []*Section         `protobuf:"bytes,2,rep,name=sections,proto3" json:"sections,omitempty"`
	Mapping              map[string]*IntMap `protobuf:"bytes,3,rep,name=mapping,proto3" json:"mapping,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Grammar) Reset()         { *m = Grammar{} }
func (m *Grammar) String() string { return proto.CompactTextString(m) }
func (*Grammar) ProtoMessage()    {}
func (*Grammar) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{4}
}

func (m *Grammar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grammar.Unmarshal(m, b)
}
func (m *Grammar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grammar.Marshal(b, m, deterministic)
}
func (m *Grammar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grammar.Merge(m, src)
}
func (m *Grammar) XXX_Size() int {
	return xxx_messageInfo_Grammar.Size(m)
}
func (m *Grammar) XXX_DiscardUnknown() {
	xxx_messageInfo_Grammar.DiscardUnknown(m)
}

var xxx_messageInfo_Grammar proto.InternalMessageInfo

func (m *Grammar) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

func (m *Grammar) GetSections() []*Section {
	if m != nil {
		return m.Sections
	}
	return nil
}

func (m *Grammar) GetMapping() map[string]*IntMap {
	if m != nil {
		return m.Mapping
	}
	return nil
}

type IntMap struct {
	Value                map[string]int32 `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{5}
}

func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (m *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(m, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetValue() map[string]int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Replacement struct {
	Probability          float64  `protobuf:"fixed64,1,opt,name=probability,proto3" json:"probability,omitempty"`
	IsTerminal           bool     `protobuf:"varint,2,opt,name=isTerminal,proto3" json:"isTerminal,omitempty"`
	Values               []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	Function             string   `protobuf:"bytes,4,opt,name=function,proto3" json:"function,omitempty"`
	Pos                  []int32  `protobuf:"varint,5,rep,packed,name=pos,proto3" json:"pos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Replacement) Reset()         { *m = Replacement{} }
func (m *Replacement) String() string { return proto.CompactTextString(m) }
func (*Replacement) ProtoMessage()    {}
func (*Replacement) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{6}
}

func (m *Replacement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Replacement.Unmarshal(m, b)
}
func (m *Replacement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Replacement.Marshal(b, m, deterministic)
}
func (m *Replacement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Replacement.Merge(m, src)
}
func (m *Replacement) XXX_Size() int {
	return xxx_messageInfo_Replacement.Size(m)
}
func (m *Replacement) XXX_DiscardUnknown() {
	xxx_messageInfo_Replacement.DiscardUnknown(m)
}

var xxx_messageInfo_Replacement proto.InternalMessageInfo

func (m *Replacement) GetProbability() float64 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func (m *Replacement) GetIsTerminal() bool {
	if m != nil {
		return m.IsTerminal
	}
	return false
}

func (m *Replacement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Replacement) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Replacement) GetPos() []int32 {
	if m != nil {
		return m.Pos
	}
	return nil
}

type Section struct {
	Type                 string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Replacements         []*Replacement `protobuf:"bytes,3,rep,name=replacements,proto3" json:"replacements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Section) Reset()         { *m = Section{} }
func (m *Section) String() string { return proto.CompactTextString(m) }
func (*Section) ProtoMessage()    {}
func (*Section) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{7}
}

func (m *Section) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Section.Unmarshal(m, b)
}
func (m *Section) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Section.Marshal(b, m, deterministic)
}
func (m *Section) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Section.Merge(m, src)
}
func (m *Section) XXX_Size() int {
	return xxx_messageInfo_Section.Size(m)
}
func (m *Section) XXX_DiscardUnknown() {
	xxx_messageInfo_Section.DiscardUnknown(m)
}

var xxx_messageInfo_Section proto.InternalMessageInfo

func (m *Section) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Section) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Section) GetReplacements() []*Replacement {
	if m != nil {
		return m.Replacements
	}
	return nil
}

type Items struct {
	PreTerminals         []*TreeItem `protobuf:"bytes,1,rep,name=preTerminals,proto3" json:"preTerminals,omitempty"`
	Terminals            []string    `protobuf:"bytes,2,rep,name=terminals,proto3" json:"terminals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Items) Reset()         { *m = Items{} }
func (m *Items) String() string { return proto.CompactTextString(m) }
func (*Items) ProtoMessage()    {}
func (*Items) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{8}
}

func (m *Items) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Items.Unmarshal(m, b)
}
func (m *Items) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Items.Marshal(b, m, deterministic)
}
func (m *Items) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Items.Merge(m, src)
}
func (m *Items) XXX_Size() int {
	return xxx_messageInfo_Items.Size(m)
}
func (m *Items) XXX_DiscardUnknown() {
	xxx_messageInfo_Items.DiscardUnknown(m)
}

var xxx_messageInfo_Items proto.InternalMessageInfo

func (m *Items) GetPreTerminals() []*TreeItem {
	if m != nil {
		return m.PreTerminals
	}
	return nil
}

func (m *Items) GetTerminals() []string {
	if m != nil {
		return m.Terminals
	}
	return nil
}

type TreeItem struct {
	Index                int32       `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Transition           int32       `protobuf:"varint,2,opt,name=transition,proto3" json:"transition,omitempty"`
	Childrens            []*TreeItem `protobuf:"bytes,3,rep,name=childrens,proto3" json:"childrens,omitempty"`
	Id                   bool        `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TreeItem) Reset()         { *m = TreeItem{} }
func (m *TreeItem) String() string { return proto.CompactTextString(m) }
func (*TreeItem) ProtoMessage()    {}
func (*TreeItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{9}
}

func (m *TreeItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeItem.Unmarshal(m, b)
}
func (m *TreeItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeItem.Marshal(b, m, deterministic)
}
func (m *TreeItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeItem.Merge(m, src)
}
func (m *TreeItem) XXX_Size() int {
	return xxx_messageInfo_TreeItem.Size(m)
}
func (m *TreeItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeItem.DiscardUnknown(m)
}

var xxx_messageInfo_TreeItem proto.InternalMessageInfo

func (m *TreeItem) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TreeItem) GetTransition() int32 {
	if m != nil {
		return m.Transition
	}
	return 0
}

func (m *TreeItem) GetChildrens() []*TreeItem {
	if m != nil {
		return m.Childrens
	}
	return nil
}

func (m *TreeItem) GetId() bool {
	if m != nil {
		return m.Id
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*ConnectResponse)(nil), "proto.ConnectResponse")
	proto.RegisterType((*ResultResponse)(nil), "proto.ResultResponse")
	proto.RegisterType((*CrackingResponse)(nil), "proto.CrackingResponse")
	proto.RegisterMapType((map[string]string)(nil), "proto.CrackingResponse.HashesEntry")
	proto.RegisterType((*Grammar)(nil), "proto.Grammar")
	proto.RegisterMapType((map[string]*IntMap)(nil), "proto.Grammar.MappingEntry")
	proto.RegisterType((*IntMap)(nil), "proto.IntMap")
	proto.RegisterMapType((map[string]int32)(nil), "proto.IntMap.ValueEntry")
	proto.RegisterType((*Replacement)(nil), "proto.Replacement")
	proto.RegisterType((*Section)(nil), "proto.Section")
	proto.RegisterType((*Items)(nil), "proto.PreTerminals")
	proto.RegisterType((*TreeItem)(nil), "proto.TreeItem")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x93, 0x3a, 0x1f, 0xe3, 0xd0, 0x56, 0x2b, 0x28, 0x56, 0x40, 0x28, 0x72, 0x2f, 0x16,
	0x82, 0x48, 0xa4, 0x02, 0x95, 0x8f, 0x5b, 0x29, 0xa1, 0x12, 0xad, 0xd0, 0xb6, 0xe2, 0xc0, 0x6d,
	0x6b, 0x0f, 0xed, 0xaa, 0xf6, 0xda, 0xda, 0xdd, 0xa0, 0xe6, 0xc4, 0x9d, 0x3b, 0xbf, 0x8b, 0x03,
	0x7f, 0x08, 0xed, 0xae, 0x1d, 0x3b, 0x50, 0xc4, 0x25, 0x99, 0x79, 0xf3, 0x66, 0x76, 0xe7, 0xcd,
	0x7a, 0x20, 0x28, 0x65, 0xa1, 0x8b, 0xa9, 0xfd, 0x25, 0xbe, 0xfd, 0x8b, 0xfa, 0xe0, 0x1f, 0xe5,
	0xa5, 0x5e, 0x46, 0x4b, 0xd8, 0x3e, 0x2c, 0x84, 0xc0, 0x44, 0x53, 0x54, 0x65, 0x21, 0x14, 0x92,
	0x18, 0xfa, 0x97, 0x92, 0xe5, 0x39, 0x93, 0xa1, 0x37, 0xf1, 0xe2, 0x60, 0xb6, 0xe5, 0x72, 0xa7,
	0x73, 0x87, 0xd2, 0x3a, 0x4c, 0xc6, 0x30, 0xb8, 0x62, 0xea, 0xea, 0x03, 0x57, 0x3a, 0xec, 0x4c,
	0xba, 0xf1, 0x90, 0xae, 0x7c, 0x32, 0x81, 0xc0, 0xd8, 0x09, 0xd3, 0x27, 0x45, 0x8a, 0x61, 0x77,
	0xe2, 0xc5, 0x43, 0xda, 0x86, 0xa2, 0x08, 0xb6, 0x28, 0xaa, 0x45, 0xd6, 0x9c, 0xbc, 0x03, 0x5d,
	0x14, 0xa9, 0x3d, 0x75, 0x40, 0x8d, 0x19, 0x7d, 0xf7, 0x60, 0xe7, 0x50, 0xb2, 0xe4, 0x9a, 0x8b,
	0xcb, 0x15, 0xed, 0x35, 0xf4, 0x4c, 0x1d, 0x54, 0xa1, 0x37, 0xe9, 0xc6, 0xc1, 0x6c, 0xaf, 0xba,
	0xdf, 0x9f, 0xc4, 0xe9, 0x7b, 0xcb, 0x3a, 0x12, 0x5a, 0x2e, 0x69, 0x95, 0x32, 0x7e, 0x09, 0x41,
	0x0b, 0x36, 0x47, 0x5e, 0xe3, 0xd2, 0x1e, 0x39, 0xa4, 0xc6, 0x24, 0x77, 0xc1, 0xff, 0xca, 0xb2,
	0x05, 0x86, 0x1d, 0x8b, 0x39, 0xe7, 0x55, 0xe7, 0xc0, 0x8b, 0x7e, 0x79, 0xd0, 0x9f, 0x37, 0xad,
	0xcb, 0x45, 0x86, 0xa7, 0x2c, 0xc7, 0x2a, 0x79, 0xe5, 0x93, 0xc7, 0x30, 0x50, 0x98, 0x68, 0x5e,
	0x08, 0x65, 0x65, 0x69, 0x14, 0x3c, 0x73, 0x30, 0x5d, 0xc5, 0xc9, 0x73, 0xe8, 0xe7, 0xac, 0x2c,
	0xb9, 0xb8, 0x0c, 0xbb, 0x96, 0xfa, 0x60, 0x5d, 0xec, 0xe9, 0x89, 0x8b, 0xba, 0x26, 0x6a, 0xee,
	0xf8, 0x18, 0x46, 0xed, 0xc0, 0x2d, 0x6d, 0xec, 0xb5, 0xdb, 0x08, 0x66, 0x77, 0xaa, 0xb2, 0xc7,
	0x42, 0x9f, 0xb0, 0xb2, 0xdd, 0x95, 0x84, 0x9e, 0x03, 0xc9, 0xb4, 0x4e, 0x71, 0xb2, 0x86, 0x6b,
	0x29, 0xd3, 0x4f, 0x26, 0xe4, 0xae, 0xe1, 0x68, 0xe3, 0x03, 0x80, 0x06, 0xfc, 0x9f, 0x92, 0x7e,
	0xfb, 0xcc, 0x1f, 0x1e, 0x04, 0x14, 0xcb, 0x8c, 0x25, 0x98, 0xa3, 0xb0, 0x8f, 0xa5, 0x94, 0xc5,
	0x05, 0xbb, 0xe0, 0x19, 0xd7, 0xae, 0x86, 0x47, 0xdb, 0x10, 0x79, 0x04, 0xc0, 0xd5, 0x39, 0xca,
	0x9c, 0x0b, 0x96, 0xd9, 0x82, 0x03, 0xda, 0x42, 0xc8, 0x2e, 0xf4, 0x6c, 0x79, 0x65, 0x65, 0x1c,
	0xd2, 0xca, 0x33, 0x73, 0xfa, 0xb2, 0x10, 0x56, 0xec, 0x70, 0xd3, 0xcd, 0xa9, 0xf6, 0xcd, 0x8d,
	0xcb, 0x42, 0x85, 0xfe, 0xa4, 0x1b, 0xfb, 0xd4, 0x98, 0x11, 0x87, 0x7e, 0x35, 0x22, 0x42, 0x60,
	0x53, 0x2f, 0xcb, 0x7a, 0xb8, 0xd6, 0x36, 0x98, 0x30, 0x03, 0x77, 0x2f, 0xc3, 0xda, 0xe4, 0x05,
	0x8c, 0x64, 0xd3, 0x89, 0xaa, 0xa6, 0x48, 0x2a, 0xed, 0x5a, 0x4d, 0xd2, 0x35, 0x5e, 0xf4, 0x19,
	0xfc, 0x63, 0x8d, 0xb9, 0x22, 0xfb, 0x30, 0x2a, 0x25, 0xd6, 0x8d, 0xd4, 0x6f, 0x7a, 0xbb, 0x2a,
	0x70, 0x2e, 0x11, 0x0d, 0x8f, 0xae, 0x91, 0xc8, 0x43, 0x18, 0xea, 0x55, 0x86, 0xfb, 0xf4, 0x1a,
	0x20, 0xfa, 0x06, 0x83, 0x3a, 0xcf, 0x0c, 0x81, 0x8b, 0x14, 0x6f, 0x6c, 0x23, 0x3e, 0x75, 0x8e,
	0x91, 0x53, 0x4b, 0x26, 0x14, 0xb7, 0xc2, 0xb8, 0xf9, 0xb4, 0x10, 0xf2, 0x14, 0x86, 0xc9, 0x15,
	0xcf, 0x52, 0x89, 0xa2, 0x6e, 0xe9, 0xaf, 0x1b, 0x35, 0x0c, 0xb2, 0x05, 0x1d, 0x9e, 0x5a, 0x7d,
	0x07, 0xb4, 0xc3, 0xd3, 0xd9, 0x4f, 0x0f, 0x36, 0x3f, 0x1e, 0xbe, 0x9b, 0x93, 0x67, 0xd0, 0xaf,
	0xd6, 0x0b, 0x19, 0x55, 0xf9, 0x76, 0xef, 0x8c, 0x77, 0xeb, 0x6f, 0x76, 0x7d, 0xf9, 0x44, 0x1b,
	0x24, 0x06, 0x78, 0xcb, 0x55, 0x72, 0x6b, 0xd6, 0x9a, 0x47, 0x9e, 0xc0, 0x68, 0x8e, 0xfa, 0x14,
	0x6f, 0xb4, 0x53, 0xf2, 0x76, 0xae, 0x8d, 0x45, 0x1b, 0xe4, 0x0d, 0xc0, 0x19, 0x8a, 0xd4, 0xad,
	0x1c, 0x72, 0xff, 0x1f, 0x3b, 0x63, 0x7c, 0x6f, 0x35, 0xb9, 0xf6, 0x6a, 0xba, 0xe8, 0x59, 0x74,
	0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xeb, 0x72, 0x52, 0x4d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PCFGClient is the client API for PCFG service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PCFGClient interface {
	Connect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConnectResponse, error)
	Disconnect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetNextItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Items, error)
	SendResult(ctx context.Context, in *CrackingResponse, opts ...grpc.CallOption) (*ResultResponse, error)
}

type pCFGClient struct {
	cc *grpc.ClientConn
}

func NewPCFGClient(cc *grpc.ClientConn) PCFGClient {
	return &pCFGClient{cc}
}

func (c *pCFGClient) Connect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/proto.PCFG/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCFGClient) Disconnect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.PCFG/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCFGClient) GetNextItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Items, error) {
	out := new(Items)
	err := c.cc.Invoke(ctx, "/proto.PCFG/GetNextItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pCFGClient) SendResult(ctx context.Context, in *CrackingResponse, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/proto.PCFG/SendResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PCFGServer is the server API for PCFG service.
type PCFGServer interface {
	Connect(context.Context, *Empty) (*ConnectResponse, error)
	Disconnect(context.Context, *Empty) (*Empty, error)
	GetNextItems(context.Context, *Empty) (*Items, error)
	SendResult(context.Context, *CrackingResponse) (*ResultResponse, error)
}

// UnimplementedPCFGServer can be embedded to have forward compatible implementations.
type UnimplementedPCFGServer struct {
}

func (*UnimplementedPCFGServer) Connect(ctx context.Context, req *Empty) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedPCFGServer) Disconnect(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (*UnimplementedPCFGServer) GetNextItems(ctx context.Context, req *Empty) (*Items, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextItems not implemented")
}
func (*UnimplementedPCFGServer) SendResult(ctx context.Context, req *CrackingResponse) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResult not implemented")
}

func RegisterPCFGServer(s *grpc.Server, srv PCFGServer) {
	s.RegisterService(&_PCFG_serviceDesc, srv)
}

func _PCFG_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFGServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PCFG/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFGServer).Connect(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCFG_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFGServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PCFG/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFGServer).Disconnect(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCFG_GetNextItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFGServer).GetNextItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PCFG/GetNextItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFGServer).GetNextItems(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PCFG_SendResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrackingResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PCFGServer).SendResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PCFG/SendResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PCFGServer).SendResult(ctx, req.(*CrackingResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _PCFG_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PCFG",
	HandlerType: (*PCFGServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _PCFG_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _PCFG_Disconnect_Handler,
		},
		{
			MethodName: "GetNextItems",
			Handler:    _PCFG_GetNextItems_Handler,
		},
		{
			MethodName: "SendResult",
			Handler:    _PCFG_SendResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
