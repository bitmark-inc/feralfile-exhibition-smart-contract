const FeralFileAirdropV1 = artifacts.require("FeralFileAirdropV1");

const argv = require("minimist")(process.argv.slice(2), {
    string: ["token_uri", "type"],
});

module.exports = function (deployer) {
    let token_uri =
        argv.token_uri ||
        "https://ipfs.bitmark.com/ipfs/QmNVdQSp1AvZonLwHzTbbZDPLgbpty15RMQrbPEWd4ooTU/{id}";
    let type = argv.type || 0;
    let burnable = argv.burnable || true;
    let bridgeable = argv.bridgeable || true;

    deployer.deploy(FeralFileAirdropV1, type, token_uri, burnable, bridgeable);
};
