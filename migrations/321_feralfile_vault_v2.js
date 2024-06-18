var FeralfileVaultV2 = artifacts.require("FeralfileVaultV2");

const argv = require("minimist")(process.argv.slice(2), {
    string: ["exhibition_signer"],
});

module.exports = function (deployer) {
    let exhibition_signer =
        argv.exhibition_signer || "0xBEb9F810862c40A144925f568b1853d72Acc492F";
    deployer.deploy(FeralfileVaultV2, exhibition_signer);
};
