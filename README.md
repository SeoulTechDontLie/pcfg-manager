# Parallel/distributed PCFG Manager
**A high-performance password geuess generator**
David Mikus, Radek Hranicky, Filip Listiak, Lukas Zobal
This is a proof-of-concept tool created under NES@FIT research group in cooperation with the Fitcrack team: https://fitcrack.fit.vutbr.cz/team/
(Currently used for distributed PCFG attacks in the Fitcrack system)

# The tool can run:
* As standalone high-performance password guess generator
* As distributed password guessing server
* As distributed password guessing client (support for direct cracking with hashcat)

# How to use this?
Get Matt Weir's PCFG trainer: https://github.com/lakiw/legacy-pcfg/tree/master/python_pcfg_cracker_version3/pcfg_trainer
Generate a PCFG from a training dictionary. Have converage set to 1.0. For example:
```
./pcfg_trainer.py --training test.txt --rule test --coverage 1.0
```

Use the generated grammar as the input of this tool, for example:
```
./pcfg-manager -r Rules/test
```


# Command line parameters:
```
Usage:
  pcfg-manager [flags]
  pcfg-manager [command]

Available Commands:
  client      run client
  help        Help about any command
  marshal     used for marshalling grammar
  server      run server

Flags:
  -d, --debug
  -g, --go-routines uint           how many go routines will be used (default 1)
      --grammar-file string        it uses marshaled grammar file instead of parsing
  -h, --help                       help for pcfg-manager
  -m, --max-guesses uint           max guesses before exit (generates at least m terminals, could be more)
      --preterminals-file string   generates terminals from specified preterminals
  -r, --rules string               specifies rule directory of the PCFG (default "Rules/Default")

Use "pcfg-manager [command] --help" for more information about a command.
```

# Example 1: Standalone parallel guessing:
```
./pcfg-manager -r Rules/twitter-banned --go-routines 4
INFO[0000] Config file loaded
monkey
asdfgh
access
11111111
12345678
password
...
```

# Example 2: Distributed guessing:
```
./pcfg-manager server -r Rules/facebook-pastebay -p 12345 &
INFO[0000] Config file loaded
INFO[0000] Listening on port 12345

./pcfg-manager client -s localhost:12345 --generate-only
INFO[0011] client [::1]:35566 connected
INFO[0011] sending chunk[1], preTerminals: 46, terminals: 999 to [::1]:35566 in 51.836µs, size: 883
INFO[0000] received 46 preTerminals and 0 terminals
aaparken
123456789
didierdemaeyer
22042009
gemigencola
2251084185
2406923269
larahornby
moravanska
...
```

# Example 3: Distributed cracking with hashcat:
```
./pcfg-manager server -r Rules/test -p 12345 --hashcat-mode 100 --hashlist my-sha1-hashlist.hash &
INFO[0000] Config file loaded
INFO[0000] Listening on port 12345

./pcfg-manager client -s localhost:12345 --hashcat-folder hashcat --stats
INFO[0050] client [::1]:35440 connected
INFO[0050] sending chunk[1], preTerminals: 1251, terminals: 10026 to [::1]:35440 in 1.740827ms, size: 63603
INFO[0000] received 1251 preTerminals and 0 terminals
hashcat (v5.1.0) starting...

Hashes: 1 digests; 1 unique digests, 1 unique salts
Bitmaps: 16 bits, 65536 entries, 0x0000ffff mask, 262144 bytes, 5/13 rotates
...
```

# The work was presented in:
HRANICKÝ Radek, LIŠTIAK Filip, MIKUŠ Dávid and RYŠAVÝ Ondrej.
On Practical Aspects of PCFG Password Cracking.
In: Data and Applications Security and Privacy. Charleston: Springer Nature Switzerland AG, 2019,
pp. 43-60. ISBN 978-3-030-22478-3. ISSN 0302-9743.
Available from: https://link.springer.com/chapter/10.1007%2F978-3-030-22479-0_3

HRANICKÝ Radek, ZOBAL Lukáš, RYŠAVÝ Ondrej, KOLÁR Dušan and MIKUŠ Dávid.
Distributed PCFG Password Cracking.
In: Computer Security - ESORICS 2020. Lecture notes in Computer Science.
Guildford: Springer Nature Switzerland AG, 2020, pp. 701-719. ISBN 978-3-030-58950-9.
Available from: https://link.springer.com/chapter/10.1007/978-3-030-58951-6_34
