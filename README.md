# W3bstream Sprout
W3bstream Sprout (Alpha) which supports native Halo2 circuits (WIP) as well as zkVMs like zkWASM (WIP) and RISC0.

wsctl is the command line tool that interact with W3bstream protocol which can be used by node operators as well as project developers.


### Minimum requirements

| Components | Version | Description |
|----------|-------------|-------------|
| [Golang](https://golang.org) | &ge; 1.21 | Go programming language |

### Step 1
install wsctl
`curl https://raw.githubusercontent.com/machinefi/sprout/main/scripts/install_wsctl.sh | bash`

### Step 2
install docker & docker-compose

### Step 3
```bash
mkdir sprout && cd sprout
curl https://raw.githubusercontent.com/machinefi/sprout/main/docker-compose.yaml -o docker-compose.yaml
```

### Step 4
```bash
wsctl node up --private-key "your private key"
```

### Step 5
```bash
wsctl node log
```

### Step 6 
open a new terminal and execute
```bash
wsctl message send -p "test01" -v "0.1" -d "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Stark\"}"
```
will see log on `wsctl node log` terminal

### Step 7
```bash
wsctl node down
```

### Step 8
can watch config by 
```bash
wsctl config get
```
and set config by 
```bash
wsctl config set
```
