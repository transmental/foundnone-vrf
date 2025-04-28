pragma circom 2.1.4;
include "../node_modules/circomlib/circuits/poseidon.circom";

template GenProof() {
    // **private**
    signal input secret;

    // **public**
    signal input seed;
    signal input entropy;
    signal input commitment;

    component h1 = Poseidon(2);
    h1.inputs[0] <== secret;
    h1.inputs[1] <== seed;
    h1.out      === entropy;

    signal zero <== 0;
    component h2 = Poseidon(2);
    h2.inputs[0] <== secret;
    h2.inputs[1] <== zero;
    h2.out      === commitment;
}

component main { public [ seed, entropy, commitment ] } = GenProof();
