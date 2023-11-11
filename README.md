# W3bstream Sprout
W3bstream Sprout (Alpha) which supports native Halo2 circuits (WIP) as well as zkVMs like zkWASM (WIP) and RISC0.

`wsctl`` is the command line tool that interact with W3bstream protocol which can be used by node operators as well as project developers.

## Minimum requirements

| Components | Version | Description |
|----------|-------------|-------------|
| [Golang](https://golang.org) | &ge; 1.21 | Go programming language |

## Preparation

### Install docker
Docker is needed to run the node service. Please make sure your docker is up to date.
```bash
install docker & docker-compose
```
### Install wsctl
`wsctl` is the W3bstream command line.
```bash
curl https://raw.githubusercontent.com/machinefi/sprout/main/scripts/install_wsctl.sh | bash
```

## Run server
### Download source code
```bash
mkdir sprout && cd sprout
curl https://raw.githubusercontent.com/machinefi/sprout/main/docker-compose.yaml -o docker-compose.yaml
```

### Start w3bstream node
```bash
wsctl node up --private-key "your private key"
```

### Monitor w3bstream node status
```bash
wsctl node log
```

### Shut down w3bstream node
```bash
wsctl node down
```

### Compile your own w3bstream code
After modifying w3bstream source code, the image could be rebuild by running
```bash
make docker
```
Replace the image name in the docker-compose.yaml to the new image built above. 

## Send message from client
### Set remote w3bstream server
Set server endpoint and language via subcommand `config`:
```bash
wsctl config set endpoint localhost:8888
```
The value of the variables in config could be fetched via
```bash
wsctl config get endpoint
```

### Send message to remote server
open a new terminal and execute
```bash
wsctl message send -p "test01" -v "0.1" -d "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```
It will send a message to project test01 running on the remote server. The processing status could be checked via `wsctl node log` on the server.
