var FeralfileExhibition = artifacts.require("FeralfileExhibition");

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator']
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || "Social Codes";
  let exhibition_symbol = argv.exhibition_symbol || "SSC";
  let exhibition_curator = argv.exhibition_curator || "0x8fd310de32848798eB64Bd88f9C5656Eea32415e";
  let exhibition_secondary_sale_royalty_bps = argv.exhibition_secondary_sale_royalty_bps || 1000;
  let exhibition_edition_size = argv.exhibition_edition_size || 50;
  let exhibition_initial_price = argv.exhibition_initial_price || 10;

  deployer.deploy(FeralfileExhibition, exhibition_name, exhibition_symbol,
    exhibition_curator.toString(), exhibition_edition_size, exhibition_initial_price, exhibition_secondary_sale_royalty_bps,
    "https://ipfs.bitmark.com/ipfs/QmUn95tx6p9XeJUfNyoAbqVRmPaWM3UUzbvJNvXfcmEpVb",
    "https://ipfs.bitmark.com/ipfs/");
};
