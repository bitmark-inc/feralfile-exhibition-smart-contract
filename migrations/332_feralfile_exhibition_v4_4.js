var FeralfileExhibitionV4_4 = artifacts.require("FeralfileExhibitionV4_4");
var RendererStorageV0 = artifacts.require("RendererStorageV0");
var LibBytes = artifacts.require("LibBytes");

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
        argv.exhibition_signer || "0xdB33365a8730de2F7574ff1189fB9D337bF4c36d";
    let exhibition_vault =
        argv.exhibition_vault || "0x1AB40A882AE5373bAD177470951ecB5F17cDF090";
    let exhibition_cost_receiver =
        argv.exhibition_cost_receiver ||
        "0xdB33365a8730de2F7574ff1189fB9D337bF4c36d";
    let burnable = argv.burnable || true;
    let bridgeable = argv.bridgeable || true;
    let contract_uri =
        argv.contract_uri ||
        "ipfs://QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe";
    let series_ids = argv.series_ids || [1];
    let max_supplies = argv.max_supplies || [10];

    deployer.deploy(LibBytes);
    deployer.link(LibBytes, FeralfileExhibitionV4_4);
    deployer.deploy(RendererStorageV0);
    deployer.link(RendererStorageV0, FeralfileExhibitionV4_4);

    deployer.deploy(
        FeralfileExhibitionV4_4,
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
