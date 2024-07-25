# W3bstream contracts

npx hardhat test

## Deploy

```bash
export PROJECT_ADDRESS=0x6972C35dB95258DB79D662959244Eaa544812c5A
export PROJECT_REGISTRATION_FEE=1.0
export PROVER_REGISTRATION_FEE=1.0
export MIN_STAKE=1.0
yarn hardhat run scripts/deploy.ts --network testnet
```

### Deployment

#### Testnet

```
W3bstreamProject deployed to 0x6AfCB0EB71B7246A68Bb9c0bFbe5cD7c11c4839f
ProjectRegistrar deployed to 0x4888bfbf39Dc83C19cbBcb307ccE8F7F93b72E38
W3bstreamProver deployed to 0xAD480a9c1B9fA8dD118c26Ac26880727160D0448
W3bstreamCredit deployed to 0x0ad7d4bBC1c839b33404Cd32fB8FB65D9ec5d5b6
FleetManagement deployed to 0xDBA78C8eCaeE2DB9DDE0c4168f7E8626d4Ff0010
W3bstreamRouter deployed to 0x90A27ab74E790Cef6e258aabee1B361a9c993e8b
ProjectDevice deployed to 0x3d6b6c7bDB72e8BF73780f433342759d8b329Ca5
W3bstreamVMType deployed to 0x7f0B05758914e8B1C8Bf2DA3419BED995625fd99
```
