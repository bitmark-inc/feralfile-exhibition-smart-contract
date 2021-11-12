const FeralfileExhibitionV2 = artifacts.require("FeralfileExhibitionV2");

const IPFS_GATEWAY_PREFIX = "https://cloudflare-ipfs.com/ipfs/"

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000"

contract("FeralfileExhibitionV2", async accounts => {

  before(async function () {
    this.exhibition = await FeralfileExhibitionV2.new(
      "Feral File V2 Test 001", "FFV2", "Exhibition Title",
      50, 1000, "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
      "https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ",
      IPFS_GATEWAY_PREFIX);
  })

  it("create an artwork correctly with no editions on it", async function () {
    let r = await this.exhibition.createArtwork("TestArtwork", "TestArtwork", "TestUser", 5);
    this.artworkID = r.logs[0].args.artworkID

    let artwork = await this.exhibition.artworks(this.artworkID)
    assert.equal(artwork.title, "TestArtwork");

    let artworkEditions = await this.exhibition.totalEditionOfArtwork(this.artworkID)
    assert.equal(artworkEditions, 0);
  })

  it("fail to create an artwork if the edition size is greater than exhibition limits", async function () {
    try {
      let r = await this.exhibition.createArtwork("TestArtworkFail", "TestArtworkFail", "TestUser", 100);
    } catch (error) {
      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert artwork edition size exceeds the maximum edition size of the exhibition -- Reason given: artwork edition size exceeds the maximum edition size of the exhibition.")
    }
  })

  it("fail to create an artwork with duplicated fingerprint", async function () {
    try {
      let r = await this.exhibition.createArtwork("TestArtwork", "TestArtwork", "TestUser", 5);
    } catch (error) {

      console.log(error.message)

      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert an artwork with the same fingerprint has already registered -- Reason given: an artwork with the same fingerprint has already registered.")
    }
  })

  it("fail to create an artwork with empty fingerprint", async function () {
    try {
      let r = await this.exhibition.createArtwork("", "TestArtwork", "TestUser", 5);
    } catch (error) {

      console.log(error.message)

      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert fingerprint can not be empty -- Reason given: fingerprint can not be empty.")
    }
  })


  it("fail to create an artwork with empty title", async function () {
    try {
      let r = await this.exhibition.createArtwork("TestArtwork", "", "TestUser", 5);
    } catch (error) {

      console.log(error.message)

      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert title can not be empty -- Reason given: title can not be empty.")
    }
  })

  it("fail to create an artwork with empty artist name", async function () {
    try {
      let r = await this.exhibition.createArtwork("TestArtwork", "TestArtwork", "", 5);
    } catch (error) {

      console.log(error.message)

      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert artist can not be empty -- Reason given: artist can not be empty.")
    }
  })



  it("can swap tokens from bitmark blockchain correctly", async function () {
    let tokenOwner = "0xbbed1c61b6e68c397b021c1274080a2005042c08"

    let bitmarkID = Buffer.from("793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e382", "hex")

    let artworkCID = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc"

    let testEditionNumber = 3

    await this.exhibition.swapArtworkFromBitmark(this.artworkID, bitmarkID, testEditionNumber, tokenOwner, artworkCID)

    let editionID = BigInt(this.artworkID) + BigInt(testEditionNumber)
    let artworkEdition = await this.exhibition.artworkEditions(editionID)

    assert.equal(artworkEdition.editionID, editionID)
    assert.equal(artworkEdition.ipfsCID, artworkCID)
  })


  it("can not set royalty payout address to zero address", async function () {
    try {
      await this.exhibition.setRoyaltyPayoutAddress(ZERO_ADDRESS)
    } catch (error) {
      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert invalid royalty payout address -- Reason given: invalid royalty payout address.")
    }
  })
})
