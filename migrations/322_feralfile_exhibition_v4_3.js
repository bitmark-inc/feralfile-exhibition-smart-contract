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
        argv.exhibition_vault || "0x455464F0d369dAC13002e81e9fAB857f6aD21795";
    let exhibition_cost_receiver =
        argv.exhibition_cost_receiver ||
        "0x080FEB125bA730D6D12789B6AAAB01f4E31D8Bd1";
    let burnable = argv.burnable || true;
    let bridgeable = argv.bridgeable || true;
    let contract_uri =
        argv.contract_uri ||
        "ipfs://QmUWB3RBh42qtT4aRRGiZCrvx98LPq8GmDmioE8EQi5Vm6";
    let series_ids = argv.series_ids || [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
    let max_supplies = argv.max_supplies || [
        512, 256, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    ];
    let merge_artwork_info = argv.merge_artwork_info || {
        singleSeriesId: 1,
        mergedSeriesId: 2,
        nextTokenId:
            "27271357221255084324537793443996264744604553267670416617055471525879030383744",
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
