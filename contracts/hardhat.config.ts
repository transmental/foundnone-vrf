import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-ignition";
import "@nomicfoundation/hardhat-verify";
import "dotenv/config";
import "@nomicfoundation/hardhat-viem";
import "hardhat-gas-reporter";
import "@nomicfoundation/hardhat-chai-matchers";
import "@tenderly/hardhat-tenderly";
import "solidity-coverage";

const config: HardhatUserConfig = {
  networks: {
    base: {
      url: "https://mainnet.base.org",
      accounts: [process.env.PRIVATE_KEY!],
      chainId: 8453,
    },
    baseSepolia: {
      url: "https://sepolia.base.org",
      accounts: [process.env.PRIVATE_KEY!],
      chainId: 84532,
    },
    curtis: {
      url: "https://curtis.rpc.caldera.xyz/http",
      accounts: [process.env.PRIVATE_KEY!],
      chainId: 33111,
    },
    apeChain: {
      url: "https://apechain.io",
      accounts: [process.env.PRIVATE_KEY!],
      chainId: 33139,
    },
  },
  solidity: {
    version: "0.8.28",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  etherscan: {
    apiKey: {
      base: process.env.BASESCAN_API_KEY!,
      baseSepolia: process.env.BASESCAN_API_KEY!,
      curtis: process.env.APESCAN_API_KEY!,
      apeChain: process.env.APESCAN_API_KEY!,
    },
    customChains: [
      {
        chainId: 84532,
        network: "base",
        urls: {
          apiURL: "https://api.basescan.org/api",
          browserURL: "https://basescan.org",
        },
      },
      {
        chainId: 84532,
        network: "baseSepolia",
        urls: {
          apiURL: "https://api-sepolia.basescan.org/api",
          browserURL: "https://sepolia.basescan.org",
        },
      },
      {
        network: "apeChain",
        chainId: 33139,
        urls: {
          apiURL: "https://api.apescan.io/api",
          browserURL: "https://apescan.io",
        },
      },
      {
        network: "curtis",
        chainId: 33111,
        urls: {
          apiURL: "https://api-curtis.apescan.io/api",
          browserURL: "https://curtis.apescan.io",
        },
      },
    ],
  },
};

export default config;
