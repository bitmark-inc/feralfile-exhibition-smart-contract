var FeralfileExhibitionV3_2 = artifacts.require('FeralfileExhibitionV3_2');

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator'],
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || 'FFV3_3 - Test 1';
  let exhibition_symbol = argv.exhibition_symbol || 'FFDV3_3';
  let exhibition_royalty_payout_address =
    argv.exhibition_royalty_payout_address ||
    '0x2033606bE146405870F92Ea3144ef5057b9DEA48';
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
    'https://ipfs.bitmark.com/ipfs/QmU1tkyTQ9jb6RaMZMVLZh4W4q9cWeJ3NiNPNw9ruSps6j',
    'https://ipfs.bitmark.com/ipfs/',
    burnable,
    bridgeable
  );
};
