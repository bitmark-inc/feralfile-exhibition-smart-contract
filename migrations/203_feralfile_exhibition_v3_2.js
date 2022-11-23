var FeralfileExhibitionV3_2 = artifacts.require('FeralfileExhibitionV3_2');

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator'],
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || 'Feral File V3_2 Test';
  let exhibition_symbol = argv.exhibition_symbol || 'FFDV3_2';
  let exhibition_royalty_payout_address =
    argv.exhibition_royalty_payout_address ||
    '0xa704AA575172F84D917CcaF5EC775Da8227Ea6c3';
  let exhibition_secondary_sale_royalty_bps =
    argv.exhibition_secondary_sale_royalty_bps || 1000;
  let burnable = argv.burnable || true;
  let bridgeable = argv.bridgeable || true;

  deployer.deploy(
    FeralfileExhibitionV3_2,
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
