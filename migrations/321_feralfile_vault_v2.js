var FeralfileVault = artifacts.require("FeralfileVaultV2");

const argv = require("minimist")(process.argv.slice(2), {
    string: ["exhibition_signer"],
});

module.exports = function (deployer) {
    let exhibition_signer =
        argv.exhibition_signer || "0x6732389c6d47d01487dcDc96e2Cc6BAf108452f2";
    deployer.deploy(FeralfileVault, exhibition_signer);
};
