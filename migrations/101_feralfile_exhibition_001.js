var FeralfileExhibitionV2 = artifacts.require("FeralfileExhibitionV2");

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator']
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || "Feral File — Social Codes";
  let exhibition_symbol = argv.exhibition_symbol || "FF001";
  let exhibition_royalty_payout_address = argv.exhibition_royalty_payout_address || "0xf6d099037Df564a32237705e5005F4450b0EB6c3";
  let exhibition_secondary_sale_royalty_bps = argv.exhibition_secondary_sale_royalty_bps || 1000;
  let exhibition_edition_size = argv.exhibition_edition_size || 75;

  deployer.deploy(FeralfileExhibitionV2, exhibition_name, exhibition_symbol,
    exhibition_edition_size, exhibition_secondary_sale_royalty_bps, exhibition_royalty_payout_address,
    "https://ipfs.bitmark.com/ipfs/QmdcgZ8Uav5v99KCL5KjxZvSyYHyhLXE3CLGDfmhTtYpLX",
    "https://ipfs.bitmark.com/ipfs/");
};
