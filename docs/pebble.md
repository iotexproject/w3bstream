### Prerequisites

https://github.com/foundry-rs/foundry
need the tools for interacting with smart contracts

### Mainnet 

#### Endpoint
https://dragonfruit-mainnet.w3bstream.com/v1/task

#### bind w3bstream project
```bash
cast send 0x425D3FD5e8e0d0d7c73599adeb9B395505581ec7 "register(uint256)" 6 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```
```bash
cast send 0x6EF4559f2023C93F78d27E0151deF083638478d2 "updateConfig(uint256,string,bytes32)" 6 ipfs://ipfs.mainnet.iotex.io/QmQfaXAr7ZHbn1inDKigx9GcaHkZqxJiiHsA81GfCTRYZ6 0x9273b57144d4df4c6fb2a5850a1a95891c1b12f92d06909d875c3a4afe18bae4 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```
```bash
cast send 0x6EF4559f2023C93F78d27E0151deF083638478d2 "resume(uint256)" 6 --private-key "your private key" --rpc-url "https://babel-api.mainnet.iotex.io" --legacy
```
