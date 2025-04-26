import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import hre from "hardhat";
import { expect } from "chai";

describe("PlonkVerifier standalone", () => {
  async function deployVerifier() {
    // deploy only the verifier (no EntropyLog wrapper)
    const verifier = await hre.viem.deployContract("PlonkVerifier", []);
    const publicClient = await hre.viem.getPublicClient();
    return { verifier, publicClient };
  }

  it("callStatic.verifyProof returns true for a valid proof", async () => {
    const { verifier, publicClient } = await loadFixture(deployVerifier);
    // static call the verifier directly
    const ok = await verifier.read.verifyProof([
      [
        "0x280c08e52dd258a1d313f7302c33785b4ccc19cca4d4e383b2e8a84ff1b06a0f",
        "0x21ca4ef0974dd7841bce0a8ce2dbba401f3c160ed90e86c1dc79968e14141938",
        "0x0a6c94c9855ece378b9f1692d55b3bff90ccf378cd26cb3df9175238257a8ed1",
        "0x2578121e0cc871d662c838c7aef2dd15f61036a5d8309ac6d7d03d0c7daee526",
        "0x01f3a42c8646e59cac0c4d85700070cf930f37b8f30695c75df148a543cfda81",
        "0x138a3a11cce346e1294c369fcc6d508bc15be59c15ff41a37f5a5d628a29a2c4",
        "0x250c7df7ff046aa424ab644d1e420b2bb3f1e292b598c8d2b327db9204c98dd2",
        "0x1aacd968e70860bfab6a7d9673b4a119f567df944e9fc63ce84ef4588f4b15de",
        "0x0a09ad6846ad330ba58a28fa5840021fa0c7d3cda7d79999db377c2ca2dc3957",
        "0x0a2374e5d5967808da73f500ca402bda2d4b7124a64c6750c367058706789049",
        "0x0bed69e08a4aa23f7f796acba3df3e832ca6ff34c8b1f89a1b0a8c9a4356d534",
        "0x0a5d01888815b82241d76f78a463150c0bfbe740f3acbdbf97fe61fc7df979df",
        "0x07f3181b247b0ca69ec6c97519f626680eaaff4fb0a3f5e8b1159e245490ad22",
        "0x25f754fb7b166ac38005696e4265df73edd9ac03228f44e54734b9c2c30837d1",
        "0x1e3217fef26d8768fabde5ed2392ffc12e4b289b5323b9d93436bbbf5648ea59",
        "0x168437dd24ce177e6fcb27c1d2286d4145c4213183bb94f3148981fe81c783ed",
        "0x22d7d99e10ccc50281850d6e0ce9b4e9c730c7e0b888ab04b262e84f18e19641",
        "0x13c16671e60d309f0124778f1226df7c7423b2ad96b3df341e7b82148fed20e8",
        "0x22d2492393d6ec46922cbe165a183fbbe4cf3b6019013f32b8924653231e2524",
        "0x2b7abf738fe244d52c9e23e4750e6f172fae06544653f465c422632276e01f95",
        "0x0444dd39356d07402bc7796bac587d8c14969782ada9ebf76dd8fb58f00dea37",
        "0x23aa48eb5cbcda1c31f3d2b8f682e9f7933a27acc25a90796a8a71b4e05132ab",
        "0x0060674b0d7aac36e97ffdddd1976ca16a145f5a7bf88be4835d92310198bfaf",
        "0x0e012a39d0a9fa4141edea6c9a13eccef0073f208ec2373ab591514ce85b7421",
      ],
      [
        "0x000e71fa66637cba544f5832f2372e08a14d95dda4c7278d28c6f40a9036b1fc",
        "0x139ffc1d7097500c28f4e6827729a2e06014c5f5adb2df2d6ec9f6e0c025893a",
        "0x14da730266710e01739ca813057174e0a3a6966168a7cc6ad0011be371c0b836",
      ],
    ] as any);

    expect(ok).to.be.true;
  });

  it("callStatic.verifyProof returns false for a bad proof", async () => {
    const { verifier } = await loadFixture(deployVerifier);

    const ok = await verifier.read.verifyProof([
      [
        "0x280c08e52dd258a1d313f7302c33785b4ccc19cca4d4e383b2e8a84ff1b06a0f",
        "0x21ca4ef0974dd7841bce0a8ce2dbba401f3c160ed90e86c1dc79968e14141938",
        "0x0a6c94c9855ece378b9f1692d55b3bff90ccf378cd26cb3df9175238257a8ed1",
        "0x2578121e0cc871d662c838c7aef2dd15f61036a5d8309ac6d7d03d0c7daee526",
        "0x01f3a42c8646e59cac0c4d85700070cf930f37b8f30695c75df148a543cfda81",
        "0x138a3a11cce346e1294c369fcc6d508bc15be59c15ff41a37f5a5d628a29a2c4",
        "0x250c7df7ff046aa424ab644d1e420b2bb3f1e292b598c8d2b327db9204c98dd2",
        "0x1aacd968e70860bfab6a7d9673b4a119f567df944e9fc63ce84ef4588f4b15de",
        "0x0a09ad6846ad330ba58a28fa5840021fa0c7d3cda7d79999db377c2ca2dc3957",
        "0x0a2374e5d5967808da73f500ca402bda2d4b7124a64c6750c367058706789049",
        "0x0bed69e08a4aa23f7f796acba3df3e832ca6ff34c8b1f89a1b0a8c9a4356d534",
        "0x0a5d01888815b82241d76f78a463150c0bfbe740f3acbdbf97fe61fc7df979df",
        "0x07f3181b247b0ca69ec6c97519f626680eaaff4fb0a3f5e8b1159e245490ad22",
        "0x25f754fb7b166ac38005696e4265df73edd9ac03228f44e54734b9c2c30837d1",
        "0x1e3217fef26d8768fabde5ed2392ffc12e4b289b5323b9d93436bbbf5648ea59",
        "0x168437dd24ce177e6fcb27c1d2286d4145c4213183bb94f3148981fe81c783ed",
        "0x22d7d99e10ccc50281850d6e0ce9b4e9c730c7e0b888ab04b262e84f18e19641",
        "0x13c16671e60d309f0124778f1226df7c7423b2ad96b3df341e7b82148fed20e8",
        "0x22d2492393d6ec46922cbe165a183fbbe4cf3b6019013f32b8924653231e2524",
        "0x2b7abf738fe244d52c9e23e4750e6f172fae06544653f465c422632276e01f95",
        "0x0444dd39356d07402bc7796bac587d8c14969782ada9ebf76dd8fb58f00dea37",
        "0x23aa48eb5cbcda1c31f3d2b8f682e9f7933a27acc25a90796a8a71b4e05132ab",
        "0x0060674b0d7aac36e97ffdddd1976ca16a145f5a7bf88be4835d92310198bfaf",
        "0x0e012a39d0a9fa4141edea6c9a13eccef0073f208ec2373ab591514ce85b7421",
      ],
      [
        "0x000e71fa66637cba544f5832f2372e08a14d95dda4c7278d28c6f40a9036b1fc",
        "0x139ffc1d7097500c28f4e6827729a2e06014c5f5adb2df2d6ec9f6e0c025893a",
        "0x04da730266710e01739ca813057174e0a3a6966168a7cc6ad0011be371c0b836",
      ],
    ] as any);

    expect(ok).to.be.false;
  });
});
