var FeralfileExhibitionV31 = artifacts.require('FeralfileExhibitionV31');

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator'],
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || 'Feral File V3';
  let exhibition_symbol = argv.exhibition_symbol || 'FFDV3002';
  let exhibition_royalty_payout_address =
    argv.exhibition_royalty_payout_address ||
    '0x2760869A50D48F1C67253C4461c0A6f9e1440Cac';
  let exhibition_secondary_sale_royalty_bps =
    argv.exhibition_secondary_sale_royalty_bps || 1000;
  let burnable = argv.burnable || true;
  let bridgeable = argv.bridgeable || true;

  deployer.deploy(
    FeralfileExhibitionV31,
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
