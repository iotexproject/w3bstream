# W3bstream contracts

npx hardhat test

## Deploy

```bash
export PROJECT_REGISTRATION_FEE=1.0
export PROVER_REGISTRATION_FEE=1.0
export MIN_STAKE=1.0
yarn hardhat run scripts/deploy.ts --network testnet
```

### Deployment

#### Testnet

```
W3bstreamProject deployed to 0xe2267bC7fF61371d0Ad85f5A8e44063786266495
ProjectRegistrar deployed to 0xF6BF9f1E7ec17b72Defbd90874359fbb513DeD38
W3bstreamProver deployed to 0xb5De017C1E1f1b98AC093A26299f8F95Ed4e1F82
W3bstreamCredit deployed to 0x5ac81e25e3916896143BEC07176aA5B47573f053
FleetManagement deployed to 0xD4C35cdCbE43f6567eD85cD62ad075aA52FD1377
W3bstreamRouter deployed to 0x301EEe80c6d9d0299F8E778C3cfA415682EA3417
```
