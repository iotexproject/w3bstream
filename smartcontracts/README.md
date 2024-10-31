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

| Contract                   | Address                                    |
|-----------------------------|--------------------------------------------|
| W3bstreamProject            | `0x0abec44FC786e8da12267Db5fdeB4311AD1A0A8A` |
| ProjectRegistrar            | `0x74309Bc83fF7Ba8aBaB901936927976792a6d9B6` |
| W3bstreamProver             | `0xab6836908d15E42D30bdEf14cbFA4ad45dCAF3a3` |
| W3bstreamVMType             | `0x5b27FC853058C1e50C252c017e5859AcF854F3b4` |
| W3bstreamDebits             | `0x0AD341EfF116eeee2451d105133F7759FE4c2e4f` |
| W3bstreamProjectReward      | `0xfb3E89d1ED4b43F2D4D76400D95f4C158Fc02aC0` |
| W3bstreamTaskManager        | `0xF0714400a4C0C72007A9F910C5E3007614958636` |
| W3bstreamRouter             | `0x28E0A99A76a467E7418019cBBbF79E4599C73B5B` |
| W3bstreamRewardDistributor  | `0x058f2F501EC0505B9CF8AB361FFFBFd36C83a8aF` |
| W3bstreamMinter             | `0x49C096AE869A3054Db06ffF221b917b41f94CEf3` |

| Setting                   | Value                                    |
|-----------------------------|--------------------------------------------|
| ProjectRegistrar Fee        | `1.0 IOTX`                                      |
| W3bstreamMinter Reward      | `0 IOTX`                                        |
