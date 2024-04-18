import * as dotenv from 'dotenv';
import { HardhatUserConfig } from 'hardhat/config';
import '@nomicfoundation/hardhat-toolbox';
import '@openzeppelin/hardhat-upgrades';

dotenv.config();

const PRIVATE_KEY = process.env.PRIVATE_KEY;
const accounts = PRIVATE_KEY !== undefined ? [PRIVATE_KEY] : [];

const config: HardhatUserConfig = {
  networks: {
    hardhat: {
      allowUnlimitedContractSize: true,
    },
    dev: {
      url: 'http://127.0.0.1:8545',
      accounts: accounts,
    },
    mainnet: {
      url: 'https://babel-api.mainnet.iotex.io/',
      accounts: accounts,
    },
    testnet: {
      url: 'https://babel-api.testnet.iotex.io/',
      accounts: accounts,
    },
  },
  solidity: {
    compilers: [
      {
        version: '0.8.19',
        settings: {
          viaIR: true,
          optimizer: {
            enabled: true,
            runs: 10000,
          },
          metadata: {
            bytecodeHash: 'none',
          },
        },
      },
    ],
  },
};

export default config;
