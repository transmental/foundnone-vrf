# Foundnone VRF

A democratized Verifiable Random Function (VRF) system allowing anyone to request and fulfill entropy requests onchain for rewards. It can be used with a single EOA or an external transaction relayer service enabling concurrency and scalability. The requester can provide a callback address to receive the entropy, or it can be retrieved by polling the contract. On the fulfillment side, the fulfiller can opt to whitelist specific callback addresses and/or specify a maximum amount of callback gas to prevent spam.

# Deployed Addresses:

- Base Sepolia: `0x1ec945E267CF78c53306d48D89f2cdb500026811`
- Curtis (Apechain Testnet): `0x25f27467377DaC26B79784603A0b2DcDaa3b67cf`
- Base: `pending further testing...`

# Quick Start

## Environment Variables

`cp fulfiller/.env.example fulfiller/.env`

```ini
# RPC URLs
WS_RPC_URL=ws://...
HTTP_RPC_URL=http://...

# Deployed FoundnoneVRF contract address on Base Sepolia
CONTRACT_ADDRESS=0x1ec945E267CF78c53306d48D89f2cdb500026811

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

## Requesting Entropy

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
- **Plonk-Based Verification**: Onchain proof verification via Plonk verification. `/contracts/contracts/PlonkVerifier.sol`.
- **Commitment Scheme**: Enforces a predetermined Poseidon-based fulfiller commitment to prevent brute forcing desired outcomes.
- **Separate Reward Receiver**: Fulfillers can allocate rewards to a distinct address.
- **Trusted Setup**: Reliably generated via a reputable Phase 2 Ptau file.
- **Concurrent Fulfillment**: Multiple fulfillers can compete to fulfill requests, with the first valid proof winning the reward.
- **External Transaction Relayer**: Supports running fulfillers as a service, allowing for concurrent processing of multiple requests.

## Trusted Setup

Uses the Zcash Phase 2 Powers of Tau ceremony artifacts from the Privacy Scaling Explorations repository:

- **Source**: https://github.com/privacy-scaling-explorations/perpetualpowersoftau
- **File**: `ppot_0080_11.ptau` (11th-degree, 2,048 Points file)
- **Proof System Assumptions**: Security relies on the Plonk trusted setup (via Phase 2 Ptau). Assumes no participant compromised the Powers of Tau ceremony.

This Phase 2 file was used to generate the circuit zkey (`fulfiller/zk/vrf_final.zkey`).

## Commitment Scheme

Each fulfiller generates a secret `s` and computes a Poseidon hash

```
commitment = Poseidon([s, 0])
```

This `commitment` is set onchain before processing requests. Every proof must include the previous commitment as a public input, binding the entropy generation to the fulfiller’s secret and preventing replay or spoofing.

## Fulfiller Rewards

- Fulfillers earn (100 - contractFeePercentage)% of the request fee.
- Contract retains contractFeePercentage% as platform fees.
- Fulfillers claim their rewards manually via the `withdrawRewardReceiverBalance` function (this should be called by the PAYOUT_ADDRESS specified in the fulfiller service .env file).

## Unfulfilled Requests

- Requests not fulfilled within 256 blocks are considered invalid and can be refunded.
- This is due to the fact that the blockhash of the block - 1 of the request is used as a part of the seed during entropy generation.
- We use the block - 1 blockhash to ensure that fulfillment can still be made inside of the same block as the request.
- Refunds are processed via the `refundUnfulfilledRequest` function, which returns the full request fee to the sender (the contract fee is not taken until fulfillment).

## Why Foundnone VRF is Different

- **Decentralized Fulfillment**: No single oracle—anyone can run a fulfiller service.
- **Fulfiller Incentivization**: Earn rewards for fulfilling requests.
- **Speed and Efficiency**: Fast and efficient onchain verification, proof to fulfillment times can be under 3 seconds depending on network congestion and fulfillers hardware.
- **Competitive Fulfillment**: Fulfiller rewards are based on the number of requests processed.
- **Full Transparency and Easy Setup**: Open-source prover pipeline and Dockerized deployment.

## Contract Testing

- 100% coverage of the smart contract (with full snarkjs prover pipeline) at `contracts/test/FoundnoneVRF.ts`.

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

# Glossary

| Term                                 | Definition                                                                                                                     |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| **VRF (Verifiable Random Function)** | A cryptographic function that produces random outputs along with a proof that the output was correctly generated.              |
| **Entropy**                          | The random output generated by the VRF, stored onchain after successful proof verification.                                    |
| **Fulfillment**                      | The act of submitting a valid proof and entropy for an open VRF request.                                                       |
| **Commitment**                       | A Poseidon hash of a secret known only to the fulfiller, submitted to the contract to prevent precomputing favorable outcomes. |
| **Plonk**                            | A zk-SNARK proof system used for efficient onchain verification of entropy proofs.                                             |
| **Poseidon Hash**                    | A zkSNARK-optimized cryptographic hash function used to create commitments efficiently in zero-knowledge circuits.             |
| **Phase 2 Powers of Tau**            | A publicly verifiable trusted setup ceremony used to generate parameters necessary for zk-SNARK circuits.                      |
| **Request Fee**                      | The amount of ETH paid to submit a VRF request. Split between the fulfiller and the protocol.                                  |
| **Reward Receiver**                  | The address designated by a fulfiller to receive rewards from fulfilled requests.                                              |
| **Contract Fee Percentage**          | The percentage of each request fee that is retained by the contract for operational purposes.                                  |
| **Request ID**                       | A unique identifier for each VRF entropy request.                                                                              |
| **Blockhash**                        | The hash of a recent block, used as an unpredictable seed in entropy generation.                                               |
| **Fulfiller**                        | An entity that runs the prover and submits entropy proofs to the Foundnone VRF contract for rewards.                           |
| **Callback Address**                 | An optional address provided by the requester to receive the entropy directly, or it can be retrieved by polling the contract. |
| **Callback Gas Limit**               | The maximum amount of gas that can be used for the callback function when sending the entropy to the callback address.         |

# FAQ

## ❓ What is Foundnone VRF?

Foundnone VRF is a decentralized Verifiable Random Function system that allows anyone to request and fulfill randomness (entropy) onchain, with provable validity via zero-knowledge proofs (Plonk).

---

## ❓ How do I request random entropy?

Call the `requestRng` function on the FoundnoneVRF contract while sending at least the required `requestFee` in ETH.  
Listen for the `RngRequested` event or poll `entropies[_requestId]` to retrieve the result after fulfillment.

---

## ❓ How do I run a fulfiller and earn rewards?

1. Set up a `.env` file with RPCs, private key, contract address, and payout address.
2. Build the Docker image and run the fulfiller service.
3. The service listens for new requests and submits proofs to earn rewards.

---

## ❓ What is the purpose of the commitment?

The commitment prevents fulfillers from brute-forcing favorable entropy.  
Each fulfiller must set a Poseidon hash commitment onchain before fulfilling any requests, binding their prover to a specific secret.

---

## ❓ How are rewards distributed?

When a fulfiller submits a valid proof:

- `(100 - contractFeePercentage)%` of the request fee is credited to their designated reward receiver address.
- `contractFeePercentage%` goes to the contract’s operational balance.

Fulfillers must call `withdrawRewardReceiverBalance()` to redeem accumulated rewards.

---

## ❓ What happens if a request is not fulfilled?

Requests are only valid for approximately 256 blocks (~1 hour depending on the chain).  
If not fulfilled within that window, requesters can call `refundUnfulfilledRequest(requestId)` to reclaim their full request fee.

---

## ❓ Can multiple fulfillers compete to fulfill a request?

Yes.  
Anyone can attempt to fulfill a request.  
The first valid proof submitted wins the reward for that request.

---

## ❓ What if the blockhash needed for fulfillment expires?

If more than 256 blocks pass, the blockhash is no longer retrievable.  
In that case, fulfillment becomes impossible, and the requester must manually refund the request.

---

## ❓ How secure is the trusted setup?

The Plonk zk-SNARK system used by Foundnone VRF relies on the security of the Phase 2 Powers of Tau ceremony.  
We use artifacts from a reputable source (Privacy Scaling Explorations), assuming no participant in the ceremony colluded to compromise it.

---

## ❓ Can I modify the circuit or the zk artifacts?

Yes, but you must regenerate the `.zkey` and `.wasm` artifacts located in `fulfiller/zk/`.  
If the circuit changes, proofs generated using outdated artifacts will fail verification.

---

## ❓ Is there a limit to how many fulfillers can run simultaneously?

No.  
The system is open — anyone can run a fulfiller at any time.  
Rewards are paid out competitively based on proof submissions.

---

## ❓ How does a fulfiller protect against callback spam?
Fulfillers can:
- Whitelist specific callback addresses to ensure only trusted recipients receive entropy.
- Set a maximum gas limit for callbacks to prevent excessive costs from spammy requests.
