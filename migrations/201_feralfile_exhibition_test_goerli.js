var FeralfileExhibitionV3 = artifacts.require('FeralfileExhibitionV3');

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator'],
});

module.exports = function (deployer) {
  let exhibition_name = argv.exhibition_name || 'FFV3 - Test 1';
  let exhibition_symbol = argv.exhibition_symbol || 'FFV3Test1';
  let exhibition_royalty_payout_address =
    argv.exhibition_royalty_payout_address ||
    '0x2033606bE146405870F92Ea3144ef5057b9DEA48';
  let exhibition_secondary_sale_royalty_bps =
    argv.exhibition_secondary_sale_royalty_bps || 1000;
  let burnable = argv.burnable || true;
  let bridgeable = argv.bridgeable || true;

  deployer.deploy(
    FeralfileExhibitionV3,
    exhibition_name,
    exhibition_symbol,
    exhibition_secondary_sale_royalty_bps,
    exhibition_royalty_payout_address,
    'ipfs://QmWC4bECUTBMwNLsEGSFtTBpCDqjNkoAoxcrsJkFHdfCr6',
    '',
    burnable,
    bridgeable
  );
};
