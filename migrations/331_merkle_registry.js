let MerkleRegistry = artifacts.require("MerkleRegistry");

const argv = require("minimist")(process.argv.slice(2), {
    string: ["source_code_uri"],
});

module.exports = function (deployer) {
    let sourceCodeURI =
        argv.source_code_uri ||
        "ipfs://QmRRfMJyncYMdNmSUk1DBTzYXHqh5Xmdgw6ijeFwdRh3Vj";

    deployer.deploy(MerkleRegistry, sourceCodeURI);
};
