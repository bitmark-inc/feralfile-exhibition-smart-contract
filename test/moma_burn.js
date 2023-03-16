const FeralfileExhibitionV2 = artifacts.require('FeralfileExhibitionV2');
const FeralfileExhibitionV3_3 = artifacts.require('FeralfileExhibitionV3_3');
const MoMABurn = artifacts.require('MoMABurn');

const IPFS_GATEWAY_PREFIX = 'https://cloudflare-ipfs.com/ipfs/';

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

const originArtworkCID = 'QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc';

contract('MoMABurn', async (accounts) => {
  before(async function () {
    this.mmBurn = await MoMABurn.new();

    this.exhibitionV2 = await FeralfileExhibitionV2.new(
      'FeralFile V2 Test',
      'FFV2',
      5000,
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX
    );
    this.exhibitionTier1 = await FeralfileExhibitionV3_3.new(
      'FeralFile V3 Tier 1',
      'FFV3',
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX,
      true,
      true,
      this.mmBurn.address
    );
    this.exhibitionTier2 = await FeralfileExhibitionV3_3.new(
      'FeralFile V3 Tier 2',
      'FFV3',
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX,
      true,
      true,
      this.mmBurn.address
    );
    this.exhibitionTier3 = await FeralfileExhibitionV3_3.new(
      'FeralFile V3 Tier 3',
      'FFV3',
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX,
      true,
      true,
      this.mmBurn.address
    );

    // Init burn artwork
    let burnArtwork = await this.exhibitionV2.createArtwork(
      'Tier1',
      'Tier1',
      'TestUser',
      5000
    );
    this.burnArtworkID = burnArtwork.logs[0].args.artworkID;

    // Init mint artwork
    let newArtworks1 = await this.exhibitionTier1.createArtworks([
      ['Tier1', 'Tier1', 'TestUser1', 800, 0, 0],
    ]);
    this.artworkID1 = newArtworks1.logs[0].args.artworkID;
    await this.exhibitionTier1.updateArtworkCIDs(
      [this.artworkID1],
      ['TestCID1']
    );

    let newArtworks2 = await this.exhibitionTier2.createArtworks([
      ['Tier2', 'Tier2', 'TestUser2', 600, 0, 0],
    ]);
    this.artworkID2 = newArtworks2.logs[0].args.artworkID;
    await this.exhibitionTier2.updateArtworkCIDs(
      [this.artworkID2],
      ['TestCID2']
    );

    let newArtworks3 = await this.exhibitionTier3.createArtworks([
      ['Tier3', 'Tier3', 'TestUser3', 100, 0, 0],
    ]);
    this.artworkID3 = newArtworks3.logs[0].args.artworkID;
    await this.exhibitionTier3.updateArtworkCIDs(
      [this.artworkID3],
      ['TestCID3']
    );
    await this.mmBurn.setBurnAndMintParams(
      this.exhibitionV2.address,
      this.exhibitionTier1.address,
      this.exhibitionTier2.address,
      this.exhibitionTier3.address,
      this.burnArtworkID,
      this.artworkID1,
      this.artworkID2,
      this.artworkID3
    );
  });

  it('test mint and burn not enabled', async function () {
    try {
      await this.mmBurn.burnAndMint([BigInt(0), BigInt(1), BigInt(2)]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert Burn and mint status is not enabled'
      );
    }
  });

  it('test mint and burn successfully for tier 1', async function () {
    let bitmarkID1 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e382',
      'hex'
    );
    let bitmarkID2 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e383',
      'hex'
    );
    let bitmarkID3 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e384',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID1,
      1,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID2,
      2,
      accounts[0],
      'QmeH4njLdkVyUB2QMVdiCMkwzBDxaJJivxpabGGPFcBbAo'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID3,
      3,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7j'
    );

    const editionID1 = BigInt(this.burnArtworkID) + BigInt(1);
    const editionID2 = BigInt(this.burnArtworkID) + BigInt(2);
    const editionID3 = BigInt(this.burnArtworkID) + BigInt(3);

    const owner1 = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2 = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3 = await this.exhibitionV2.ownerOf(editionID3.toString());
    assert.equal(owner1.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3.toLowerCase(), accounts[0].toLowerCase());

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, true);

    await this.mmBurn.setBurnEnabled(true);

    let result = await this.mmBurn.burnAndMint([
      editionID1,
      editionID2,
      editionID3,
    ]);

    const logLength = result.receipt.rawLogs.length;

    const newEditionID = result.receipt.rawLogs[logLength - 1].topics[3];
    const ownerNewEdition = await this.exhibitionTier1.ownerOf(newEditionID);

    const owner1New = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2New = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3New = await this.exhibitionV2.ownerOf(editionID3.toString());

    assert.equal(
      owner1New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner2New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner3New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );

    assert.equal(ownerNewEdition.toLowerCase(), accounts[0].toLowerCase());
  });

  it('test mint and burn successfully for tier 3', async function () {
    let bitmarkID1 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e362',
      'hex'
    );
    let bitmarkID2 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e363',
      'hex'
    );
    let bitmarkID3 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e364',
      'hex'
    );
    let bitmarkID4 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e365',
      'hex'
    );
    let bitmarkID5 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e366',
      'hex'
    );
    let bitmarkID6 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e367',
      'hex'
    );
    let bitmarkID7 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e368',
      'hex'
    );
    let bitmarkID8 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e369',
      'hex'
    );
    let bitmarkID9 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e370',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID1,
      10,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUK1'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID2,
      11,
      accounts[0],
      'QmeH4njLdkVyUB2QMVdiCMkwzBDxaJJivxpabGGPFcBbAK2'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID3,
      12,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7K3'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID4,
      13,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7K4'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID5,
      14,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7K5'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID6,
      15,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7K6'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID7,
      16,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7K7'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID8,
      17,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7K8'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID9,
      18,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7K9'
    );

    let editionID1 = BigInt(this.burnArtworkID) + BigInt(10);
    let editionID2 = BigInt(this.burnArtworkID) + BigInt(11);
    let editionID3 = BigInt(this.burnArtworkID) + BigInt(12);
    let editionID4 = BigInt(this.burnArtworkID) + BigInt(13);
    let editionID5 = BigInt(this.burnArtworkID) + BigInt(14);
    let editionID6 = BigInt(this.burnArtworkID) + BigInt(15);
    let editionID7 = BigInt(this.burnArtworkID) + BigInt(16);
    let editionID8 = BigInt(this.burnArtworkID) + BigInt(17);
    let editionID9 = BigInt(this.burnArtworkID) + BigInt(18);

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, true);

    await this.mmBurn.setBurnEnabled(true);

    let result = await this.mmBurn.burnAndMint([
      editionID1,
      editionID2,
      editionID3,
      editionID4,
      editionID5,
      editionID6,
      editionID7,
      editionID8,
      editionID9,
    ]);

    const logLength = result.receipt.rawLogs.length;

    const newEditionID = result.receipt.rawLogs[logLength - 1].topics[3];
    const ownerNewEdition = await this.exhibitionTier3.ownerOf(newEditionID);

    const owner1New = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2New = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3New = await this.exhibitionV2.ownerOf(editionID3.toString());
    const owner4New = await this.exhibitionV2.ownerOf(editionID4.toString());
    const owner5New = await this.exhibitionV2.ownerOf(editionID5.toString());
    const owner6New = await this.exhibitionV2.ownerOf(editionID6.toString());
    const owner7New = await this.exhibitionV2.ownerOf(editionID7.toString());
    const owner8New = await this.exhibitionV2.ownerOf(editionID8.toString());
    const owner9New = await this.exhibitionV2.ownerOf(editionID9.toString());

    assert.equal(
      owner1New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner2New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner3New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner4New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner5New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner6New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner7New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner8New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );
    assert.equal(
      owner9New.toLowerCase(),
      '0x000000000000000000000000000000000000dEaD'.toLowerCase()
    );

    assert.equal(ownerNewEdition.toLowerCase(), accounts[0].toLowerCase());
  });

  it('test mint and burn fail because number of edition is invalid', async function () {
    let bitmarkID1 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e385',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID1,
      4,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_2'
    );

    let editionID1 = BigInt(this.burnArtworkID) + BigInt(4);

    const owner1 = await this.exhibitionV2.ownerOf(editionID1.toString());
    assert.equal(owner1.toLowerCase(), accounts[0].toLowerCase());

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, true);

    try {
      await this.mmBurn.burnAndMint([editionID1]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert Invalid number of burning editions'
      );
    }

    const owner1New = await this.exhibitionV2.ownerOf(editionID1.toString());

    assert.equal(owner1New.toLowerCase(), accounts[0].toLowerCase());
  });

  it('test mint and burn fail because wrong owner', async function () {
    let bitmarkID1 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e386',
      'hex'
    );
    let bitmarkID2 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e387',
      'hex'
    );
    let bitmarkID3 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e388',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID1,
      5,
      accounts[1], // Mint to account 1
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_3'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID2,
      6,
      accounts[0],
      'QmeH4njLdkVyUB2QMVdiCMkwzBDxaJJivxpabGGPFcBbAo_4'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID3,
      7,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7j_5'
    );

    let editionID1 = BigInt(this.burnArtworkID) + BigInt(5);
    let editionID2 = BigInt(this.burnArtworkID) + BigInt(6);
    let editionID3 = BigInt(this.burnArtworkID) + BigInt(7);

    const owner1 = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2 = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3 = await this.exhibitionV2.ownerOf(editionID3.toString());
    assert.equal(owner1.toLowerCase(), accounts[1].toLowerCase());
    assert.equal(owner2.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3.toLowerCase(), accounts[0].toLowerCase());

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, true);

    try {
      await this.mmBurn.burnAndMint([editionID1, editionID2, editionID3]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert ERC721: caller is not token owner nor approved'
      );
    }

    const owner1New = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2New = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3New = await this.exhibitionV2.ownerOf(editionID3.toString());

    assert.equal(owner1New.toLowerCase(), accounts[1].toLowerCase());
    assert.equal(owner2New.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3New.toLowerCase(), accounts[0].toLowerCase());
  });

  it('test mint and burn fail because not approved', async function () {
    let bitmarkID1 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e489',
      'hex'
    );
    let bitmarkID2 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e490',
      'hex'
    );
    let bitmarkID3 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e491',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID1,
      30,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_6'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID2,
      31,
      accounts[0],
      'QmeH4njLdkVyUB2QMVdiCMkwzBDxaJJivxpabGGPFcBbAo_7'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID3,
      32,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7j_8'
    );

    let editionID1 = BigInt(this.burnArtworkID) + BigInt(30);
    let editionID2 = BigInt(this.burnArtworkID) + BigInt(31);
    let editionID3 = BigInt(this.burnArtworkID) + BigInt(32);

    const owner1 = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2 = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3 = await this.exhibitionV2.ownerOf(editionID3.toString());
    assert.equal(owner1.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3.toLowerCase(), accounts[0].toLowerCase());

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, false);
    try {
      await this.mmBurn.burnAndMint([editionID1, editionID2, editionID3]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert ERC721: caller is not token owner nor approved'
      );
    }

    const owner1New = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2New = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3New = await this.exhibitionV2.ownerOf(editionID3.toString());

    assert.equal(owner1New.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2New.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3New.toLowerCase(), accounts[0].toLowerCase());
  });

  it('test mint and burn fail because using same token input', async function () {
    let bitmarkID1 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e589',
      'hex'
    );
    let bitmarkID2 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e590',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID1,
      33,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_8'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID2,
      34,
      accounts[0],
      'QmeH4njLdkVyUB2QMVdiCMkwzBDxaJJivxpabGGPFcBbAo_9'
    );

    let editionID1 = BigInt(this.burnArtworkID) + BigInt(33);
    let editionID2 = BigInt(this.burnArtworkID) + BigInt(34);

    const owner1 = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2 = await this.exhibitionV2.ownerOf(editionID2.toString());
    assert.equal(owner1.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2.toLowerCase(), accounts[0].toLowerCase());

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, true);
    try {
      await this.mmBurn.burnAndMint([editionID1, editionID2, editionID2]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert Invalid burning editions'
      );
    }

    const owner1New = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2New = await this.exhibitionV2.ownerOf(editionID2.toString());

    assert.equal(owner1New.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2New.toLowerCase(), accounts[0].toLowerCase());
  });

  it('test mint and burn fail because wrong artwork ID', async function () {
    let wrongBurnArtwork = await this.exhibitionV2.createArtwork(
      'Wrong MM',
      'Wrong MM',
      'TestUser',
      5000
    );
    const wrongBurnArtworkID = wrongBurnArtwork.logs[0].args.artworkID;

    let bitmarkID1 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e591',
      'hex'
    );
    let bitmarkID2 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e592',
      'hex'
    );
    let bitmarkID3 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e593',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      wrongBurnArtworkID,
      bitmarkID1,
      35,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_11'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      wrongBurnArtworkID,
      bitmarkID2,
      36,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_12'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      wrongBurnArtworkID,
      bitmarkID3,
      37,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_13'
    );

    const editionID1 = BigInt(wrongBurnArtworkID) + BigInt(35);
    const editionID2 = BigInt(wrongBurnArtworkID) + BigInt(36);
    const editionID3 = BigInt(wrongBurnArtworkID) + BigInt(37);

    const owner1 = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2 = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3 = await this.exhibitionV2.ownerOf(editionID3.toString());
    assert.equal(owner1.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3.toLowerCase(), accounts[0].toLowerCase());

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, true);

    await this.mmBurn.setBurnEnabled(true);

    try {
      await this.mmBurn.burnAndMint([editionID1, editionID2, editionID3]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert Edition ID is not support for burning'
      );
    }

    const owner1New = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2New = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3New = await this.exhibitionV2.ownerOf(editionID3.toString());

    assert.equal(owner1New.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2New.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3New.toLowerCase(), accounts[0].toLowerCase());
  });
});
