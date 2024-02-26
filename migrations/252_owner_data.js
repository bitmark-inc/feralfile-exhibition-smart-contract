var OwnerData = artifacts.require("OwnerData");
const argv = require("minimist")(process.argv.slice(2), {
    string: ["trustee"],
});

module.exports = function (deployer) {
    const trustee =
        argv.trustee || "0xdB33365a8730de2F7574ff1189fB9D337bF4c36d";
    deployer.deploy(OwnerData, trustee);
};
