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
cast send 0x97c3696E5f9A17569711B002152fd1603f8F06eB "register(uint256)" 9 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```

```bash
cast send 0xee8e318C712aB1731f9c3b708a5Caf2533614AF3 "updateConfig(uint256,string,bytes32)" 9 ipfs://ipfs.mainnet.iotex.io/QmPmnceezQsgWQRwR9seYLQ666rEkfxi4LgLCiLJeBqMpA 0xba270fc9a9a0817e1086ce2ecfd9c951b644a1aa628beb38b18e734c68a7e1f0 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```

```bash
cast send 0xee8e318C712aB1731f9c3b708a5Caf2533614AF3 "resume(uint256)" 9 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```

#### bind dapp

```bash
cast send 0xeBf9Ab649f9952F9B6e85e59Fac9fED43594e3E0 "bindDapp(uint256,address)" 9 0x5197646593d536546F583009F0057C6b48aBAa6E --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```
