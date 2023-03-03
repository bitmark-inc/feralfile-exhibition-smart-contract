const FeralfileExhibitionV2 = artifacts.require('FeralfileExhibitionV2');
const FeralfileExhibitionV3_3 = artifacts.require('FeralfileExhibitionV3_3');
const MoMABurn = artifacts.require('MoMABurn');

const IPFS_GATEWAY_PREFIX = 'https://cloudflare-ipfs.com/ipfs/';

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

const originArtworkCID = 'QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc';

contract('MoMABurn', async (accounts) => {
  before(async function () {
    this.exhibitionV2 = await FeralfileExhibitionV2.new(
      'FeralFile V2 Test',
      'FFV2',
      5000,
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX
    );
    this.exhibitionV3 = await FeralfileExhibitionV3_3.new(
      'FeralFile V3 Test',
      'FFV3',
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX,
      true,
      true
    );

    // Init burn contract
    let burnArtwork = await this.exhibitionV2.createArtwork(
      'Tier1',
      'Tier1',
      'TestUser',
      5000
    );

    this.burnArtworkID = burnArtwork.logs[0].args.artworkID;

    let newArtworks = await this.exhibitionV3.createArtworks([
      ['Tier1', 'Tier1', 'TestUser1', 800, 0, 0],
      ['Tier2', 'Tier2', 'TestUser2', 600, 0, 0],
      ['Tier3', 'Tier3', 'TestUser3', 100, 0, 0],
    ]);

    this.artworkID1 = newArtworks.logs[0].args.artworkID;
    this.artworkID2 = newArtworks.logs[1].args.artworkID;
    this.artworkID3 = newArtworks.logs[2].args.artworkID;

    this.mmBurn = await MoMABurn.new(
      this.exhibitionV2.address,
      this.exhibitionV3.address,
      this.artworkID1,
      this.artworkID2,
      this.artworkID3,
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS',
      'QmeH4njLdkVyUB2QMVdiCMkwzBDxaJJivxpabGGPFcBbAo',
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7j'
    );

    await this.exhibitionV3.addTrustee(this.mmBurn.address);
  });

  it('test mint and burn successfully', async function () {
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

    let editionID1 = BigInt(this.burnArtworkID) + BigInt(1);
    let editionID2 = BigInt(this.burnArtworkID) + BigInt(2);
    let editionID3 = BigInt(this.burnArtworkID) + BigInt(3);

    const owner1 = await this.exhibitionV2.ownerOf(editionID1.toString());
    const owner2 = await this.exhibitionV2.ownerOf(editionID2.toString());
    const owner3 = await this.exhibitionV2.ownerOf(editionID3.toString());
    assert.equal(owner1.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner2.toLowerCase(), accounts[0].toLowerCase());
    assert.equal(owner3.toLowerCase(), accounts[0].toLowerCase());

    await this.exhibitionV2.setApprovalForAll(this.mmBurn.address, true);

    let result = await this.mmBurn.burnAndMint([
      editionID1,
      editionID2,
      editionID3,
    ]);

    const newEditionID = result.receipt.rawLogs[7].topics[3];
    const ownerNewEdition = await this.exhibitionV3.ownerOf(newEditionID);

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
      console.log(error.message);

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
      console.log(error.message);

      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert Invalid owner of burned token'
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
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e389',
      'hex'
    );
    let bitmarkID2 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e390',
      'hex'
    );
    let bitmarkID3 = Buffer.from(
      '793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e391',
      'hex'
    );

    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID1,
      8,
      accounts[0],
      'QmSd1xQLhKTRsYzUYLmBmYeNfA9ufj2MHN2G87UjpDrMUS_6'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID2,
      9,
      accounts[0],
      'QmeH4njLdkVyUB2QMVdiCMkwzBDxaJJivxpabGGPFcBbAo_7'
    );
    await this.exhibitionV2.swapArtworkFromBitmark(
      this.burnArtworkID,
      bitmarkID3,
      10,
      accounts[0],
      'QmRMg9m6M9CaBotYaKzGr47z6hjRQM2xuyRyBFWfjwGz7j_8'
    );

    let editionID1 = BigInt(this.burnArtworkID) + BigInt(8);
    let editionID2 = BigInt(this.burnArtworkID) + BigInt(9);
    let editionID3 = BigInt(this.burnArtworkID) + BigInt(10);

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
      console.log(error.message);

      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert Token need to approve for this contract'
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
