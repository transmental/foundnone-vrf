# Foundnone VRF

A democratized Verifiable Random Function (VRF) system allowing anyone to request and fulfill entropy requests on-chain for rewards.

# Quick Start With Docker

## Environment Variables

`cp fulfiller/.env.example fulfiller/.env`

```ini
# RPC URLs
WS_RPC_URL=ws://...
HTTP_RPC_URL=http://...

# Deployed FoundnoneVRF contract address
CONTRACT_ADDRESS=0x...

# Private key of the ENTROPY_ROLE (do not prefix with 0x)
FULFILLER_PK=...

# Address to receive rewards
PAYOUT_ADDRESS=0x...

# Chain ID (e.g. 84532)
CHAIN_ID=84532

# Optional retry settings for WebSocket re-subscription, defaults to 5
CONNECTION_RETRIES=5
```

## Build & Run Fulfiller

#### Build Docker Image

```bash
docker build --no-cache -t foundnone-vrf-fulfiller .
```

#### Run

```bash
docker run --name foundnone-vrf-fulfiller --env-file fulfiller/.env foundnone-vrf-fulfiller
```

## Request Entropy

Call the `requestRng` function on the `FoundnoneVRF` contract. This will emit a `RngRequested` event with the request ID.

Poll the contract for entropies[_requestId] or listen for the `RequestFulfilled` event for your request ID. The entropy will be available in the `event.publicInputs[1]` field.

## Run the POC Frontend

```bash
cd poc-frontend
npm i
npm run dev
```

# Key Features

- **Open Fulfillment**: Anyone can run a fulfiller to process VRF requests and earn rewards.
- **Plonk-Based Verification**: On-chain proof verification via a Plonk verifier.
- **Commitment Scheme**: Uses a predetermined Poseidon-based commitment to prevent spoofing.
- **Separate Reward Receiver**: Fulfillers can allocate rewards to a distinct address.
- **Trusted Setup**: Reliably generated via a reputable Phase 2 Ptau file.

## Trusted Setup

Uses the Zcash Phase 2 Powers of Tau ceremony artifacts from the Privacy Scaling Explorations repository:

- **Source**: https://github.com/privacy-scaling-explorations/perpetualpowersoftau
- **File**: `ppot_0080_11.ptau` (11th-degree, 2,048 Points file)

This Phase 2 file was used to generate the circuit zkey (`fulfiller/zk/vrf_final.zkey`).

## Commitment Scheme

Each fulfiller generates a secret `s` and computes a Poseidon hash

```
commitment = Poseidon([s, 0])
```

This `commitment` is set on-chain before processing requests. Every proof must include the previous commitment as a public input, binding the entropy generation to the fulfiller’s secret and preventing replay or spoofing.

## Why Foundnone VRF is Different

- **Decentralized Fulfillment**: No single oracle—anyone can run a fulfiller service.
- **Gas Efficiency**: PLONK verifier instead of Groth16, reducing verifier gas to a fraction of other implementations.
- **Extra Flexibility**: Reward allocation and commitment rotation on each fulfillment.
- **Full Transparency**: Open-source prover pipeline and Dockerized deployment.

## Repository Structure

```
.
├── README.md
├── circom
│   └── vrf.circom
├── contracts
│   ├── contracts
│   │   ├── FoundnoneVRF.sol
│   │   └── PlonkVerifier.sol
│   ├── hardhat.config.ts
│   ├── ignition              # Deployment artifacts via Hardhat Ignition
│   ├── test
│   │   └── FoundnoneVRF.ts
│   └── tsconfig.json
├── docker-compose.yml
├── Dockerfile
├── fulfiller                # Go-based fulfiller service
│   ├── abi
│   ├── main.go
│   ├── prover.js            # Lightweight local prover server
│   ├── zk                   # zkey and wasm for node runtime
│   └── go.mod/go.sum
├── package.json             # Node.js scripts and dependencies
└── ...
```

## Contract Testing

```bash
cd contracts
npx hardhat test
```

## Redeeming Rewards

Go to the contract address you have made fulfillments on, and call the `withdrawRewardReceiverBalance` function from the address that you set as the `PAYOUT_ADDRESS` in your `.env` file.
This will transfer the rewards to that address.

## Contributing

1. Fork and clone the repo
2. Install dependencies (`npm install`, `go mod tidy`)
3. Make your changes and add tests
4. Submit a pull request

## License

MIT © 2025 Zachary Owens
