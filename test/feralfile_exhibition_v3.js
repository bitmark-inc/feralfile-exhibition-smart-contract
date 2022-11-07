const FeralfileExhibitionV3 = artifacts.require('FeralfileExhibitionV3');

const axios = require('axios');

const IPFS_GATEWAY_PREFIX = 'https://cloudflare-ipfs.com/ipfs/';

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

const originArtworkCID = 'QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc';

contract('FeralfileExhibitionV3', async (accounts) => {
  before(async function () {
    this.exhibition = await FeralfileExhibitionV3.new(
      'Feral File V3 Test 001',
      'FFV3',
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX,
      true,
      true
    );
  });

  it('check contract is burnable', async function () {
    let isBurnable = await this.exhibition.isBurnable();
    assert.equal(isBurnable, true);
  });

  it('check contract is bridgeable', async function () {
    let isBridgeable = await this.exhibition.isBridgeable();
    assert.equal(isBridgeable, true);
  });

  it('create an artwork correctly with no editions on it', async function () {
    let r = await this.exhibition.createArtworks([
      ['TestArtwork', 'TestUser', 'TestArtwork', 10, 2, 1],
    ]);
    this.artworkID = r.logs[0].args.artworkID;

    let artwork = await this.exhibition.artworks(this.artworkID);
    assert.equal(artwork.title, 'TestArtwork');

    let artworkEditions = await this.exhibition.totalEditionOfArtwork(
      this.artworkID
    );
    assert.equal(artworkEditions, 0);
  });

  it('fail to create an artwork with duplicated fingerprint', async function () {
    try {
      let r = await this.exhibition.createArtworks([
        ['TestArtwork', 'TestUser', 'TestArtwork', 5, 2, 1],
      ]);
    } catch (error) {
      console.log(error.message);

      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert an artwork with the same fingerprint has already registered'
      );
    }
  });

  it('fail to create an artwork with empty fingerprint', async function () {
    try {
      let r = await this.exhibition.createArtworks([
        ['TestArtwork', 'TestUser', '', 10, 2, 1],
      ]);
    } catch (error) {
      console.log(error.message);

      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert fingerprint can not be empty'
      );
    }
  });

  it('fail to create an artwork with empty title', async function () {
    try {
      let r = await this.exhibition.createArtworks([
        ['', 'TestUser', 'TestArtwork', 5, 2, 1],
      ]);
    } catch (error) {
      console.log(error.message);

      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert title can not be empty'
      );
    }
  });

  it('fail to create an artwork with empty artist name', async function () {
    try {
      let r = await this.exhibition.createArtworks([
        ['TestArtwork', '', 'TestArtwork', 5, 2, 1],
      ]);
    } catch (error) {
      console.log(error.message);

      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert artist can not be empty'
      );
    }
  });

  it('fail to create an artwork with empty edition size', async function () {
    try {
      let r = await this.exhibition.createArtworks([
        ['TestArtwork', 'TestArtwork', 'TestArtwork', 0, 2, 1],
      ]);
    } catch (error) {
      console.log(error.message);

      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert edition size needs to be at least 1'
      );
    }
  });

  it('can update the IPFS to a new one', async function () {
    let newCID = 'QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXcNew';

    let testEditionNumber = 1;
    let editionID = BigInt(this.artworkID) + BigInt(testEditionNumber);
    await this.exhibition.batchMint([
      [
        this.artworkID,
        testEditionNumber,
        '0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2',
        '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
        'test',
      ],
    ]);
    await this.exhibition.updateArtworkEditionIPFSCid(editionID, newCID);

    let artworkEdition = await this.exhibition.artworkEditions(editionID);

    assert.equal(artworkEdition.ipfsCID, newCID);
  });

  it('cannot update the IPFS cid using an existent IPFS cid', async function () {
    let testEditionNumber = 1;
    let editionID = BigInt(this.artworkID) + BigInt(testEditionNumber);
    try {
      await this.exhibition.updateArtworkEditionIPFSCid(
        editionID,
        originArtworkCID
      );
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert artwork edition is not found'
      );
    }
  });

  it('cannot update the IPFS cid for a non-existent token', async function () {
    let newCID = 'QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXcNew';

    let editionID = 1;
    try {
      await this.exhibition.updateArtworkEditionIPFSCid(editionID, newCID);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert artwork edition is not found'
      );
    }
  });

  it('can not set royalty payout address to zero address', async function () {
    try {
      await this.exhibition.setRoyaltyPayoutAddress(ZERO_ADDRESS);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert invalid royalty payout address'
      );
    }
  });

  it('can not set royalty payout address to zero address', async function () {
    try {
      await this.exhibition.setRoyaltyPayoutAddress(ZERO_ADDRESS);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert invalid royalty payout address'
      );
    }
  });

  it('test batch mint successfully in 1 transaction', async function () {
    let editionNumber1 = 3;
    let editionNumber2 = 4;

    // Mint to 0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2 and transfer to 0xD588e5EC7900C881Cec1843f2EbC73601D75e584
    await this.exhibition.batchMint([
      [this.artworkID, editionNumber1, accounts[0], accounts[0], 'test1'],
      [
        this.artworkID,
        editionNumber2,
        '0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2',
        '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
        'test2',
      ],
    ]);

    let editionID1 = BigInt(this.artworkID) + BigInt(editionNumber1);
    let ownerOfEdition1 = await this.exhibition.ownerOf(editionID1.toString());
    assert.equal(ownerOfEdition1.toLowerCase(), accounts[0].toLowerCase());

    let editionID2 = BigInt(this.artworkID) + BigInt(editionNumber2);
    let ownerOfEdition2 = await this.exhibition.ownerOf(editionID2.toString());
    assert.equal(
      ownerOfEdition2.toLowerCase(),
      '0xD588e5EC7900C881Cec1843f2EbC73601D75e584'.toLowerCase()
    );
  });

  it('test batch transfer artworks successfully in 1 transaction', async function () {
    let editionNumber = 5;

    // Mint to 0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2 and transfer to 0xD588e5EC7900C881Cec1843f2EbC73601D75e584
    await this.exhibition.batchMint([
      [
        this.artworkID,
        editionNumber,
        '0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2',
        '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
        'test',
      ],
    ]);

    let editionID = BigInt(this.artworkID) + BigInt(editionNumber);

    let timestamp = (new Date().getTime() / 1000).toFixed(0);

    // TODO:
    let signature = await axios.post(
      'http://localhost:8081/users/e1h4Q15yN29aVgyowe4oD8nvG2wbvf1r6JFNnLhso17aMBjDi4/sign_abi_parameter',
      {
        types: ['address', 'address', 'uint256', 'uint256'],
        values: [
          '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
          '0x487ba00d91015dcc905bb93b528c12a05fbc7a4f',
          editionID.toString(),
          timestamp,
        ],
      },
      {
        headers: {
          'Content-Type': 'application/json',
          'X-API-KEY': 'w729FVB1S4tCkLiNcs3ljxvd',
        },
      }
    );

    // Transfer item to 0x487ba00d91015dcc905bb93b528c12a05fbc7a4f
    let transferredResult = await this.exhibition.authorizedTransfer([
      [
        '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
        '0x487ba00d91015dcc905bb93b528c12a05fbc7a4f',
        editionID.toString(),
        timestamp,
        signature.data.r,
        signature.data.s,
        signature.data.v,
      ],
    ]);

    let ownerOfEdition = await this.exhibition.ownerOf(editionID.toString());
    assert.equal(
      ownerOfEdition.toLowerCase(),
      '0x487ba00d91015dcc905bb93b528c12a05fbc7a4f'.toLowerCase()
    );
  });

  it('fail to batch transfer out of recv window', async function () {
    let editionNumber = 6;

    // Mint to 0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2 and transfer to 0xD588e5EC7900C881Cec1843f2EbC73601D75e584
    await this.exhibition.batchMint([
      [
        this.artworkID,
        editionNumber,
        '0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2',
        '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
        'test6',
      ],
    ]);

    let editionID = BigInt(this.artworkID) + BigInt(editionNumber);

    let timestamp = (new Date().getTime() / 1000 - 400).toFixed(0);

    try {
      // Transfer item to 0x487ba00d91015dcc905bb93b528c12a05fbc7a4f
      let transferredResult = await this.exhibition.authorizedTransfer([
        [
          '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
          '0x487ba00d91015dcc905bb93b528c12a05fbc7a4f',
          editionID.toString(),
          timestamp,
          '0x4b1be9d4a91bf5f0462e96be0a67e738ac3ee2b191896c51b6b071f35c590563',
          '0x5a3c40da5f67db32aa3e8026f55d3305bc40e025f9ccac5067ef47256bb49483',
          '0x1b',
        ],
      ]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert FeralfileExhibitionV3: timestamp is over recv window'
      );
    }
  });

  it('fail to batch transfer invalid signature', async function () {
    let editionNumber = 7;

    // Mint to 0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2 and transfer to 0xD588e5EC7900C881Cec1843f2EbC73601D75e584
    await this.exhibition.batchMint([
      [
        this.artworkID,
        editionNumber,
        '0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2',
        '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
        'test7',
      ],
    ]);

    let editionID = BigInt(this.artworkID) + BigInt(editionNumber);

    let timestamp = (new Date().getTime() / 1000).toFixed(0);

    try {
      // Transfer item to 0x487ba00d91015dcc905bb93b528c12a05fbc7a4f
      let transferredResult = await this.exhibition.authorizedTransfer([
        [
          '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
          '0x487ba00d91015dcc905bb93b528c12a05fbc7a4f',
          editionID.toString(),
          timestamp,
          '0x4b1be9d4a91bf5f0462e96be0a67e738ac3ee2b191896c51b6b071f35c590563',
          '0x5a3c40da5f67db32aa3e8026f55d3305bc40e025f9ccac5067ef47256bb49483',
          '0x1b',
        ],
      ]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert FeralfileExhibitionV3: the transfer request is not authorized'
      );
    }
  });

  it('test burn edition successfully', async function () {
    let editionNumber = 8;

    await this.exhibition.batchMint([
      [this.artworkID, editionNumber, accounts[0], accounts[0], 'test8'],
    ]);

    let editionID = BigInt(this.artworkID) + BigInt(editionNumber);

    try {
      let burnedResult = await this.exhibition.burnEditions([editionID]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert FeralfileExhibitionV3: the transfer request is not authorized'
      );
    }
  });

  it('test burn edition successfully', async function () {
    let editionNumber = 9;

    await this.exhibition.batchMint([
      [
        this.artworkID,
        editionNumber,
        '0xb824edfc5dced3ac86b6f5816763a35c2ba66fa2',
        '0xD588e5EC7900C881Cec1843f2EbC73601D75e584',
        'test9',
      ],
    ]);

    let editionID = BigInt(this.artworkID) + BigInt(editionNumber);

    try {
      let burnedResult = await this.exhibition.burnEditions([editionID]);
    } catch (error) {
      assert.equal(
        error.message,
        'Returned error: VM Exception while processing transaction: revert ERC721: caller is not token owner nor approved'
      );
    }
  });
});
