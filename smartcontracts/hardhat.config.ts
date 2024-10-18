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
      url: process.env.URL != undefined ? process.env.URL : 'http://127.0.0.1:8545',
      accounts: accounts,
    },
    nightly: {
      url: 'https://babel-nightly.iotex.io',
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
  etherscan: {
    apiKey: 'YOUR_ETHER',
    customChains: [
      {
        network: 'mainnet',
        chainId: 4689,
        urls: {
          apiURL: 'https://IoTeXscout.io/api',
          browserURL: 'https://IoTeXscan.io',
        },
      },
      {
        network: 'testnet',
        chainId: 4690,
        urls: {
          apiURL: 'https://testnet.IoTeXscout.io/api',
          browserURL: 'https://testnet.IoTeXscan.io',
        },
      },
    ],
  },
};

export default config;
