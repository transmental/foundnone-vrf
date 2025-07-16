# Foundnone VRF

A democratized Verifiable Random Function (VRF) system allowing anyone to request and fulfill entropy requests onchain for rewards. It can be used with a single EOA, an external transaction relayer service, or KMS Wallets for secure and concurrent transaction handling, enabling scalability and competition. The requester can provide a callback address to receive the entropy, or it can be retrieved by polling the contract. On the fulfillment side, the fulfiller can opt to whitelist specific callback addresses and/or specify a maximum amount of callback gas to prevent spam.

# Deployed Addresses:

- Base Sepolia: `0x6011C31271b321FcE089FB898ecd487BA96CC73f`
- Curtis (Apechain Testnet): `0x25f27467377DaC26B79784603A0b2DcDaa3b67cf`
- Base: `pending further testing...`

# Quick Start

## Environment Variables

`cp fulfiller/.env.example fulfiller/.env`

```ini
# Required

# The URL of the websocket RPC endpoint to connect to.
WS_RPC_URL=
# The URL of the HTTP RPC endpoint to connect to.
HTTP_RPC_URL=
# The address of the contract to listen to requests on and send fulfillment transactions to.
CONTRACT_ADDRESS=0xc907a3187c3900D98F2fB94aED4f3c2E54cf8cD9
# The address of the account that will be credited with the fulfillment rewards.
PAYOUT_ADDRESS=
# The chain ID of the network you are connecting to.
CHAIN_ID=84532
# The private key of the account that will be used to sign fulfillment transactions. Do not prefix with '0x'.
FULFILLER_PK=

# Optional
# The number of retries to make if the websocket connection fails.
CONNECTION_RETRIES=5
# The base URL of the relayer to use for sending transactions, this is preferably an internal URL.
# When set, the fulfillment is sent optimistically and the relayer service will handle fulfillment. (non blocking)
RELAYER_URL=
# Must be set if the RELAYER_URL is not set, if fulfilling is done by this golang service
# it is blocking and will wait for each fulfillment transaction to be mined.
FULFILLER_PK=
#Optional, defaults to 10, the number of concurrent fulfillments that can be processed by an external transaction relayer service.
RELAYER_CONCURRENCY_LIMIT=10
# Optional, defaults to 100000, the maximum gas limit for each callback transaction.
MAX_CALLBACK_GAS_LIMIT=100000
# Optionally set a list of whitelisted callback addresses that can be used to fulfill requests.
# If set, only events with either the zero address or one of the whitelisted addresses will be processed.
WHITELISTED_CALLBACK_ADDRESSES=0x1ec945E267CF78c53306d48D89f2cdb500026811,0x...

# KMS Wallet Configuration
# Enables transaction concurrency for higher throughput and competition.
KMS_KEY=<your_kms_key>
KMS_REGION=<your_kms_region>
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
- **KMS Wallet Support**: Enables transaction concurrency for higher throughput and competition by securely managing keys and automating account handling.
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

## KMS Wallet Support

The Foundnone VRF system includes support for KMS (Amazon Key Management Service) wallets, enabling secure and scalable management of private keys. This feature is designed to achieve transaction concurrency, allowing fulfillers to handle multiple requests simultaneously, thereby increasing competition and throughput.

### Key Features of KMS Wallet Integration

- **Secure Key Management**: Private keys are securely managed within AWS KMS, ensuring they are never exposed outside the secure environment.
- **Automated Account Handling**: Automatically generates, encrypts, and stores private keys along with their commitments and secrets in a Postgres database.
- **Higher Throughput**: Supports concurrent transactions for faster processing and increased competition among fulfillers.
- **Seamless Ethereum Integration**: Provides a `SignerFn` compatible with go-ethereum for signing transactions directly using KMS.

### How KMS Wallet Works

Important: you must run a Postgres database and have AWS KMS configured with the necessary permissions to use the KMS key for encryption and decryption.

1. **KeyVault**:
   - The `KeyVault` is the core component for managing private keys.
   - It interacts with AWS KMS to encrypt and decrypt private keys and stores the ciphertext in a Postgres database.
   - **Key Functions**:
     - **EncryptAndStoreKey**: Encrypts a private key using AWS KMS and stores it in the database along with its associated Ethereum address, commitment, and secret.
     - **LoadAllKeys**: Retrieves all keys, commitments, and secrets from the database and decrypts them using AWS KMS.
     - **GenerateAndStoreKeys**: Generates new Ethereum private keys, computes their commitments and secrets, and securely stores them.

2. **KMSWallet**:
   - Represents a single Ethereum wallet backed by AWS KMS.
   - Derives the Ethereum address from the public key stored in KMS.
   - Provides a `SignerFn` for signing Ethereum transactions directly using KMS.
   - **Key Functions**:
     - **deriveAddressFromKMS**: Fetches the public key from AWS KMS and computes the Ethereum address.
     - **SignerFn**: Signs transactions using AWS KMS and computes the recovery ID (`V`) to produce a valid Ethereum signature.
     - **GatherFunds**: (Placeholder) Sweeps all ETH from the KMS account to a specified address.

3. **Transaction Signing**:
   - Transactions are signed using AWS KMS, ensuring that private keys never leave the secure KMS environment.
   - The `SignerFn` function handles Ethereum-specific signing requirements, including recovery ID computation.

4. **Concurrency and Scalability**:
   - The `KeyVault` can manage multiple keys, making it suitable for applications requiring high throughput.
   - The `KMSWallet` uses a mutex (`mu`) to ensure thread-safe operations, enabling concurrent transaction handling.

### How to Use KMS Wallet

1. **Set Up Environment Variables**:
   - Add the following variables to your `.env` file:

     ```ini
     KMS_KEY=<your_kms_key>
     KMS_REGION=<your_kms_region>
     ```

2. **Monitor Logs**:
   - Check the logs to verify that accounts are loaded, funded, and committed successfully.

### Additional Requirements for KMS Wallet Mode

When running the fulfiller in KMS mode, the following environment variables are required or optional for proper configuration:

#### Required Variables

- **Primary Wallet Private Key (`FULFILLER_PK`)**:
  - The private key of the primary wallet is still required to fund the KMS wallets that are spun up.
  - This wallet is used to bootstrap the KMS-backed accounts with initial ETH for gas.

- **Postgres Connection String (`PG_CONN_STRING`)**:
  - The connection string for the Postgres database where encrypted keys, commitments, and secrets are stored.

- **AWS Credentials**:
  - **`AWS_ACCESS_KEY_ID`**: The access key ID for AWS.
  - **`AWS_SECRET_ACCESS_KEY`**: The secret access key for AWS.
  - **`AWS_DEFAULT_REGION`**: The default AWS region for KMS operations.

#### Optional Variables

- **`MAX_ACCOUNTS`**:
  - Specifies the maximum number of KMS-backed accounts to manage.
  - Default: `5`.

- **`WALLET_MODE`**:
  - Specifies the wallet mode. Set to `kms` to enable KMS wallet functionality.

- **`POOL_MIN_GAS_WEI`**:
  - The minimum gas threshold (in wei) for KMS-backed accounts.
  - Default: `100000000000000` (0.005 ETH).

- **`POOL_REFILL_AMOUNT_WEI`**:
  - The amount of ETH (in wei) to refill KMS-backed accounts when their balance falls below the threshold.
  - Default: `1000000000000000` (0.01 ETH).

### Example `.env` Configuration for KMS Mode

```ini
# Required Variables
FULFILLER_PK="<primary_wallet_private_key>"
PG_CONN_STRING="<postgres_connection_string>"
AWS_ACCESS_KEY_ID="<aws_access_key_id>"
AWS_SECRET_ACCESS_KEY="<aws_secret_access_key>"
AWS_DEFAULT_REGION="<aws_default_region>"

# Optional Variables
MAX_ACCOUNTS=5
WALLET_MODE=kms
POOL_MIN_GAS_WEI=100000000000000
POOL_REFILL_AMOUNT_WEI=1000000000000000
```

### Running Database Migrations

To apply all unapplied migrations in the `fulfiller/migrations` folder, you can use the `golang-migrate` tool. Follow these steps:

1. **Install `golang-migrate`**:

   ```zsh
   brew install golang-migrate
   ```

2. **Run All Unapplied Migrations**:

   ```zsh
   migrate -path fulfiller/migrations -database "<postgres_connection_string>" up
   ```

   Replace `<postgres_connection_string>` with your actual Postgres connection string.

3. **Verify Applied Migrations**:
   To check which migrations have been applied, run:
   ```zsh
   migrate -path fulfiller/migrations -database "<postgres_connection_string>" version
   ```

### Benefits of KMS Wallet Integration

- **Enhanced Security**: Keys are securely managed within AWS KMS, reducing the risk of exposure.
- **Simplified Management**: Automates key generation, funding, and commitment setup.
- **Improved Scalability**: Handles multiple accounts and transactions concurrently, increasing throughput and competition.
- **Ethereum Compatibility**: Provides seamless integration with Ethereum's transaction signing requirements.

### Debugging and Logging

- Debug logs are included to provide visibility into key storage, updates, and signing operations.
- Example log output:
  ```
  [KMS DEBUG] Storing key for address 0x123... with commitment abc123 and secret xyz456 and ciphertext ...
  [KMS DEBUG] Updated commitment and secret for address 0x123... with comm: abc123, secret xyz456
  ```

## Why Foundnone VRF is Different

- **Decentralized Fulfillment**: No single oracle—anyone can run a fulfiller service.
- **Fulfiller Incentivization**: Earn rewards for fulfilling requests.
- **Speed and Efficiency**: Fast and efficient onchain verification, proof to fulfillment times can be under 2 seconds depending on network congestion and fulfiller hardware.
- **Competitive Fulfillment**: Fulfiller rewards are based on the number of requests processed.
- **KMS Wallet Integration**: Achieves higher throughput and competition by enabling concurrent transaction handling.
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
| **KMS Wallet**                       | A wallet backed by AWS Key Management Service (KMS) for secure and scalable private key management.                            |
| **KeyVault**                         | A component that manages private keys, encrypts them using KMS, and stores them in a Postgres database.                        |
| **SignerFn**                         | A function provided by the KMS wallet to sign Ethereum transactions directly using KMS.                                        |
| **POOL_MIN_GAS_WEI**                 | The minimum gas threshold (in wei) for KMS-backed accounts. Default: 0.005 ETH.                                                |
| **POOL_REFILL_AMOUNT_WEI**           | The amount of ETH (in wei) to refill KMS-backed accounts when their balance falls below the threshold. Default: 0.01 ETH.      |
| **MAX_ACCOUNTS**                     | The maximum number of KMS-backed accounts to manage. Default: 5.                                                               |
| **WALLET_MODE**                      | Specifies the wallet mode. Set to `kms` to enable KMS wallet functionality.                                                    |

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

---
