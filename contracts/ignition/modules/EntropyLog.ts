import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const EntropyLogModule = buildModule("EntropyLogModule", (m) => {
  const admin = m.getParameter("admin", "0x230988f9ab3019F100Ad6bb50723bE27dd9783E3");

  const entropyLog = m.contract("FoundnoneVRF", [admin]);

  return { entropyLog };
});

export default EntropyLogModule;
