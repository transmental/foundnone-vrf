import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const MockEntropyReceiverModule = buildModule(
  "MockEntropyReceiverModule",
  (m) => {
    const vrfAddress = "0x25f27467377DaC26B79784603A0b2DcDaa3b67cf";

    const entropyReceiver = m.contract("MockEntropyReceiver", [vrfAddress]);

    return { entropyReceiver };
  }
);

export default MockEntropyReceiverModule;
