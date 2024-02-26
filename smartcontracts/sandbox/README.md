# DePIN SandBox Contracts

## Design

### NodeRegistry

Register node and operator contract, one node can only register once and will get an NFT, the NFT tokenId is node id. One operator address can only bind to one node, and through operator address can query node info.

### FleetManager

The contract that manage node and project relationship.

### W3bstreamRouter

The router that route node message to project recevier contract.

## Deployment

### Testnet

```
PROJECT_REGISTRY: 0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26
NodeRegistry: 0x16ca331641a9537e346e12C7403fDA014Da72F16
FleetManager: 0x8D3c113805f970839940546D5ef88afE98Ba76E4
W3bstreamRouter: 0x1BFf17c79b5fa910cC77e95Ca82C7De26fC3C3b0

MockRisc0SnarkReceiver deployed to 0xC3e814db991924c64d94EfCb7a1ad6A479b1D728
```

## Getting started

1. Register project in ProjectRegistrar

```
./ioctl ws project create -u $CONFIG_FILE_PATH
```

2. Register node in NodeRegistry

```
export ETH_RPC_URL=https://babel-api.testnet.iotex.io
export PRIVATE_KEY=$PROJECT_OWNER_PRIVATE_KEY
cast send 0x16ca331641a9537e346e12C7403fDA014Da72F16 "register(address)" $ENODE_OPERATOR_ADDRESS --legacy --private-key=$PRIVATE_KEY
```

3. Allow enode for project

```
export ETH_RPC_URL=https://babel-api.testnet.iotex.io
export PRIVATE_KEY=$PROJECT_OWNER_PRIVATE_KEY
cast send 0x8D3c113805f970839940546D5ef88afE98Ba76E4 "allow(uint256,uint256)" $PROJECT_ID $NODE_ID --legacy --private-key=$PRIVATE_KEY
```

4. Register project receiver

```
export ETH_RPC_URL=https://babel-api.testnet.iotex.io
export PRIVATE_KEY=$PROJECT_OWNER_PRIVATE_KEY
cast send 0x1BFf17c79b5fa910cC77e95Ca82C7De26fC3C3b0 "register(uint256,address)" $PROJECT_ID $REVEIVER_ADDRESS --legacy --private-key=$PRIVATE_KEY
```
