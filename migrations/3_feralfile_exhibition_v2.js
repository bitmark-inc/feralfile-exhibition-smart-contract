var FeralfileExhibitionV2 = artifacts.require("FeralfileExhibitionV2");

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator']
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || "Feral File — Feral File Test Exhibition";
  let exhibition_symbol = argv.exhibition_symbol || "FFDV2001";
  let exhibition_royalty_payout_address = argv.exhibition_royalty_payout_address || "0x2760869A50D48F1C67253C4461c0A6f9e1440Cac";
  let exhibition_secondary_sale_royalty_bps = argv.exhibition_secondary_sale_royalty_bps || 1000;
  let exhibition_edition_size = argv.exhibition_edition_size || 50;

  deployer.deploy(FeralfileExhibitionV2, exhibition_name, exhibition_symbol,
    exhibition_edition_size, exhibition_secondary_sale_royalty_bps, exhibition_royalty_payout_address,
    "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
    "https://ipfs.bitmark.com/ipfs/");
};
