var FeralfileExhibitionV4_5 = artifacts.require("FeralfileExhibitionV4_5");

const argv = require("minimist")(process.argv.slice(2), {
    string: [
        "exhibition_signer",
        "exhibition_vault",
        "exhibition_cost_receiver",
        "exhibition_token_id_prefix_shard",
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
    let exhibition_tokenId_prefix_shard =
        argv.exhibition_token_id_prefix_shard ||
        "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6";
    let burnable = argv.burnable || true;
    let bridgeable = argv.bridgeable || true;
    let contract_uri =
        argv.contract_uri ||
        "ipfs://QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe";
    let series_ids = argv.series_ids || [1, 2, 3, 4];
    let max_supplies = argv.max_supplies || [1000, 1000, 1000, 1000];

    deployer.deploy(
        FeralfileExhibitionV4_5,
        exhibition_name,
        exhibition_symbol,
        burnable,
        bridgeable,
        exhibition_signer,
        exhibition_vault,
        exhibition_cost_receiver,
        contract_uri,
        series_ids,
        max_supplies,
        exhibition_tokenId_prefix_shard
    );
};
