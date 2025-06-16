import { tenderly } from "hardhat";

async function main() {
  console.log("Verifying");

  await tenderly.verify({
    name: "FoundnoneVRF",
    address: "0x857655697c706a774c47a72Ca93018Fb139C9e4a",
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
