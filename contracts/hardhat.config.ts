import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-ignition";
import "@nomicfoundation/hardhat-verify";
import "dotenv/config";
import "@nomicfoundation/hardhat-viem";
import "hardhat-gas-reporter";
import "@nomicfoundation/hardhat-chai-matchers";
import "solidity-coverage";

const config: HardhatUserConfig = {
  networks: {
    baseSepolia: {
      url: "https://sepolia.base.org",
      accounts: [process.env.PRIVATE_KEY!],
      chainId: 84532,
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
      baseSepolia: process.env.BASESCAN_API_KEY!,
    },
    customChains: [
      {
        chainId: 84532,
        network: "baseSepolia",
        urls: {
          apiURL: "https://api-sepolia.basescan.org/api",
          browserURL: "https://sepolia.basescan.org",
        },
      },
    ],
  },
};

export default config;
