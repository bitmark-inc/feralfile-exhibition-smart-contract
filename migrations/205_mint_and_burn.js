var MoMABurn = artifacts.require('MoMABurn');

const argv = require('minimist')(process.argv.slice(2), {
  string: ['exhibition_curator'],
});

module.exports = function (deployer) {
  let burn_token =
    argv.burn_token || '0x1D6D02bfCb4978fCf8300361F4AEfbbd01552907';
  let mint_token =
    argv.mint_token || '0x4cD36964D4fA1426DECF8032a6b99d9858a41Bb7';
  let tier1_artwork_id =
    argv.tier1_artwork_id ||
    '30904499295864258639419635169992685342210506105330261237398947502082530823740';
  let tier2_artwork_id =
    argv.tier2_artwork_id ||
    '33885430236022805519928116200885204946158099275331952512558554217060697131713';
  let tier3_artwork_id =
    argv.tier3_artwork_id ||
    '6513207159759019468158447468875791685910005088450594565135634337538606641299';
  let tier1_ipfs_cid =
    argv.tier1_ipfs_cid || 'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS';
  let tier2_ipfs_cid =
    argv.tier2_ipfs_cid || 'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7j';
  let tier3_ipfs_cid =
    argv.tier3_ipfs_cid || 'QmRRm2iErRsq7hZ12VDoSA1aPHrHudSTcE8vu9kWkmAYtR';

  deployer.deploy(
    MoMABurn,
    burn_token,
    mint_token,
    tier1_artwork_id,
    tier2_artwork_id,
    tier3_artwork_id,
    tier1_ipfs_cid,
    tier2_ipfs_cid,
    tier3_ipfs_cid
  );
};
