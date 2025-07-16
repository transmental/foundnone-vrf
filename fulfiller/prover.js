const express = require("express");
const { plonk } = require("snarkjs");
const fs = require("fs");
const path = require("path");

const app = express();
app.use(express.json());

const wasm = fs.readFileSync(path.join(__dirname, "zk/vrf_js/vrf.wasm"));
const zkey = fs.readFileSync(path.join(__dirname, "zk/vrf_final.zkey"));

app.post("/prove", async (req, res) => {
  const startTime = Date.now();
  try {
    const input = req.body;
    console.log("Received input for proof generation:", input);
    const { proof, publicSignals } = await plonk.fullProve(input, wasm, zkey);

    const rawCalldata = await plonk.exportSolidityCallData(
      proof,
      publicSignals
    );

    // Split the rawCalldata into two arrays
    const splitIndex = rawCalldata.indexOf("][");
    if (splitIndex === -1) {
      throw new Error("Unexpected calldata format");
    }

    const proofPart = rawCalldata.substring(0, splitIndex + 1);
    const publicSignalsPart = rawCalldata.substring(splitIndex + 1);

    const proofArray = JSON.parse(proofPart);
    const publicSignalsArray = JSON.parse(publicSignalsPart);

    // Concatenate the two arrays
    const calldata = proofArray.concat(publicSignalsArray);

    res.json({ calldata });
  } catch (error) {
    console.error("Error generating proof:", error);
    res.status(500).json({ error: "Failed to generate proof" });
  } finally {
    const endTime = Date.now();
    console.log("Proof generation time:", endTime - startTime, "ms");
  }
});

app.listen(3001, () => console.log("Snarkjs prover listening on port 3001"));
