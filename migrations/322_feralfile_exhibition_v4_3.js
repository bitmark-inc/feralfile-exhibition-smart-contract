var FeralfileExhibitionV4_3 = artifacts.require("FeralfileExhibitionV4_3");

const argv = require("minimist")(process.argv.slice(2), {
    string: [
        "exhibition_signer",
        "exhibition_vault",
        "exhibition_cost_receiver",
    ],
});

module.exports = function (deployer) {
    let exhibition_name = argv.exhibition_name || "CRAWL";
    let exhibition_symbol = argv.exhibition_symbol || "FERALFILE";
    let exhibition_signer =
        argv.exhibition_signer || "0xBEb9F810862c40A144925f568b1853d72Acc492F";
    let exhibition_vault =
        argv.exhibition_vault || "0x0c51e8becb17ba3203cd04d3fc31fcb90de412a1";
    let exhibition_cost_receiver =
        argv.exhibition_cost_receiver ||
        "0x6732389c6d47d01487dcDc96e2Cc6BAf108452f2";
    let burnable = argv.burnable || true;
    let bridgeable = argv.bridgeable || true;
    let contract_uri =
        argv.contract_uri ||
        "ipfs://QmZuygbgeVDZ8NBBpD3oSVUAibTTgKYnKvYWikdGp7HwNb";
        let series_ids = argv.series_ids || [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
        let max_supplies = argv.max_supplies || [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 512, 511];
    let merge_artwork_info = argv.merge_artwork_info || {
        singleSeriesId: 11,
        mergedSeriesId: 12,
        nextTokenId: "65139289060393321767190134082157117717885840752315574009980337524583053413184"
      };

    deployer.deploy(
        FeralfileExhibitionV4_3,
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
        merge_artwork_info
    );
};
