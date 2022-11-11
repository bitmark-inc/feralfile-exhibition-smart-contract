const FeralfileExhibitionV3 = artifacts.require('FeralfileExhibitionV31');
const MockDecentraland = artifacts.require('MockDecentraland');

const axios = require('axios');

const IPFS_GATEWAY_PREFIX = 'https://ipfs.bitmark.com/ipfs/';

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

const originArtworkCID = 'QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc';

contract('FeralfileExhibitionV31', async (accounts) => {
  before(async function () {
    this.decentraland = await MockDecentraland.new()

    this.exhibition = await FeralfileExhibitionV3.new(
      'Feral File V3 Test 002',
      'FFV3',
      1000,
      '0x8fd310de32848798eB64Bd88f9C5656Eea32415e',
      'https://ipfs.bitmark.com/ipfs/QmaptARVxNSP36PQai5oiCPqbrATvpydcJ8SPx6T6Yp1CZ',
      IPFS_GATEWAY_PREFIX,
      true,
      true
    );

    this.exhibition.updateDecentralandInfo(this.decentraland.address, 0)
  });

  it('check contract is burnable', async function () {
    let isBurnable = await this.exhibition.isBurnable();
    assert.equal(isBurnable, true);
  });

  it('check contract is bridgeable', async function () {
    let isBridgeable = await this.exhibition.isBridgeable();
    assert.equal(isBridgeable, true);
  });

  it('create an artwork with an external IPFS link', async function () {
    let testIPFSCid = "QmaqM7RypBoHJZyuNwsMNQ7e7vUMyGXuxHrAmrs4khPYyz" // AKG test IPFS artwork CID
    let r = await this.exhibition.createArtworkWithIPFSCID(
      ['TestArtwork', 'TestUser', 'TestArtwork-fingerprint', 10, 2, 1],
      testIPFSCid
    );
    this.artworkID = r.logs[0].args.artworkID;

    let artwork = await this.exhibition.artworks(this.artworkID);
    assert.equal(artwork.title, 'TestArtwork');

    let artworkEditions = await this.exhibition.totalEditionOfArtwork(
      this.artworkID
    );
    assert.equal(artworkEditions, 0);

    let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(this.artworkID)
    assert.equal(testIPFSCid, artworkCIDInContract);
  });

  it('test batch mint 4 artworks for external artwork', async function () {
    let editionNumber0 = 0;
    let editionNumber1 = 1;
    let editionNumber2 = 2;
    let editionNumber3 = 3;


    await this.exhibition.batchMint([
      [this.artworkID, editionNumber0, accounts[0], accounts[0], 'test0'],
      [this.artworkID, editionNumber1, accounts[0], accounts[1], 'test1'],
      [this.artworkID, editionNumber2, accounts[0], accounts[2], 'test2'],
      [this.artworkID, editionNumber3, accounts[0], accounts[3], 'test3'],
    ]);

    let editionID0 = BigInt(this.artworkID) + BigInt(editionNumber0);
    let ownerOfEdition0 = await this.exhibition.ownerOf(editionID0.toString());
    assert.equal(ownerOfEdition0.toLowerCase(), accounts[0].toLowerCase());

    let editionID1 = BigInt(this.artworkID) + BigInt(editionNumber1);
    let ownerOfEdition1 = await this.exhibition.ownerOf(editionID1.toString());
    assert.equal(ownerOfEdition1.toLowerCase(), accounts[1].toLowerCase());

    let editionID2 = BigInt(this.artworkID) + BigInt(editionNumber2);
    let ownerOfEdition2 = await this.exhibition.ownerOf(editionID2.toString());
    assert.equal(ownerOfEdition2.toLowerCase(), accounts[2].toLowerCase());

    let editionID3 = BigInt(this.artworkID) + BigInt(editionNumber3);
    let ownerOfEdition3 = await this.exhibition.ownerOf(editionID3.toString());
    assert.equal(ownerOfEdition3.toLowerCase(), accounts[3].toLowerCase());
  })

  it('test build metadata for artwork with IPFS', async function () {
    let editionNumber1 = 1;
    let editionID1 = BigInt(this.artworkID) + BigInt(editionNumber1);

    let tokenURI = await this.exhibition.tokenURI(editionID1)
    assert.equal(tokenURI, "data:application/json;base64,eyJuYW1lIjoiVGVzdEFydHdvcmsiLCAiYXJ0d29ya19pZCI6IjMyNzkxMjY5MzU4MDE5MTk0ODk2MzMwMTk1MjYzMzIyNzUyNTY2ODM2MTkzMjE4MzY5MzU2NTAxODc5MTA0MjI4MjI4NzcxNzAwNjA5IiwgImltYWdlIjoiIiwgImFuaW1hdGlvbl91cmwiOiJkYXRhOnRleHQvaHRtbDtiYXNlNjQsUENGRVQwTlVXVkJGSUdoMGJXdytQR2gwYld3Z2JHRnVaejBpWlc0aVBqeG9aV0ZrUGp4elkzSnBjSFErSUhaaGNpQmhjblIzYjNKclJHRjBZU0E5SUh0c1lXNWtUM2R1WlhJNklqQjRNREUzTjJKaE1tRmtNRFZqWWpFeVpEWmtaRFl6TnprMk0yVXdOREF3T1RJMk56WmxaVGhrWlNJc0lHOTNibVZ5UVhKeVlYazZXeUl3ZUdZeVltWXlaVGswWTJGak1qTmhNelptTmpNNE5USXlaR0V6TkdaaE1qVXdNREUzT0RNd04yVWlMQ0l3ZURJMk1HVTNOekV3TldVNU1UTmxZbVJtWlRRek5UTXpNVGczTURBek9UbGlNVFJrTXpFNE1HVWlMQ0l3ZUdZek5ESmpOVEV4WTJGak1HSXhPRFZpT1dKaE9UQTBaV05qTTJFd01HUmlZMlpqTTJNMlpXUWlMQ0l3ZURrNE1EQXpZemt5WmpaaE9EYzJaVFEzTm1VeE9XRTJPRGRrTW1NMlpXSTVaR1V5T0dWak9UZ2lMRjE5SUR3dmMyTnlhWEIwUGp4elkzSnBjSFErZDJsdVpHOTNMbUZrWkVWMlpXNTBUR2x6ZEdWdVpYSW9JbTFsYzNOaFoyVWlMR1oxYm1OMGFXOXVLR1VwZTJ4bGRDQjBQV1J2WTNWdFpXNTBMbWRsZEVWc1pXMWxiblJDZVVsa0tDSnRZV2x1Wm5KaGJXVWlLVHRwWmlnaWIySnFaV04wSWowOWRIbHdaVzltSUdVdVpHRjBZU2w3YkdWMElHNDlaUzVrWVhSaExtNWxkMGhsYVdkb2REdHVKaVlvZEM1emRIbHNaUzV0YVc1SVpXbG5hSFE5YmlzaWNIZ2lMSFF1YzNSNWJHVXVhR1ZwWjJoMFBXNHJJbkI0SWlsOWZTazdablZ1WTNScGIyNGdhVzVwZEVSaGRHRW9LWHRzWlhRZ1pUMWtiMk4xYldWdWRDNW5aWFJGYkdWdFpXNTBRbmxKWkNnaWJXRnBibVp5WVcxbElpazdZWEowZDI5eWEwUmhkR0VtSm1VdVkyOXVkR1Z1ZEZkcGJtUnZkeTV3YjNOMFRXVnpjMkZuWlNoaGNuUjNiM0pyUkdGMFlTd2lLaUlwZlR3dmMyTnlhWEIwUGp3dmFHVmhaRDQ4WW05a2VTQnpkSGxzWlQwaWIzWmxjbVpzYjNjdGVEb2dhR2xrWkdWdU95QndZV1JrYVc1bk9pQXdPeUJ0WVhKbmFXNDZJREE3SUhkcFpIUm9PaUF4TURBbE95SWdiMjVzYjJGa1BTSnBibWwwUkdGMFlTZ3BJajQ4YVdaeVlXMWxJR2xrUFNKdFlXbHVabkpoYldVaUlITjBlV3hsUFNKa2FYTndiR0Y1T21Kc2IyTnJPeUJ3WVdSa2FXNW5PaUF3T3lCdFlYSm5hVzQ2SURBN0lHSnZjbVJsY2pwdWIyNWxPeUIzYVdSMGFEb2dNVEF3SlRzaUlITnlZejBpYUhSMGNITTZMeTlwY0daekxtSnBkRzFoY21zdVkyOXRMMmx3Wm5NdlVXMWhjVTAzVW5sd1FtOUlTbHA1ZFU1M2MwMU9VVGRsTjNaVlRYbEhXSFY0U0hKQmJYSnpOR3RvVUZsNWVpSStQQzlwWm5KaGJXVStJRHd2WW05a2VUNDhMMmgwYld3KyJ9");
  })


  it('create a regular artwork', async function () {
    let r = await this.exhibition.createArtworks([
      ['TestArtwork-regular', 'TestUser', 'TestArtwork-regular-fingerprint', 10, 2, 1]
    ]);
    this.artworkRegularID = r.logs[0].args.artworkID;

    let artwork = await this.exhibition.artworks(this.artworkRegularID);
    assert.equal(artwork.title, 'TestArtwork-regular');

    let artworkEditions = await this.exhibition.totalEditionOfArtwork(
      this.artworkRegularID
    );
    assert.equal(artworkEditions, 0);

    let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(this.artworkRegularID)
    assert.equal("", artworkCIDInContract);
  });

  it('test batch mint 2 artworks for regular artwork', async function () {
    let editionNumber0 = 0;
    let editionNumber1 = 1;

    await this.exhibition.batchMint([
      [this.artworkRegularID, editionNumber0, accounts[0], accounts[0], 'test1-0'],
      [this.artworkRegularID, editionNumber1, accounts[0], accounts[1], 'test1-1'],
    ]);

    let editionID0 = BigInt(this.artworkRegularID) + BigInt(editionNumber0);
    let editionID1 = BigInt(this.artworkRegularID) + BigInt(editionNumber1);

    let tokenURI0 = await this.exhibition.tokenURI(editionID0)
    assert.equal(tokenURI0, IPFS_GATEWAY_PREFIX + "test1-0");

    let tokenURI1 = await this.exhibition.tokenURI(editionID1)
    assert.equal(tokenURI1, IPFS_GATEWAY_PREFIX + "test1-1");
  });
});
