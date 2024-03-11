var FeralfileExhibitionV5 = artifacts.require("FeralfileExhibitionV5");

const argv = require("minimist")(process.argv.slice(2), {
    string: [
        "exhibition_signer",
        "exhibition_vault",
        "exhibition_cost_receiver",
    ],
});

module.exports = function (deployer) {
    let exhibition_signer =
        argv.exhibition_signer || "0xdB33365a8730de2F7574ff1189fB9D337bF4c36d";
    let exhibition_vault =
        argv.exhibition_vault || "0x0c51e8becb17ba3203cd04d3fc31fcb90de412a1";
    let exhibition_cost_receiver =
        argv.exhibition_cost_receiver ||
        "0xdB33365a8730de2F7574ff1189fB9D337bF4c36d";
    let burnable = argv.burnable || true;
    let contract_uri =
        argv.contract_uri ||
        "ipfs://QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe";
    let token_uri =
        argv.token_uri ||
        "ipfs://QmZt5r6bU7r3BvCp6b6Z1yYK5Hd9M9WnQV2g8VfNnMmM3T/{id}";
    let series_ids = argv.series_ids || [1, 2, 3, 4];
    let series_max_supplies = argv.series_max_supplies || [10, 10, 10, 1];
    let series_artwork_max_supplies = argv.series_artwork_max_supplies || [
        1, 1, 1, 100,
    ];

    deployer.deploy(
        FeralfileExhibitionV5,
        token_uri,
        contract_uri,
        exhibition_signer,
        exhibition_vault,
        exhibition_cost_receiver,
        burnable,
        series_ids,
        series_max_supplies,
        series_artwork_max_supplies
    );
};
