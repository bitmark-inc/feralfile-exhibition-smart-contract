const FeralfileExhibitionV3_2 = artifacts.require('FeralfileExhibitionV3_2');
const MockOperatorFilterRegistry = artifacts.require(
  'MockOperatorFilterRegistry'
);

const IPFS_GATEWAY_PREFIX = 'https://ipfs.bitmark.com/ipfs/';

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

const ALLOWED_EXCHANGE = '0x2de783974f53AC98b16332f943431df0FD66C9E4';

contract('FeralfileExhibitionV3_2', async (accounts) => {
  before(async function () {
    this.operatorFilterRegistry = await MockOperatorFilterRegistry.new(
      ALLOWED_EXCHANGE
    );
    this.testThumbnailCid = 'QmV68mphFwMraCE9J6KpQc89Sz8ppvJx5CP6XFruhGQrX8'; // AKG test IPFS thumbnail CID
    this.testArtworkCid = 'QmTn3PfHHvoDHKawTPXutqxAk2k8ynFK9cZfsSwggryjkX'; // AKG test IPFS artwork CID

    this.exhibition = await FeralfileExhibitionV3_2.new(
      'Feral File V3_2 Test 002',
      'FFV3',
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX,
      true,
      true
    );

    await this.exhibition.updateOperatorFilterRegistry(
      this.operatorFilterRegistry.address
    );
  });

  it('can register artwork', async function () {
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

  it('test batch mint 1 artwork', async function () {
    let editionIndex0 = 0;
    let editionIndex1 = 1;
    let editionIndex2 = 2;
    let editionIndex3 = 3;

    await this.exhibition.batchMint([
      [this.artworkID, editionIndex0, accounts[0], accounts[0], 'test0'],
      [this.artworkID, editionIndex1, accounts[0], accounts[0], 'test1'],
      [this.artworkID, editionIndex2, accounts[0], accounts[0], 'test2'],
      [this.artworkID, editionIndex3, accounts[0], accounts[0], 'test3'],
    ]);

    let account0 = accounts[0].toLowerCase();

    this.editionID0 = (
      BigInt(this.artworkID) + BigInt(editionIndex0)
    ).toString();
    let ownerOfEdition0 = await this.exhibition.ownerOf(this.editionID0);
    assert.equal(ownerOfEdition0.toLowerCase(), account0);

    this.editionID1 = (
      BigInt(this.artworkID) + BigInt(editionIndex1)
    ).toString();
    let ownerOfEdition1 = await this.exhibition.ownerOf(this.editionID1);
    assert.equal(ownerOfEdition1.toLowerCase(), account0);

    this.editionID2 = (
      BigInt(this.artworkID) + BigInt(editionIndex2)
    ).toString();
    let ownerOfEdition2 = await this.exhibition.ownerOf(this.editionID2);
    assert.equal(ownerOfEdition2.toLowerCase(), account0);

    this.editionID3 = (
      BigInt(this.artworkID) + BigInt(editionIndex3)
    ).toString();
    let ownerOfEdition3 = await this.exhibition.ownerOf(this.editionID3);
    assert.equal(ownerOfEdition3.toLowerCase(), account0);
  });

  it('can not approve an filtered exchange', async function () {
    let onError = false;
    try {
      await this.exhibition.approve(
        this.operatorFilterRegistry.address,
        this.editionID0
      );
    } catch (error) {
      onError = true;
      assert.equal(error.reason, 'operator is not allowed');
    }
    assert.equal(onError, true);
    let addr = await this.exhibition.getApproved(this.editionID0);
    assert.equal(addr, ZERO_ADDRESS);
  });

  it('can approve the allowed exchange', async function () {
    await this.exhibition.approve(ALLOWED_EXCHANGE, this.editionID0);
    let addr = await this.exhibition.getApproved(this.editionID0);
    assert.equal(addr, ALLOWED_EXCHANGE);
  });

  it('can not transfer before approve', async function () {
    let onError = false;
    try {
      await this.exhibition.safeTransferFrom(
        accounts[0],
        accounts[1],
        this.editionID1,
        {
          from: accounts[1],
        }
      );
    } catch (error) {
      onError = true;
      assert.equal(
        error.reason,
        'ERC721: caller is not token owner nor approved'
      );
    }
    assert.equal(onError, true);
  });

  it('can approve the regular address and transfer', async function () {
    let account1 = accounts[1];
    await this.exhibition.approve(account1, this.editionID1);
    let addr = await this.exhibition.getApproved(this.editionID1);
    assert.equal(addr, account1);
    await this.exhibition.safeTransferFrom(
      accounts[0],
      account1,
      this.editionID1,
      {
        from: account1,
      }
    );

    let ownerOfEdition1 = await this.exhibition.ownerOf(this.editionID1);
    assert.equal(ownerOfEdition1, account1);
  });

  it('can approve when unset the registry', async function () {
    await this.exhibition.updateOperatorFilterRegistry(ZERO_ADDRESS);
    let addressBefore = await this.exhibition.getApproved(this.editionID2);
    assert.equal(addressBefore, ZERO_ADDRESS);
    await this.exhibition.approve(
      this.operatorFilterRegistry.address,
      this.editionID2
    );
    let addressAfter = await this.exhibition.getApproved(this.editionID2);
    assert.equal(addressAfter, this.operatorFilterRegistry.address);
  });
});
