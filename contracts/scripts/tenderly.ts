import { tenderly } from "hardhat";

async function main() {
  console.log("Verifying");

  await tenderly.verify({
    name: "FoundnoneVRF",
    address: "0x25f27467377DaC26B79784603A0b2DcDaa3b67cf",
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
