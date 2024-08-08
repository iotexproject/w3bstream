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
W3bstreamProject deployed to 0x280b75d2d0b35FAfdF7f97594908C2C95ddBb033
ProjectRegistrar deployed to 0x892e082B399e578a589994811c5347720DBd67d4
W3bstreamProver deployed to 0x951Ca40DC508d8F2Aa28F3fa802705C57A553490
W3bstreamCredit deployed to 0xF0A34EeC1a22415D7426D2277cBEa69fabFaa022
FleetManagement deployed to 0xc44264387042474c64B3914874bd6E02f669d294
W3bstreamRouter deployed to 0x1798380df49C0e73aCd6254fd10d91E853Ff6CfF
W3bstreamVMType deployed to 0x0B35Bc3c077A15be050D96C122c0fDf7e6a41AaF
ProjectDevice deployed to 0xEA0B75d277AE1D13BBeAAe4537291319E2d3d1C2
```
