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
    assert.equal(tokenURI, "data:application/json;base64,eyJuYW1lIjoiVGVzdEFydHdvcmsiLCAiYXJ0d29ya19pZCI6IjMyNzkxMjY5MzU4MDE5MTk0ODk2MzMwMTk1MjYzMzIyNzUyNTY2ODM2MTkzMjE4MzY5MzU2NTAxODc5MTA0MjI4MjI4NzcxNzAwNjA5IiwgImltYWdlIjoiIiwgImFuaW1hdGlvbl91cmwiOiJkYXRhOnRleHQvaHRtbDtiYXNlNjQsUENGRVQwTlVXVkJGSUdoMGJXdytQR2gwYld3Z2JHRnVaejBpWlc0aVBqeG9aV0ZrUGp4elkzSnBjSFErSUhaaGNpQmtaV1poZFd4MFFYSjBkMjl5YTBSaGRHRTlJSHRzWVc1a1QzZHVaWEk2SWpCNE1ERTNOMkpoTW1Ga01EVmpZakV5WkRaa1pEWXpOemsyTTJVd05EQXdPVEkyTnpabFpUaGtaU0lzSUc5M2JtVnlRWEp5WVhrNld5SXdlREZtWkdGak5UbGtaalUzWVdKbE1qTTVaRGd3TlRrek0yRmhPRGcwWmpZNVl6ZGlPR1ppTlRJaUxDSXdlR0UzTjJNM1pXRm1ZbVJoTmpZM05UUmlPV1EyWmpNeU9HTmlaVGt3WXpsaE1qUTRabVl6WXpjaUxDSXdlREptTkRaa1lUTXhNamhsTkRSaE9XWm1NR05sTTJFNU5tRmxOakV4Wmpoak0yWXhZVFkyWW1RaUxDSXdlR0U0TURFeU5HRTFaV1F3TmpOa1ltWmlaVEUxTlRObVptVTBOak5rTldaa016VmhOR0kwT0RraUxGMTlQQzl6WTNKcGNIUStQSE5qY21sd2RENXNaWFFnWVd4c2IzZFBjbWxuYVc1elBYc2lhSFIwY0hNNkx5OW1aWEpoYkdacGJHVXVZMjl0SWpvaE1IMDdablZ1WTNScGIyNGdjbVZ6YVhwbFNXWnlZVzFsS0hRcGUyeGxkQ0JsUFdSdlkzVnRaVzUwTG1kbGRFVnNaVzFsYm5SQ2VVbGtLQ0p0WVdsdVpuSmhiV1VpS1R0MEppWW9aUzV6ZEhsc1pTNXRhVzVJWldsbmFIUTlkQ3NpY0hnaUxHVXVjM1I1YkdVdWFHVnBaMmgwUFhRckluQjRJaWw5Wm5WdVkzUnBiMjRnYVc1cGRFUmhkR0VvS1h0d2RYTm9RWEowZDI5eWEwUmhkR0ZVYjBsbWNtRnRaU2hrWldaaGRXeDBRWEowZDI5eWEwUmhkR0VwZldaMWJtTjBhVzl1SUhCMWMyaEJjblIzYjNKclJHRjBZVlJ2U1daeVlXMWxLSFFwZTNRbUptUnZZM1Z0Wlc1MExtZGxkRVZzWlcxbGJuUkNlVWxrS0NKdFlXbHVabkpoYldVaUtTNWpiMjUwWlc1MFYybHVaRzkzTG5CdmMzUk5aWE56WVdkbEtIUXNJaW9pS1gxbWRXNWpkR2x2YmlCMWNHUmhkR1ZCY25SdmQzSnJSR0YwWVNoMEtYdGtiMk4xYldWdWRDNW5aWFJGYkdWdFpXNTBRbmxKWkNnaWJXRnBibVp5WVcxbElpa3VZMjl1ZEdWdWRGZHBibVJ2ZHk1d2IzTjBUV1Z6YzJGblpTaDBMQ0lxSWlsOWQybHVaRzkzTG1Ga1pFVjJaVzUwVEdsemRHVnVaWElvSW0xbGMzTmhaMlVpTEdaMWJtTjBhVzl1S0hRcGUyRnNiRzkzVDNKcFoybHVjMXQwTG05eWFXZHBibDAvSW5Wd1pHRjBaUzFoY25SM2IzSnJMV1JoZEdFaVBUMDlkQzVrWVhSaExuUjVjR1VtSm5Wd1pHRjBaVUZ5ZEc5M2NtdEVZWFJoS0hRdVpHRjBZUzVoY25SM2IzSnJSR0YwWVNrNkltOWlhbVZqZENJOVBYUjVjR1Z2WmlCMExtUmhkR0VtSmlKeVpYTnBlbVV0YVdaeVlXMWxJajA5UFhRdVpHRjBZUzUwZVhCbEppWnlaWE5wZW1WSlpuSmhiV1VvZEM1a1lYUmhMbTVsZDBobGFXZG9kQ2w5S1RzOEwzTmpjbWx3ZEQ0OEwyaGxZV1ErUEdKdlpIa2djM1I1YkdVOUltOTJaWEptYkc5M0xYZzZJR2hwWkdSbGJqc2djR0ZrWkdsdVp6b2dNRHNnYldGeVoybHVPaUF3T3lCM2FXUjBhRG9nTVRBd0pUc2lJRzl1Ykc5aFpEMGlhVzVwZEVSaGRHRW9LU0krUEdsbWNtRnRaU0JwWkQwaWJXRnBibVp5WVcxbElpQnpkSGxzWlQwaVpHbHpjR3hoZVRwaWJHOWphenNnY0dGa1pHbHVaem9nTURzZ2JXRnlaMmx1T2lBd095QmliM0prWlhJNmJtOXVaVHNnZDJsa2RHZzZJREV3TUNVN0lpQnpjbU05SW1oMGRIQnpPaTh2YVhCbWN5NWlhWFJ0WVhKckxtTnZiUzlwY0daekwxRnRZWEZOTjFKNWNFSnZTRXBhZVhWT2QzTk5UbEUzWlRkMlZVMTVSMWgxZUVoeVFXMXljelJyYUZCWmVYb2lQand2YVdaeVlXMWxQaUE4TDJKdlpIaytQQzlvZEcxc1BnPT0ifQ==");
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
