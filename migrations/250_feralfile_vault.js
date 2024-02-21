var FeralfileVault = artifacts.require("FeralfileVault");

const argv = require("minimist")(process.argv.slice(2), {
    string: ["exhibition_signer"],
});

module.exports = function (deployer) {
    let exhibition_signer =
        argv.exhibition_signer || "0xdB33365a8730de2F7574ff1189fB9D337bF4c36d";
    deployer.deploy(FeralfileVault, exhibition_signer);
};
