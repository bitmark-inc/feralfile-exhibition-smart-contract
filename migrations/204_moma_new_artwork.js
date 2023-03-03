var FeralfileExhibitionV3_3 = artifacts.require('FeralfileExhibitionV3_3');

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator'],
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || 'MoMA Burn';
  let exhibition_symbol = argv.exhibition_symbol || 'FeralfileExhibitionV3_3';
  let exhibition_royalty_payout_address =
    argv.exhibition_royalty_payout_address ||
    '0x487bA00d91015dcc905Bb93b528c12a05FbC7A4F';
  let exhibition_secondary_sale_royalty_bps =
    argv.exhibition_secondary_sale_royalty_bps || 1000;
  let burnable = argv.burnable || true;
  let bridgeable = argv.bridgeable || true;

  deployer.deploy(
    FeralfileExhibitionV3_3,
    exhibition_name,
    exhibition_symbol,
    exhibition_secondary_sale_royalty_bps,
    exhibition_royalty_payout_address,
    'https://ipfs.bitmark.com/ipfs/QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe',
    'https://ipfs.bitmark.com/ipfs/',
    burnable,
    bridgeable
  );
};
