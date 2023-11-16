# W3bstream Sprout :four_leaf_clover:
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
### Download docker-compose.yaml
```bash
mkdir sprout && cd sprout
curl https://raw.githubusercontent.com/machinefi/sprout/main/docker-compose.yaml -o docker-compose.yaml
```

### Populate docker-compose.yaml fields
W3bstream-node need write proof to chain and need a private key for chain write. Set your private key at https://github.com/machinefi/sprout/blob/main/docker-compose.yaml#L20  
If you need risc0 snark proof, a bonsai key is needed at https://github.com/machinefi/sprout/blob/main/docker-compose.yaml#L40

### Use customized project code
Docker-compose will mount current work directory to containers /data https://github.com/machinefi/sprout/blob/main/docker-compose.yaml#L23  
So you can appoint the project file at https://github.com/machinefi/sprout/blob/main/docker-compose.yaml#L18


### Start w3bstream node
```bash
docker-compose up -d
```

### Monitor w3bstream node status
```bash
docker-compose logs -f w3bnode
```

### Shut down w3bstream node
```bash
docker-compose down
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
risc0 snark test
```bash
wsctl message send -p 10000 -v "0.1" -d "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```
halo2 test 
```bash
./wsctl message send -p 10001 -v "0.1" -d "{\"private_input\":\"4\"}"
```
It will send a message to project test01 running on the remote server. The processing status could be checked via `docker-compose logs -f w3bnode` on the server.
