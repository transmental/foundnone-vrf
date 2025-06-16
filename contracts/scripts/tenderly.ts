import { tenderly } from "hardhat";

async function main() {
  console.log("Verifying");

  await tenderly.verify({
    name: "FoundnoneVRF",
    address: "0x599613443B57dF5FB64C3C0e5bab72311f9aB040",
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
