var FeralfileExhibitionV4_2 = artifacts.require("FeralfileExhibitionV4_2");

const argv = require("minimist")(process.argv.slice(2), {
    string: [
        "exhibition_signer",
        "exhibition_vault",
        "exhibition_cost_receiver",
    ],
});

module.exports = function (deployer) {
    let exhibition_name = argv.exhibition_name || "crystalline work";
    let exhibition_symbol = argv.exhibition_symbol || "FERALFILE";
    let exhibition_signer =
        argv.exhibition_signer || "0xBEb9F810862c40A144925f568b1853d72Acc492F";
    let exhibition_vault =
        argv.exhibition_vault || "0xcBFaf4BDE69C9b37835761E5228f9fe9E25b452f";
    let exhibition_cost_receiver =
        argv.exhibition_cost_receiver ||
        "0x080FEB125bA730D6D12789B6AAAB01f4E31D8Bd1";
    let burnable = argv.burnable || true;
    let bridgeable = argv.bridgeable || true;
    let contract_uri =
        argv.contract_uri ||
        "ipfs://QmZuygbgeVDZ8NBBpD3oSVUAibTTgKYnKvYWikdGp7HwNb";
    let series_ids = argv.series_ids || [1];
    let max_supplies = argv.max_supplies || [9048];
    let next_purchasable_tokens = argv.next_purchasable_tokens || [
        "31946296525744824328753280828797417080692251952513183340713878305007073182561",
    ];

    deployer.deploy(
        FeralfileExhibitionV4_2,
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
        next_purchasable_tokens
    );
};
