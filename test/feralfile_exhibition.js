const FeralfileExhibition = artifacts.require("FeralfileExhibition");

const IPFS_GATEWAY_PREFIX = "https://cloudflare-ipfs.com/ipfs/"

contract("FeralfileExhibition", async accounts => {

  before(async function () {
    this.exhibition = await FeralfileExhibition.new(
      "Social Codes", "SSC", "0x8fd310de32848798eB64Bd88f9C5656Eea32415e",
      50, 10, IPFS_GATEWAY_PREFIX);
  })

  it("create an artwork correctly with no editions on it", async function () {
    let r = await this.exhibition.createArtwork(Buffer("TestArtwork"), "TestArtwork", accounts[0], "html", "", 5, 5);
    this.artworkID = r.logs[0].args._artworkID

    let artwork = await this.exhibition.artworks(this.artworkID)
    assert.equal(artwork.title, "TestArtwork");

    let artworkEditions = await this.exhibition.totalEditionOfArtwork(this.artworkID)
    assert.equal(artworkEditions, 0);
  })

  it("fail to create an artwork if the edition size is greater than exhibition limits", async function () {
    try {
      let r = await this.exhibition.createArtwork(Buffer("TestArtworkFail"), "TestArtworkFail", accounts[0], "html", "", 100, 5);
    } catch (error) {
      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert edition size exceeds the maximum edition size of the exhibition -- Reason given: edition size exceeds the maximum edition size of the exhibition.")
    }
  })

  it("fail to create an artwork with duplicated fingerprint", async function () {
    try {
      let r = await this.exhibition.createArtwork(Buffer("TestArtwork"), "TestArtwork", accounts[0], "html", "", 10, 5);
    } catch (error) {
      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert duplicated fingerprint -- Reason given: duplicated fingerprint.")
    }
  })

  it("mint two editions with different IPFS cid", async function () {
    await this.exhibition.mintArtwork(this.artworkID, ["QmYBfVBh9U5huzY72GucDdwvp35yhrWfgAdiUUfnAZPd7f", "QmZJKajwAeXFwL7zB5mhc7zLhsrw2zwXPcFp42mDexAp1u"])
    let artworkEditions = await this.exhibition.totalEditionOfArtwork(this.artworkID)
    assert.equal(artworkEditions, 2);

    let firstArtworkEditionID = (await this.exhibition.getArtworkEditionByIndex(this.artworkID, 0))
    let firstArtworkEdition = await this.exhibition.artworkEditions(firstArtworkEditionID)
    assert.equal(firstArtworkEdition.ipfsCID, "QmYBfVBh9U5huzY72GucDdwvp35yhrWfgAdiUUfnAZPd7f")
    let firstTokenURI = await this.exhibition.tokenURI(firstArtworkEditionID)
    assert.equal(firstTokenURI, "https://cloudflare-ipfs.com/ipfs/QmYBfVBh9U5huzY72GucDdwvp35yhrWfgAdiUUfnAZPd7f/metadata.json")

    let secondArtworkEditionID = (await this.exhibition.getArtworkEditionByIndex(this.artworkID, 1))
    let secondArtworkEdition = await this.exhibition.artworkEditions(secondArtworkEditionID)
    assert.equal(secondArtworkEdition.ipfsCID, "QmZJKajwAeXFwL7zB5mhc7zLhsrw2zwXPcFp42mDexAp1u")
    let secondTokenURI = await this.exhibition.tokenURI(secondArtworkEditionID)
    assert.equal(secondTokenURI, "https://cloudflare-ipfs.com/ipfs/QmZJKajwAeXFwL7zB5mhc7zLhsrw2zwXPcFp42mDexAp1u/metadata.json")
  })


  it("fails to mint duplicated IPFS cid", async function () {
    try {
      await this.exhibition.mintArtwork(this.artworkID, ["QmYBfVBh9U5huzY72GucDdwvp35yhrWfgAdiUUfnAZPd7f"])
    } catch (error) {
      assert.equal(error.message, "Returned error: VM Exception while processing transaction: revert ipfs id registered -- Reason given: ipfs id registered.")
    }
  })

  it("can swap tokens from bitmark blockchain correctly", async function () {
    let tokenOwner = "0xbbed1c61b6e68c397b021c1274080a2005042c08"

    let bitmarkID = Buffer.from("793eb498de01783e0e350830ea43d24d3c340c7ddee07dbb79c74a840e29e382", "hex")

    let artworkCID = "QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc"

    await this.exhibition.swapArtworkFromBitmarks(this.artworkID, bitmarkID, tokenOwner, artworkCID)

    let editionID = (await this.exhibition.tokenOfOwnerByIndex(tokenOwner, 0))
    let artworkEdition = await this.exhibition.artworkEditions(editionID)
    assert.equal(artworkEdition.ipfsCID, artworkCID)
  })

  it("can preserve provenance for every transfer", async function () {
    let r = await this.exhibition.createArtwork(Buffer("TestArtworkProvenance"), "TestArtworkProvenance", accounts[0], "html", "", 5, 5);
    let artworkID = r.logs[0].args._artworkID

    mintResult = (await this.exhibition.mintArtwork(artworkID, ["QmYBfVBh9U5huzY72GucDdwvp35yhrWfgAdiUUfnAZPd6f"]))

    let artworkEditionID = (await this.exhibition.getArtworkEditionByIndex(artworkID, 0))

    assert.equal(accounts[0], await this.exhibition.ownerOf(artworkEditionID))

    await this.exhibition.transferFrom(accounts[0], accounts[1], artworkEditionID)

    assert.equal(accounts[1], await this.exhibition.ownerOf(artworkEditionID))

    let provenances = (await this.exhibition.editionProvenances(artworkEditionID))

    let addresses = provenances[0]

    assert.equal(accounts[0], addresses[0])
    assert.equal(accounts[1], addresses[1])

  })
})
