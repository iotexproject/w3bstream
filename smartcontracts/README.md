# W3bstream contracts

## Deploy

```bash
export PROJECT_ADDRESS=0xf07336E1c77319B4e740b666eb0C2B19D11fc14F
export PROJECT_REGISTRATION_FEE=1.0
export PROVER_REGISTRATION_FEE=1.0
export MIN_STAKE=1.0
./deploy.sh --network mainnet|testnet|dev
```

### Deployment

#### Testnet

LivenessVerifier deployed to 0x2eAfC6918B04DBB70F2E1c19D7667cc800E3238e  
MovementVerifier deployed to 0x2842F0442bA99886503f5d87c9A2868946A433b3  
MockDapp deployed to 0xAA0cB09c6eed8B43DFb0852b4445C2Db86E85A73  
MockDappLiveness deployed to 0x24B8D8f110Be334b36ebBA567a423b052444e66B  
MockDappMovement deployed to 0x799b5dc3fa1F0522F145D37f939187c926f035Bb  
W3bstreamProject deployed to 0x7D3158166E9298fC47beA036fE162fEA17632E5D  
ProjectRegistrar deployed to 0xB564996622CE5610b9cF4ed35160f406185d7d0b  
W3bstreamProject binder set to ProjectRegistrar 0xB564996622CE5610b9cF4ed35160f406185d7d0b  
ProjectRegistrar registration fee set to 0.0  
W3bstreamProver deployed to 0x7EdDe171F0944252b5A2CCE5Ef0eD7e13310F67F  
W3bstreamVMType deployed to 0x2Fb953ca5e7E693a8Ce2C2E26dAE33616980Bc7b  
W3bstreamDebits deployed to 0x862116BBBBD18Af494A2680D5430C25A9C47ad0a  
W3bstreamProjectReward deployed to 0x65EE0E96D0818aD43d93ae8089Db5F0d658c0cCc  
W3bstreamTaskManager deployed to 0xEe96E984E6e746aAbe4a5ef687Ad8bb2aA98bDf3  
W3bstreamRouter deployed to 0x19dD7163Ad80fE550C97Affef49E1995B24941B1  
W3bstreamRewardDistributor deployed to 0xe8d9ccb6e7E1c834A1E2E1a6dF3403773eAa580e  
W3bstreamMinter deployed to 0x7cE64E8a918C9B8b6172043Bb50E99A9AFf72cd1  
MockIoID deployed to 0xDF505e158Ffe093601494D931924Bb24E8B7B92D  
W3bstreamRewardDistributor add operator to 0x7cE64E8a918C9B8b6172043Bb50E99A9AFf72cd1  
W3bstreamTaskManager add operator to 0x7cE64E8a918C9B8b6172043Bb50E99A9AFf72cd1  
W3bstreamTaskManager add operator to 0x19dD7163Ad80fE550C97Affef49E1995B24941B1  
W3bstreamRewardDistributor set operator to 0x7cE64E8a918C9B8b6172043Bb50E99A9AFf72cd1  
W3bstreamDebits set operator to 0xEe96E984E6e746aAbe4a5ef687Ad8bb2aA98bDf3  
W3bstreamMinter set reward to 0  