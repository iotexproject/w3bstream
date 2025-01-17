### Prerequisites

https://github.com/foundry-rs/foundry
need the tools for interacting with smart contracts

### Testnet 

#### Endpoint
https://dragonfruit-testnet.w3bstream.com/v1/task

#### bind w3bstream project
```bash
cast send 0xB564996622CE5610b9cF4ed35160f406185d7d0b "register(uint256)" 942 --private-key "your private key" --rpc-url "https://babel-api.testnet.iotex.io" --legacy
```
```bash
cast send 0x7D3158166E9298fC47beA036fE162fEA17632E5D "updateConfig(uint256,string,bytes32)" 942 ipfs://ipfs.mainnet.iotex.io/QmUHfDnvWrr2wiC78dw85xfctzawNWAN1TEbzosxwHdzYC 0x8153291c230dd107f102f75e826a11d9d4a8ac3f0f4e1c3619e547f82a94410e --private-key "your private key" --rpc-url "https://babel-api.testnet.iotex.io" --legacy
```
```bash
cast send 0x7D3158166E9298fC47beA036fE162fEA17632E5D "resume(uint256)" 942 --private-key "your private key" --rpc-url "https://babel-api.testnet.iotex.io" --legacy
```

#### bind dapp
```bash
cast send 0x19dD7163Ad80fE550C97Affef49E1995B24941B1 "bindDapp(uint256,address)" 942 0xB2Dda5D9E65E44749409E209d8b7b15fb4e82147 --private-key "your private key" --rpc-url "https://babel-api.testnet.iotex.io" --legacy
```


### Mainnet 

#### Endpoint
https://dragonfruit-mainnet.w3bstream.com/v1/task

#### bind w3bstream project
```bash
cast send 0x425D3FD5e8e0d0d7c73599adeb9B395505581ec7 "register(uint256)" 9 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```
```bash
cast send 0x6EF4559f2023C93F78d27E0151deF083638478d2 "updateConfig(uint256,string,bytes32)" 9 ipfs://ipfs.mainnet.iotex.io/QmUHfDnvWrr2wiC78dw85xfctzawNWAN1TEbzosxwHdzYC 0x8153291c230dd107f102f75e826a11d9d4a8ac3f0f4e1c3619e547f82a94410e --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```
```bash
cast send 0x6EF4559f2023C93F78d27E0151deF083638478d2 "resume(uint256)" 9 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```

#### bind dapp
```bash
cast send 0x580D9686A7A188746B9f4a06fb5ec9e14E937fde "bindDapp(uint256,address)" 9 0xde44BEd8c143B75deDca6A065Fdabb8AbE95ECC6 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```