var FeralfileExhibitionV4_1 = artifacts.require("FeralfileExhibitionV4_1");

const argv = require("minimist")(process.argv.slice(2), {
    string: [
        "exhibition_signer",
        "exhibition_vault",
        "exhibition_cost_receiver",
    ],
});

module.exports = function (deployer) {
    let exhibition_name = argv.exhibition_name || "FFV4 Test";
    let exhibition_symbol = argv.exhibition_symbol || "FeralfileExhibitionV4";
    let exhibition_signer =
        argv.exhibition_signer || "0x6732389c6d47d01487dcDc96e2Cc6BAf108452f2";
    let exhibition_vault =
        argv.exhibition_vault || "0x0c51e8becb17ba3203cd04d3fc31fcb90de412a1";
    let exhibition_cost_receiver =
        argv.exhibition_cost_receiver ||
        "0x6732389c6d47d01487dcDc96e2Cc6BAf108452f2";
    let burnable = argv.burnable || true;
    let bridgeable = argv.bridgeable || true;
    let contract_uri =
        argv.contract_uri ||
        "ipfs://QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe";
    let series_ids = argv.series_ids || [1, 2, 3, 4];
    let max_supplies = argv.max_supplies || [1000, 1000, 1000, 1000];

    deployer.deploy(
        FeralfileExhibitionV4_1,
        exhibition_name,
        exhibition_symbol,
        burnable,
        bridgeable,
        exhibition_signer,
        exhibition_vault,
        exhibition_cost_receiver,
        contract_uri,
        series_ids,
        max_supplies
    );
};
