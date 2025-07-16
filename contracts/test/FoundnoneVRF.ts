import {
  loadFixture,
  mine,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import hre from "hardhat";
import { expect } from "chai";
import fs from "fs";
import path from "path";
import { execSync } from "child_process";
import { parseEther, encodeAbiParameters, keccak256, zeroAddress } from "viem";
import { buildPoseidon } from "circomlibjs";

describe("FoundnoneVRF full ZK flow", function () {
  it("Poseidon RNG passes frequency and chi-squared tests (cryptographic quality)", async function () {
    // This test checks for uniformity and randomness using bucket frequency and chi-squared test
    const NUM_SAMPLES = 10000;
    const NUM_BUCKETS = 100;
    const outputs: bigint[] = [];
    const secret = BigInt(123456);
    for (let i = 0; i < NUM_SAMPLES; i++) {
      const seed = BigInt(i);
      const out = poseidon([poseidon.F.e(secret), poseidon.F.e(seed)]);
      outputs.push(poseidon.F.toObject(out));
    }
    // Frequency test: bucket the outputs
    const min = outputs.reduce((a, b) => (a < b ? a : b), outputs[0]);
    const max = outputs.reduce((a, b) => (a > b ? a : b), outputs[0]);
    const bucketSize = (max - min) / BigInt(NUM_BUCKETS);
    const buckets = Array(NUM_BUCKETS).fill(0);
    for (const out of outputs) {
      let idx = Number((out - min) / bucketSize);
      if (idx >= NUM_BUCKETS) idx = NUM_BUCKETS - 1;
      buckets[idx]++;
    }
    // Chi-squared test
    const expected = NUM_SAMPLES / NUM_BUCKETS;
    const chi2 = buckets.reduce(
      (sum, count) => sum + (count - expected) ** 2 / expected,
      0
    );
    // For 99 degrees of freedom, chi2 < ~135 for 95% confidence
    console.log("Poseidon RNG chi2:", chi2, "buckets:", buckets);
    expect(chi2).to.be.lessThan(135);
    // Optionally: check that all buckets are non-empty
    expect(buckets.every((x) => x > 0)).to.be.true;
  });
  it("brute force tests Poseidon RNG dispersion with incrementing inputs", async function () {
    // This test checks the dispersion of Poseidon RNG outputs for incrementing inputs
    // to ensure randomness quality (no clustering, good spread)
    const NUM_SAMPLES = 1000;
    const outputs: bigint[] = [];
    const seen = new Set<string>();
    // Use a fixed secret, increment the seed
    const secret = BigInt(123456);
    for (let i = 0; i < NUM_SAMPLES; i++) {
      const seed = BigInt(i);
      const out = poseidon([poseidon.F.e(secret), poseidon.F.e(seed)]);
      const outBig = poseidon.F.toObject(out);
      outputs.push(outBig);
      seen.add(outBig.toString());
    }
    // Check for uniqueness (should be no collisions)
    expect(seen.size).to.equal(NUM_SAMPLES);
    // Check for good spread: compute min, max, mean, and stddev
    const min = outputs.reduce((a, b) => (a < b ? a : b), outputs[0]);
    const max = outputs.reduce((a, b) => (a > b ? a : b), outputs[0]);
    const mean = outputs.reduce((a, b) => a + b, 0n) / BigInt(NUM_SAMPLES);
    const stddev = Math.sqrt(
      outputs.map((x) => Number(x - mean) ** 2).reduce((a, b) => a + b, 0) /
        NUM_SAMPLES
    );
    // Print stats for manual inspection
    console.log("Poseidon RNG stats:", {
      min: min.toString(),
      max: max.toString(),
      mean: mean.toString(),
      stddev,
    });
    // Assert that the spread is reasonable (stddev is a significant fraction of the range)
    const range = Number(max - min);
    expect(stddev).to.be.greaterThan(range * 0.2);
    expect(stddev).to.be.lessThan(range * 0.6);
  });
  // instantiate Poseidon once
  let poseidon: any;
  const BN128_PRIME = BigInt(
    "21888242871839275222246405745257275088548364400416034343698204186575808495617"
  );

  async function deployEntropyFixture() {
    const [admin, fulfiller, sender, otherAccount] =
      await hre.viem.getWalletClients();
    const entropy = await hre.viem.deployContract("FoundnoneVRF", [
      admin.account.address,
    ]);
    const mockCallBackReceiver = await hre.viem.deployContract(
      "MockEntropyReceiver",
      [entropy.address]
    );
    const publicClient = await hre.viem.getPublicClient();
    return {
      admin,
      fulfiller,
      sender,
      entropy,
      mockCallBackReceiver,
      publicClient,
      otherAccount,
    };
  }

  before(async () => {
    poseidon = await buildPoseidon();
  });

  it("does a full ZK prove then submitEntropy and updates commitment, fulfill callback, award the proper fee to the fulfiller and contract, and allow the withdrawal of the fee, as well as reject a refund request", async function () {
    const {
      fulfiller,
      sender,
      entropy,
      publicClient,
      admin,
      otherAccount,
      mockCallBackReceiver,
    } = await loadFixture(deployEntropyFixture);

    // 1) pick a private secret, build initial commitment = Poseidon(secret, 0)
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    // push that to chain
    const commitmentTxHash = await entropy.write.setCommitment(
      [initialCommHex as any],
      {
        account: sender.account,
      }
    );

    await publicClient.waitForTransactionReceipt({
      hash: commitmentTxHash,
    });

    // 2) request a VRF
    const reqTx = await entropy.write.requestRng(
      [mockCallBackReceiver.address, 350_000],
      {
        value: parseEther("0.000005"),
        account: otherAccount.account,
      }
    );
    const reqTx2 = await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
      account: otherAccount.account,
    });
    const reqReceipt = await publicClient.waitForTransactionReceipt({
      hash: reqTx,
    });
    const reqReceipt2 = await publicClient.waitForTransactionReceipt({
      hash: reqTx2,
    });
    expect(reqReceipt.status).to.equal("success");

    // Fetch block details using getBlock
    const block = await publicClient.getBlock({
      blockNumber: reqReceipt.blockNumber - 1n,
    });
    const blockNumber = BigInt(block.number);
    const blockHash = block.hash;

    // requestId will be 1 (first one)
    const requestId = BigInt(1);

    // 3) recompute the seed = keccak256(requestId, blockNumber) % BN128_PRIME
    const packed = encodeAbiParameters(
      [{ type: "uint256" }, { type: "uint256" }, { type: "bytes32" }],
      [requestId, blockNumber, blockHash]
    );
    const seedHash = keccak256(packed);
    const seedBig = BigInt(seedHash) % BN128_PRIME;

    // 4) recompute the entropy = Poseidon(secret, seed)
    const entField = poseidon([poseidon.F.e(secret), poseidon.F.e(seedBig)]);
    const entropySignal = poseidon.F.toObject(entField);

    // 5) write a one-off temp input.json
    const tmp = fs.mkdtempSync(path.join(__dirname, "zktemp-"));
    const inp = {
      secret: secret.toString(),
      seed: seedBig.toString(),
      entropy: entropySignal.toString(),
      commitment: initialCommitment.toString(),
    };
    const inPath = path.join(tmp, "input.json");
    fs.writeFileSync(inPath, JSON.stringify(inp));

    // 6) snarkjs fullprove + export calldata
    execSync(
      `snarkjs plonk fullprove ${inPath} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_js/vrf.wasm"
      )} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_final.zkey"
      )} ${tmp}/proof.json ${tmp}/public.json`,
      { stdio: "inherit" }
    );
    const calldata = execSync(
      `snarkjs zkey export soliditycalldata ${tmp}/public.json ${tmp}/proof.json --plonk`,
      { encoding: "utf8" }
    ).trim();

    // parse proof[] and publicSignals[]
    const [proofPart, publicPart] = calldata.split("][");
    const proof: string[] = JSON.parse(proofPart + "]");
    const publicSignals: string[] = JSON.parse("[" + publicPart);

    expect(proof.length).to.equal(24);
    expect(publicSignals.length).to.equal(3);

    // 8) call submitEntropy(...)
    const tx = await entropy.write.submitEntropy(
      [
        proof as any,
        publicSignals as any,
        requestId,
        fulfiller.account.address,
      ],
      {
        account: sender.account,
      }
    );
    const receipt = await publicClient.waitForTransactionReceipt({
      hash: tx,
    });
    expect(receipt.status).to.equal("success");

    await mine(10);
    const callbackEntropy = await mockCallBackReceiver.read.latestEntropy();
    expect(callbackEntropy).to.equal(entropySignal);

    console.log("callback entropy:", callbackEntropy);
    console.log("entropy sig:", entropySignal);

    const fulfillerContractBalance =
      await entropy.read.getRewardReceiverBalance([fulfiller.account.address]);

    const contractFeePercentage = await entropy.read.contractFeeBasisPoints();

    const requestFee = await entropy.read.requestFee();
    const fee = (requestFee * contractFeePercentage) / 10_000n;

    const expectedFulfillerBalance = requestFee - fee;

    expect(fulfillerContractBalance).to.equal(expectedFulfillerBalance);

    const res = await entropy.write.withdrawRewardReceiverBalance({
      account: fulfiller.account,
    });

    const resReceipt = await publicClient.waitForTransactionReceipt({
      hash: res,
    });
    expect(resReceipt.status).to.equal("success");
    const newFulfillerContractBalance =
      await entropy.read.getRewardReceiverBalance([fulfiller.account.address]);
    expect(newFulfillerContractBalance).to.equal(0n);

    const contractBalance = await entropy.read.contractFeeBalance();
    const expectedContractBalance = fee;
    expect(contractBalance).to.equal(expectedContractBalance);
    const res2 = await entropy.write.withdrawContractFees({
      account: admin.account,
    });
    const resReceipt2 = await publicClient.waitForTransactionReceipt({
      hash: res2,
    });
    expect(resReceipt2.status).to.equal("success");
    const newContractBalance = await entropy.read.contractFeeBalance();
    expect(newContractBalance).to.equal(0n);

    expect(entropy.read.getEntropy([requestId])).to.not.be.rejectedWith();

    // expect the refund of the request to revert
    await expect(
      entropy.write.refundUnfulfilledRequest([requestId], {
        account: otherAccount.account,
      })
    ).to.be.rejectedWith("RequestAlreadyFulfilled()");
    // cleanup
    fs.rmSync(tmp, { recursive: true, force: true });
  });

  it("should revert if the fee is not enough", async function () {
    const { sender, entropy } = await loadFixture(deployEntropyFixture);

    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });

    await expect(
      entropy.write.requestRng([zeroAddress, 0], {
        account: sender.account,
        value: parseEther("0.0000001"),
      })
    ).to.be.rejectedWith("InsufficientFee()");

    await expect(entropy.read.getEntropy([1n])).to.be.rejectedWith(
      "RequestNotFulfilled()"
    );
  });

  it("should revert if the proof is invalid", async function () {
    const { fulfiller, sender, entropy, publicClient, admin } =
      await loadFixture(deployEntropyFixture);

    // 1) pick a private secret, build initial commitment = Poseidon(secret, 0)
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    // push that to chain
    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });

    // 2) request a VRF
    const reqTx = await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
    });
    const reqReceipt = await publicClient.waitForTransactionReceipt({
      hash: reqTx,
    });
    expect(reqReceipt.status).to.equal("success");

    // Fetch block details using getBlock
    const block = await publicClient.getBlock({
      blockNumber: reqReceipt.blockNumber - 1n,
    });
    const blockNumber = BigInt(block.number);
    const blockHash = block.hash;

    // requestId will be 1 (first one)
    const requestId = BigInt(1);

    // 3) recompute the seed = keccak256(requestId, blockNumber) % BN128_PRIME
    const packed = encodeAbiParameters(
      [{ type: "uint256" }, { type: "uint256" }, { type: "bytes32" }],
      [requestId, blockNumber, blockHash]
    );
    const seedHash = keccak256(packed);
    const seedBig = BigInt(seedHash) % BN128_PRIME;

    // 4) recompute the entropy = Poseidon(secret, seed)
    const entField = poseidon([poseidon.F.e(secret), poseidon.F.e(seedBig)]);
    const entropySignal = poseidon.F.toObject(entField);

    // 5) write a one-off temp input.json
    const tmp = fs.mkdtempSync(path.join(__dirname, "zktemp-"));
    const inp = {
      secret: secret.toString(),
      seed: seedBig.toString(),
      entropy: entropySignal.toString(),
      commitment: initialCommitment.toString(),
    };
    const inPath = path.join(tmp, "input.json");
    fs.writeFileSync(inPath, JSON.stringify(inp));

    // 6) snarkjs fullprove + export calldata
    execSync(
      `snarkjs plonk fullprove ${inPath} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_js/vrf.wasm"
      )} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_final.zkey"
      )} ${tmp}/proof.json ${tmp}/public.json`,
      { stdio: "inherit" }
    );
    const calldata = execSync(
      `snarkjs zkey export soliditycalldata ${tmp}/public.json ${tmp}/proof.json --plonk`,
      { encoding: "utf8" }
    ).trim();

    // parse proof[] and publicSignals[]
    const [proofPart, publicPart] = calldata.split("][");
    const proof: string[] = JSON.parse(proofPart + "]");
    const publicSignals: string[] = JSON.parse("[" + publicPart);

    expect(proof.length).to.equal(24);
    expect(publicSignals.length).to.equal(3);

    proof[0] = "0x" + proof[0].slice(2).split("").reverse().join("");
    await expect(
      entropy.write.submitEntropy(
        [
          proof as any,
          publicSignals as any,
          requestId,
          fulfiller.account.address,
        ],
        {
          account: sender.account,
        }
      )
    ).to.be.rejectedWith("InvalidProof()");

    // cleanup
    fs.rmSync(tmp, { recursive: true, force: true });
  });
  it("should revert if the request has already been fulfilled", async function () {
    const { fulfiller, sender, entropy, publicClient, admin } =
      await loadFixture(deployEntropyFixture);

    // 1) pick a private secret, build initial commitment = Poseidon(secret, 0)
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    // push that to chain
    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });

    // 2) request a VRF
    const reqTx = await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
    });
    const reqReceipt = await publicClient.waitForTransactionReceipt({
      hash: reqTx,
    });
    expect(reqReceipt.status).to.equal("success");

    // Fetch block details using getBlock
    const block = await publicClient.getBlock({
      blockNumber: reqReceipt.blockNumber - 1n,
    });
    const blockNumber = BigInt(block.number);
    const blockHash = block.hash;

    // requestId will be 1 (first one)
    const requestId = BigInt(1);

    // 3) recompute the seed = keccak256(requestId, blockNumber) % BN128_PRIME
    const packed = encodeAbiParameters(
      [{ type: "uint256" }, { type: "uint256" }, { type: "bytes32" }],
      [requestId, blockNumber, blockHash]
    );
    const seedHash = keccak256(packed);
    const seedBig = BigInt(seedHash) % BN128_PRIME;

    // 4) recompute the entropy = Poseidon(secret, seed)
    const entField = poseidon([poseidon.F.e(secret), poseidon.F.e(seedBig)]);
    const entropySignal = poseidon.F.toObject(entField);

    // 5) write a one-off temp input.json
    const tmp = fs.mkdtempSync(path.join(__dirname, "zktemp-"));
    const inp = {
      secret: secret.toString(),
      seed: seedBig.toString(),
      entropy: entropySignal.toString(),
      commitment: initialCommitment.toString(),
    };
    const inPath = path.join(tmp, "input.json");
    fs.writeFileSync(inPath, JSON.stringify(inp));

    // 6) snarkjs fullprove + export calldata
    execSync(
      `snarkjs plonk fullprove ${inPath} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_js/vrf.wasm"
      )} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_final.zkey"
      )} ${tmp}/proof.json ${tmp}/public.json`,
      { stdio: "inherit" }
    );
    const calldata = execSync(
      `snarkjs zkey export soliditycalldata ${tmp}/public.json ${tmp}/proof.json --plonk`,
      { encoding: "utf8" }
    ).trim();

    // parse proof[] and publicSignals[]
    const [proofPart, publicPart] = calldata.split("][");
    const proof: string[] = JSON.parse(proofPart + "]");
    const publicSignals: string[] = JSON.parse("[" + publicPart);

    expect(proof.length).to.equal(24);
    expect(publicSignals.length).to.equal(3);

    await entropy.write.submitEntropy(
      [
        proof as any,
        publicSignals as any,
        requestId,
        fulfiller.account.address,
      ],
      {
        account: sender.account,
      }
    );

    await expect(
      entropy.write.submitEntropy(
        [
          proof as any,
          publicSignals as any,
          requestId,
          fulfiller.account.address,
        ],
        {
          account: sender.account,
        }
      )
    ).to.be.rejectedWith("RequestAlreadyFulfilled()");

    // cleanup
    fs.rmSync(tmp, { recursive: true, force: true });
  });
  it("should revert if the requestId is or is greater than the nextRequestId", async function () {
    const { fulfiller, sender, entropy, publicClient, admin } =
      await loadFixture(deployEntropyFixture);

    // 1) pick a private secret, build initial commitment = Poseidon(secret, 0)
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    // push that to chain
    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });

    // 2) request a VRF
    const reqTx = await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
    });
    const reqReceipt = await publicClient.waitForTransactionReceipt({
      hash: reqTx,
    });
    expect(reqReceipt.status).to.equal("success");

    // Fetch block details using getBlock
    const block = await publicClient.getBlock({
      blockNumber: reqReceipt.blockNumber - 1n,
    });
    const blockNumber = BigInt(block.number);
    const blockHash = block.hash;

    // requestId will be 1 (first one)
    const requestId = BigInt(2);

    // 3) recompute the seed = keccak256(requestId, blockNumber) % BN128_PRIME
    const packed = encodeAbiParameters(
      [{ type: "uint256" }, { type: "uint256" }, { type: "bytes32" }],
      [requestId, blockNumber, blockHash]
    );
    const seedHash = keccak256(packed);
    const seedBig = BigInt(seedHash) % BN128_PRIME;

    // 4) recompute the entropy = Poseidon(secret, seed)
    const entField = poseidon([poseidon.F.e(secret), poseidon.F.e(seedBig)]);
    const entropySignal = poseidon.F.toObject(entField);

    // 5) write a one-off temp input.json
    const tmp = fs.mkdtempSync(path.join(__dirname, "zktemp-"));
    const inp = {
      secret: secret.toString(),
      seed: seedBig.toString(),
      entropy: entropySignal.toString(),
      commitment: initialCommitment.toString(),
    };
    const inPath = path.join(tmp, "input.json");
    fs.writeFileSync(inPath, JSON.stringify(inp));

    // 6) snarkjs fullprove + export calldata
    execSync(
      `snarkjs plonk fullprove ${inPath} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_js/vrf.wasm"
      )} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_final.zkey"
      )} ${tmp}/proof.json ${tmp}/public.json`,
      { stdio: "inherit" }
    );
    const calldata = execSync(
      `snarkjs zkey export soliditycalldata ${tmp}/public.json ${tmp}/proof.json --plonk`,
      { encoding: "utf8" }
    ).trim();

    // parse proof[] and publicSignals[]
    const [proofPart, publicPart] = calldata.split("][");
    const proof: string[] = JSON.parse(proofPart + "]");
    const publicSignals: string[] = JSON.parse("[" + publicPart);

    expect(proof.length).to.equal(24);
    expect(publicSignals.length).to.equal(3);

    await expect(
      entropy.write.submitEntropy(
        [
          proof as any,
          publicSignals as any,
          requestId,
          fulfiller.account.address,
        ],
        {
          account: sender.account,
        }
      )
    ).to.be.rejectedWith("InvalidRequestId()");

    // cleanup
    fs.rmSync(tmp, { recursive: true, force: true });
  });
  it("should revert if the commitment from public input[2] does not match the current fulfiller commitment", async function () {
    const { fulfiller, sender, entropy, publicClient, admin } =
      await loadFixture(deployEntropyFixture);

    // 1) pick a private secret, build initial commitment = Poseidon(secret, 0)
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    // push that to chain
    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });

    // 2) request a VRF
    const reqTx = await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
    });
    const reqReceipt = await publicClient.waitForTransactionReceipt({
      hash: reqTx,
    });
    expect(reqReceipt.status).to.equal("success");

    // Fetch block details using getBlock
    const block = await publicClient.getBlock({
      blockNumber: reqReceipt.blockNumber - 1n,
    });
    const blockNumber = BigInt(block.number);
    const blockHash = block.hash;

    // requestId will be 1 (first one)
    const requestId = BigInt(1);

    // 3) recompute the seed = keccak256(requestId, blockNumber) % BN128_PRIME
    const packed = encodeAbiParameters(
      [{ type: "uint256" }, { type: "uint256" }, { type: "bytes32" }],
      [requestId, blockNumber, blockHash]
    );
    const seedHash = keccak256(packed);
    const seedBig = BigInt(seedHash) % BN128_PRIME;

    // 4) recompute the entropy = Poseidon(secret, seed)
    const entField = poseidon([poseidon.F.e(secret), poseidon.F.e(seedBig)]);
    const entropySignal = poseidon.F.toObject(entField);

    // 5) write a one-off temp input.json
    const tmp = fs.mkdtempSync(path.join(__dirname, "zktemp-"));
    const inp = {
      secret: secret.toString(),
      seed: seedBig.toString(),
      entropy: entropySignal.toString(),
      commitment: initialCommitment.toString(),
    };
    const inPath = path.join(tmp, "input.json");
    fs.writeFileSync(inPath, JSON.stringify(inp));

    // 6) snarkjs fullprove + export calldata
    execSync(
      `snarkjs plonk fullprove ${inPath} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_js/vrf.wasm"
      )} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_final.zkey"
      )} ${tmp}/proof.json ${tmp}/public.json`,
      { stdio: "inherit" }
    );
    const calldata = execSync(
      `snarkjs zkey export soliditycalldata ${tmp}/public.json ${tmp}/proof.json --plonk`,
      { encoding: "utf8" }
    ).trim();

    // parse proof[] and publicSignals[]
    const [proofPart, publicPart] = calldata.split("][");
    const proof: string[] = JSON.parse(proofPart + "]");
    const publicSignals: string[] = JSON.parse("[" + publicPart);

    expect(proof.length).to.equal(24);
    expect(publicSignals.length).to.equal(3);

    const fakePublicSignalCommitment = BigInt("0x" + "0".padStart(64, "0"));

    publicSignals[2] = "0x" + fakePublicSignalCommitment.toString(16);

    await expect(
      entropy.write.submitEntropy(
        [
          proof as any,
          publicSignals as any,
          requestId,
          fulfiller.account.address,
        ],
        {
          account: sender.account,
        }
      )
    ).to.be.rejectedWith("InvalidCommitment()");

    // cleanup
    fs.rmSync(tmp, { recursive: true, force: true });
  });
  it("should revert if the seed is invalid", async function () {
    const { fulfiller, sender, entropy, publicClient, admin } =
      await loadFixture(deployEntropyFixture);

    // 1) pick a private secret, build initial commitment = Poseidon(secret, 0)
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    // push that to chain
    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });

    // 2) request a VRF
    const reqTx = await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
    });
    const reqReceipt = await publicClient.waitForTransactionReceipt({
      hash: reqTx,
    });
    expect(reqReceipt.status).to.equal("success");

    // Fetch block details using getBlock
    const block = await publicClient.getBlock({
      blockNumber: reqReceipt.blockNumber - 1n,
    });
    const blockNumber = BigInt(block.number);
    const blockHash = block.hash;

    // requestId will be 1 (first one)
    const requestId = BigInt(1);

    // 3) recompute the seed = keccak256(requestId, blockNumber) % BN128_PRIME
    const packed = encodeAbiParameters(
      [{ type: "uint256" }, { type: "uint256" }, { type: "bytes32" }],
      [requestId, blockNumber + 1n, blockHash]
    );
    const seedHash = keccak256(packed);
    const seedBig = BigInt(seedHash) % BN128_PRIME;

    // 4) recompute the entropy = Poseidon(secret, seed)
    const entField = poseidon([poseidon.F.e(secret), poseidon.F.e(seedBig)]);
    const entropySignal = poseidon.F.toObject(entField);

    // 5) write a one-off temp input.json
    const tmp = fs.mkdtempSync(path.join(__dirname, "zktemp-"));
    const inp = {
      secret: secret.toString(),
      seed: seedBig.toString(),
      entropy: entropySignal.toString(),
      commitment: initialCommitment.toString(),
    };
    const inPath = path.join(tmp, "input.json");
    fs.writeFileSync(inPath, JSON.stringify(inp));

    // 6) snarkjs fullprove + export calldata
    execSync(
      `snarkjs plonk fullprove ${inPath} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_js/vrf.wasm"
      )} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_final.zkey"
      )} ${tmp}/proof.json ${tmp}/public.json`,
      { stdio: "inherit" }
    );
    const calldata = execSync(
      `snarkjs zkey export soliditycalldata ${tmp}/public.json ${tmp}/proof.json --plonk`,
      { encoding: "utf8" }
    ).trim();

    // parse proof[] and publicSignals[]
    const [proofPart, publicPart] = calldata.split("][");
    const proof: string[] = JSON.parse(proofPart + "]");
    const publicSignals: string[] = JSON.parse("[" + publicPart);

    expect(proof.length).to.equal(24);
    expect(publicSignals.length).to.equal(3);

    await expect(
      entropy.write.submitEntropy(
        [
          proof as any,
          publicSignals as any,
          requestId,
          fulfiller.account.address,
        ],
        {
          account: sender.account,
        }
      )
    ).to.be.rejectedWith("InvalidSeedOrBlockHashUnavailable()");

    // cleanup
    fs.rmSync(tmp, { recursive: true, force: true });
  });

  it("should revert if the commitment block number is greater than the request id block number, even with a valid proof", async function () {
    const { fulfiller, sender, entropy, publicClient, admin } =
      await loadFixture(deployEntropyFixture);

    // 1) pick a private secret, build initial commitment = Poseidon(secret, 0)
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");

    // 2) request a VRF
    const reqTx = await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
    });
    const reqReceipt = await publicClient.waitForTransactionReceipt({
      hash: reqTx,
    });
    expect(reqReceipt.status).to.equal("success");

    // requestId will be 1 (first one)
    const requestId = BigInt(1);

    await mine(1000);

    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });

    // 3) recompute the seed = keccak256(requestId, blockNumber) % BN128_PRIME
    const packed = encodeAbiParameters(
      [{ type: "uint256" }, { type: "uint256" }, { type: "bytes32" }],
      [requestId, reqReceipt.blockNumber - 1n, reqReceipt.blockHash]
    );
    const seedHash = keccak256(packed);
    const seedBig = BigInt(seedHash) % BN128_PRIME;

    // 4) recompute the entropy = Poseidon(secret, seed)
    const entField = poseidon([poseidon.F.e(secret), poseidon.F.e(seedBig)]);
    const entropySignal = poseidon.F.toObject(entField);

    // 5) write a one-off temp input.json
    const tmp = fs.mkdtempSync(path.join(__dirname, "zktemp-"));
    const inp = {
      secret: secret.toString(),
      seed: seedBig.toString(),
      entropy: entropySignal.toString(),
      commitment: initialCommitment.toString(),
    };
    const inPath = path.join(tmp, "input.json");
    fs.writeFileSync(inPath, JSON.stringify(inp));

    // 6) snarkjs fullprove + export calldata
    execSync(
      `snarkjs plonk fullprove ${inPath} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_js/vrf.wasm"
      )} ${path.resolve(
        __dirname,
        "../../fulfiller/zk/vrf_final.zkey"
      )} ${tmp}/proof.json ${tmp}/public.json`,
      { stdio: "inherit" }
    );
    const calldata = execSync(
      `snarkjs zkey export soliditycalldata ${tmp}/public.json ${tmp}/proof.json --plonk`,
      { encoding: "utf8" }
    ).trim();

    // parse proof[] and publicSignals[]
    const [proofPart, publicPart] = calldata.split("][");
    const proof: string[] = JSON.parse(proofPart + "]");
    const publicSignals: string[] = JSON.parse("[" + publicPart);

    expect(proof.length).to.equal(24);
    expect(publicSignals.length).to.equal(3);

    // 7) generate a _new_ secret+commitment for the next round
    const newSecret = BigInt(7777777);
    const newCommField = poseidon([poseidon.F.e(newSecret), zero]);
    const newCommitment = poseidon.F.toObject(newCommField);

    await expect(
      entropy.write.submitEntropy(
        [
          proof as any,
          publicSignals as any,
          requestId,
          fulfiller.account.address,
        ],
        {
          account: sender.account,
        }
      )
    ).to.be.rejectedWith("InvalidCommitmentBlock()");

    // cleanup
    fs.rmSync(tmp, { recursive: true, force: true });
  });

  it("should allow the admin to set the fee and the fee percentage but not any other address", async function () {
    const { admin, entropy, sender } = await loadFixture(deployEntropyFixture);

    const newFee = parseEther("0.00001");
    const newFeePercentage = 1000n;

    await entropy.write.setRequestFee([newFee], {
      account: admin.account,
    });

    await entropy.write.setContractFeeBasisPoints([newFeePercentage], {
      account: admin.account,
    });

    const fee = await entropy.read.requestFee();
    const feePercentage = await entropy.read.contractFeeBasisPoints();

    expect(fee).to.equal(newFee);
    expect(feePercentage).to.equal(newFeePercentage);

    const newFee2 = parseEther("0.00002");
    const newFeePercentage2 = 20n;
    await expect(
      entropy.write.setRequestFee([newFee2], {
        account: sender.account,
      })
    ).to.be.rejectedWith("AccessControlUnauthorizedAccount");
    await expect(
      entropy.write.setContractFeeBasisPoints([newFeePercentage2], {
        account: sender.account,
      })
    ).to.be.rejectedWith("AccessControlUnauthorizedAccount");
  });
  it("should revert on a withdraw of 0 from a fulfiller", async function () {
    const { fulfiller, entropy } = await loadFixture(deployEntropyFixture);
    await expect(
      entropy.write.withdrawRewardReceiverBalance({
        account: fulfiller.account,
      })
    ).to.be.rejectedWith("InsufficientBalance()");
  });
  it("should revert if non admin tries to withdraw contract fees", async function () {
    const { admin, sender, entropy } = await loadFixture(deployEntropyFixture);
    await expect(
      entropy.write.withdrawContractFees({
        account: sender.account,
      })
    ).to.be.rejectedWith("AccessControlUnauthorizedAccount");
  });
  it("should allow the requester to be refunded of the full amount if the request block hash is no longer available", async function () {
    const { entropy, sender, admin } = await loadFixture(deployEntropyFixture);

    // 2) request a VRF
    await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
      account: sender.account,
    });
    await mine(256);

    expect(
      await entropy.write.refundUnfulfilledRequest([1n], {
        account: sender.account,
      })
    ).to.be.ok;
  });
  it("should only allow the requester to be refunded", async function () {
    const { entropy, sender, admin } = await loadFixture(deployEntropyFixture);

    // 2) request a VRF
    await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
      account: sender.account,
    });
    await mine(1000);

    await expect(
      entropy.write.refundUnfulfilledRequest([1n], {
        account: admin.account,
      })
    ).to.be.rejectedWith("InvalidRequester()");
  });
  it("should not allow a refund of a non-existent request", async function () {
    const { entropy, sender, admin } = await loadFixture(deployEntropyFixture);

    // 2) request a VRF
    await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
      account: sender.account,
    });
    await mine(1000);

    await expect(
      entropy.write.refundUnfulfilledRequest([2n], {
        account: sender.account,
      })
    ).to.be.rejectedWith("InvalidRequester()");
  });
  it("should not allow the refund of a request if the block hash is available", async function () {
    const { entropy, sender, admin } = await loadFixture(deployEntropyFixture);

    // 2) request a VRF
    await entropy.write.requestRng([zeroAddress, 0], {
      value: parseEther("0.000005"),
      account: sender.account,
    });
    await mine(200);

    await expect(
      entropy.write.refundUnfulfilledRequest([1n], {
        account: sender.account,
      })
    ).to.be.rejectedWith("RequestStillValid()");
  });
  it("should not allow the admin to set the fee to 20% or more", async function () {
    const { admin, entropy } = await loadFixture(deployEntropyFixture);

    const newFeePercentage = 2100n;

    await expect(
      entropy.write.setContractFeeBasisPoints([newFeePercentage], {
        account: admin.account,
      })
    ).to.be.rejectedWith("InvalidFeeBasisPoints()");
  });
  it("does not allow two fulfillers to use the same commitment", async function () {
    const { fulfiller, sender, entropy } =
      await loadFixture(deployEntropyFixture);
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");
    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });
    await expect(
      entropy.write.setCommitment([initialCommHex as any], {
        account: fulfiller.account,
      })
    ).to.be.rejectedWith("CommitmentInUse()");
  });
  it("unsets the previous commitment when changing commitments", async function () {
    const { sender, entropy } = await loadFixture(deployEntropyFixture);
    const secret = BigInt(123456);
    const zero = poseidon.F.e(0);
    const commField = poseidon([poseidon.F.e(secret), zero]);
    const initialCommitment = poseidon.F.toObject(commField);
    const initialCommHex =
      "0x" + initialCommitment.toString(16).padStart(64, "0");
    await entropy.write.setCommitment([initialCommHex as any], {
      account: sender.account,
    });
    const newSecret = BigInt(7777777);
    const newCommField = poseidon([poseidon.F.e(newSecret), zero]);
    const newCommitment = poseidon.F.toObject(newCommField);
    const newCommHex = "0x" + newCommitment.toString(16).padStart(64, "0");
    await entropy.write.setCommitment([newCommHex as any], {
      account: sender.account,
    });
    expect(await entropy.read.commitmentInUse([initialCommHex as any])).to.be
      .false;
  });

  it("should reject invalid withdrawal attempts", async function () {
    const { sender, entropy } = await loadFixture(deployEntropyFixture);

    await expect(
      entropy.write.withdrawRewardReceiverBalance({ account: sender.account })
    ).to.be.rejectedWith("InsufficientBalance()");
  });
});
