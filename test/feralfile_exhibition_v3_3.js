const FeralfileExhibitionV3_3 = artifacts.require('FeralfileExhibitionV3_3');
const MockDecentraland = artifacts.require('MockDecentraland');

const axios = require('axios');

const IPFS_GATEWAY_PREFIX = 'https://ipfs.bitmark.com/ipfs/';

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';

const originArtworkCID = 'QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc';

contract('FeralfileExhibitionV3_3', async (accounts) => {
  before(async function () {
    this.decentraland = await MockDecentraland.new()
    this.testThumbnailCid = "QmV68mphFwMraCE9J6KpQc89Sz8ppvJx5CP6XFruhGQrX8" // AKG test IPFS thumbnail CID
    this.testArtworkCid = "QmXsDVEszo4XzyVK3UF8mKzUNfaUT3fpXwarca6VmGPwvq" // AKG test IPFS artwork CID

    this.exhibition = await FeralfileExhibitionV3_3.new(
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


  it('create an artwork with an external IPFS link', async function () {
    let r = await this.exhibition.createArtworkWithIPFSCID(
      ['TestArtwork', 'TestUser', 'TestArtwork-fingerprint', 10, 2, 1],
      this.testThumbnailCid, this.testArtworkCid
    );
    this.artworkID = r.logs[0].args.artworkID;

    let artwork = await this.exhibition.artworks(this.artworkID);
    assert.equal(artwork.title, 'TestArtwork');

    let artworkEditions = await this.exhibition.totalEditionOfArtwork(
      this.artworkID
    );
    assert.equal(artworkEditions, 0);

    let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(this.artworkID)
    assert.equal(this.testThumbnailCid, artworkCIDInContract.thumbnailCID);
    assert.equal(this.testArtworkCid, artworkCIDInContract.artworkCID);
  });

  it('test batch mint 4 artworks', async function () {
    let editionIndex0 = 0;
    let editionIndex1 = 1;
    let editionIndex2 = 2;
    let editionIndex3 = 3;
    let editionIndex4 = 4;

    await this.exhibition.batchMint([
      [this.artworkID, editionIndex0, accounts[0], accounts[0], 'test0'],
      [this.artworkID, editionIndex1, accounts[0], accounts[1], 'test1'],
      [this.artworkID, editionIndex2, accounts[0], accounts[2], 'test2'],
      [this.artworkID, editionIndex3, accounts[0], accounts[2], 'test3'],
      [this.artworkID, editionIndex4, accounts[0], accounts[3], 'test4'],
    ]);

    let editionID0 = BigInt(this.artworkID) + BigInt(editionIndex0);
    let ownerOfEdition0 = await this.exhibition.ownerOf(editionID0.toString());
    assert.equal(ownerOfEdition0.toLowerCase(), accounts[0].toLowerCase());
    let editionNumber0 = await this.exhibition.tokenEditionNumber(editionID0.toString());
    assert.equal(editionNumber0, "AE");

    let editionID1 = BigInt(this.artworkID) + BigInt(editionIndex1);
    let ownerOfEdition1 = await this.exhibition.ownerOf(editionID1.toString());
    assert.equal(ownerOfEdition1.toLowerCase(), accounts[1].toLowerCase());
    let editionNumber1 = await this.exhibition.tokenEditionNumber(editionID1.toString());
    assert.equal(editionNumber1, "AE");

    let editionID2 = BigInt(this.artworkID) + BigInt(editionIndex2);
    let ownerOfEdition2 = await this.exhibition.ownerOf(editionID2.toString());
    assert.equal(ownerOfEdition2.toLowerCase(), accounts[2].toLowerCase());
    let editionNumber2 = await this.exhibition.tokenEditionNumber(editionID2.toString());
    assert.equal(editionNumber2, "PP");

    let editionID3 = BigInt(this.artworkID) + BigInt(editionIndex3);
    let ownerOfEdition3 = await this.exhibition.ownerOf(editionID3.toString());
    assert.equal(ownerOfEdition3.toLowerCase(), accounts[2].toLowerCase());
    let editionNumber3 = await this.exhibition.tokenEditionNumber(editionID3.toString());
    assert.equal(editionNumber3, "#1");

    let editionID4 = BigInt(this.artworkID) + BigInt(editionIndex4);
    let ownerOfEdition4 = await this.exhibition.ownerOf(editionID4.toString());
    assert.equal(ownerOfEdition4.toLowerCase(), accounts[3].toLowerCase());
    let editionNumber4 = await this.exhibition.tokenEditionNumber(editionID4.toString());
    assert.equal(editionNumber4, "#2");
  })

  it('test build metadata for artwork with IPFS', async function () {
    let editionNumber1 = 1;

    // let expectedMetadata = "data:application/json;base64,eyJpZCI6IjMyNzkxMjY5MzU4MDE5MTk0ODk2MzMwMTk1MjYzMzIyNzUyNTY2ODM2MTkzMjE4MzY5MzU2NTAxODc5MTA0MjI4MjI4NzcxNzAwNjEwIiwgIm5hbWUiOiJUZXN0QXJ0d29yayBBRSIsICJhcnRpc3QiOiJUZXN0VXNlciIsICJhcnR3b3JrX2lkIjoiMHg0ODdmMzM2M2VhODc1ODMyODU3NGE4ODI5NDM0OTNkNGFkOTQ3YWVkYzgwZjc3MTgwY2M3NmZhMGQyZmQ1ZjgxIiwgImVkaXRpb25faW5kZXgiOiIxIiwgImVkaXRpb25fbnVtYmVyIjoiQUUiLCAiZXh0ZXJuYWxfdXJsIjoiaHR0cHM6Ly9mZXJhbGZpbGUuY29tIiwgImltYWdlIjoiaHR0cHM6Ly9pcGZzLmJpdG1hcmsuY29tL2lwZnMvUW1hcU03UnlwQm9ISlp5dU53c01OUTdlN3ZVTXlHWHV4SHJBbXJzNGtoUFl5eiIsICJhbmltYXRpb25fdXJsIjoiZGF0YTp0ZXh0L2h0bWw7YmFzZTY0LFBDRkVUME5VV1ZCRklHaDBiV3crUEdoMGJXd2diR0Z1WnowaVpXNGlQanhvWldGa1BqeHpZM0pwY0hRK0lIWmhjaUJrWldaaGRXeDBRWEowZDI5eWEwUmhkR0U5SUh0c1lXNWtUM2R1WlhJNklqQjRNREUzTjJKaE1tRmtNRFZqWWpFeVpEWmtaRFl6TnprMk0yVXdOREF3T1RJMk56WmxaVGhrWlNJc0lHOTNibVZ5UVhKeVlYazZXeUl3ZUdJMU5UWXhNMk0xWWpBNU9UaGxOREptTWpNNU1UVmtORFJqTWpReE1EYzBOR00xTTJFeFlUUWlMQ0l3ZUdWa05EZGtOREF6WldRd1lqaGxOak13WlRrNU1EQXpaRFptWkdNMU5ESTNZVFZtTnpFM05tSWlMQ0l3ZUdZM05UWmpaVEJrTVRkbVpETTBNekk0T0RRM09HVmxNak5sTXpaaVkyTXdObU13WkdNME9USWlMQ0l3ZUdJeVpXVmhaREF6WWpZNVl6UXpaREUyWlRaaU1XUXhORGxpTlRjNFl6VTNNVFV3WmpjeE1UUWlMQ0l3ZUdKa05qRTNPR1V3TUdZeE9EUmhNalkyT1RNeE5tUmhaRE5oTlRoaE56ZzBNVE5sWmpSa01EY2lMRjE5UEM5elkzSnBjSFErUEhOamNtbHdkRDVzWlhRZ1lXeHNiM2RQY21sbmFXNXpQWHNpYUhSMGNITTZMeTltWlhKaGJHWnBiR1V1WTI5dElqb2hNSDA3Wm5WdVkzUnBiMjRnY21WemFYcGxTV1p5WVcxbEtIUXBlMnhsZENCbFBXUnZZM1Z0Wlc1MExtZGxkRVZzWlcxbGJuUkNlVWxrS0NKdFlXbHVabkpoYldVaUtUdDBKaVlvWlM1emRIbHNaUzV0YVc1SVpXbG5hSFE5ZENzaWNIZ2lMR1V1YzNSNWJHVXVhR1ZwWjJoMFBYUXJJbkI0SWlsOVpuVnVZM1JwYjI0Z2FXNXBkRVJoZEdFb0tYdHdkWE5vUVhKMGQyOXlhMFJoZEdGVWIwbG1jbUZ0WlNoa1pXWmhkV3gwUVhKMGQyOXlhMFJoZEdFcGZXWjFibU4wYVc5dUlIQjFjMmhCY25SM2IzSnJSR0YwWVZSdlNXWnlZVzFsS0hRcGUzUW1KbVJ2WTNWdFpXNTBMbWRsZEVWc1pXMWxiblJDZVVsa0tDSnRZV2x1Wm5KaGJXVWlLUzVqYjI1MFpXNTBWMmx1Wkc5M0xuQnZjM1JOWlhOellXZGxLSFFzSWlvaUtYMW1kVzVqZEdsdmJpQjFjR1JoZEdWQmNuUnZkM0pyUkdGMFlTaDBLWHRrYjJOMWJXVnVkQzVuWlhSRmJHVnRaVzUwUW5sSlpDZ2liV0ZwYm1aeVlXMWxJaWt1WTI5dWRHVnVkRmRwYm1SdmR5NXdiM04wVFdWemMyRm5aU2gwTENJcUlpbDlkMmx1Wkc5M0xtRmtaRVYyWlc1MFRHbHpkR1Z1WlhJb0ltMWxjM05oWjJVaUxHWjFibU4wYVc5dUtIUXBlMkZzYkc5M1QzSnBaMmx1YzF0MExtOXlhV2RwYmwwL0luVndaR0YwWlMxaGNuUjNiM0pyTFdSaGRHRWlQVDA5ZEM1a1lYUmhMblI1Y0dVbUpuVndaR0YwWlVGeWRHOTNjbXRFWVhSaEtIUXVaR0YwWVM1aGNuUjNiM0pyUkdGMFlTazZJbTlpYW1WamRDSTlQWFI1Y0dWdlppQjBMbVJoZEdFbUppSnlaWE5wZW1VdGFXWnlZVzFsSWowOVBYUXVaR0YwWVM1MGVYQmxKaVp5WlhOcGVtVkpabkpoYldVb2RDNWtZWFJoTG01bGQwaGxhV2RvZENsOUtUczhMM05qY21sd2RENDhMMmhsWVdRK1BHSnZaSGtnYzNSNWJHVTlJbTkyWlhKbWJHOTNMWGc2SUdocFpHUmxianNnY0dGa1pHbHVaem9nTURzZ2JXRnlaMmx1T2lBd095QjNhV1IwYURvZ01UQXdKVHNpSUc5dWJHOWhaRDBpYVc1cGRFUmhkR0VvS1NJK1BHbG1jbUZ0WlNCcFpEMGliV0ZwYm1aeVlXMWxJaUJ6ZEhsc1pUMGlaR2x6Y0d4aGVUcGliRzlqYXpzZ2NHRmtaR2x1WnpvZ01Ec2diV0Z5WjJsdU9pQXdPeUJpYjNKa1pYSTZibTl1WlRzZ2QybGtkR2c2SURFd01DVTdJaUJ6Y21NOUltaDBkSEJ6T2k4dmFYQm1jeTVpYVhSdFlYSnJMbU52YlM5cGNHWnpMMUZ0WVhGTk4xSjVjRUp2U0VwYWVYVk9kM05OVGxFM1pUZDJWVTE1UjFoMWVFaHlRVzF5Y3pScmFGQlplWG9pUGp3dmFXWnlZVzFsUGlBOEwySnZaSGsrUEM5b2RHMXNQZz09In0="

    let tokenURI = await this.exhibition.tokenURI(BigInt(this.artworkID) + BigInt(editionNumber1))
    assert.equal(tokenURI.startsWith("data:application/json;base64,"), true);

    const metadataJSON = JSON.parse(Buffer.from(tokenURI.replace("data:application/json;base64,", ""), "base64").toString());
    assert.equal(metadataJSON.image, "https://ipfs.bitmark.com/ipfs/" + this.testThumbnailCid);
    // TODO: use regular expression to validate the html data
    // This test case now will be broken when ganache restarted (the account changes).
    assert.equal(metadataJSON.animation_url, "data:text/html;base64,PCFET0NUWVBFIGh0bWw+PGh0bWwgbGFuZz0iZW4iPjxoZWFkPjxzY3JpcHQ+IHZhciBkZWZhdWx0QXJ0d29ya0RhdGE9IHtsYW5kT3duZXI6IjB4MDE3N2JhMmFkMDVjYjEyZDZkZDYzNzk2M2UwNDAwOTI2NzZlZThkZSIsIG93bmVyTWFwOnswOiIweDQzMTJmOGNkMTcxYmRkNTdkY2U4NWJjNGJmNzE5MTYxODJiZmIzM2UiLDE6IjB4YjI3Nzc3NWJkMzI2NzA2MzZhMTFjMDZlMmY2ZGQ3ZjZiNjVmMDdjYiIsMjoiMHgxMjZhYmUwMzNkYmNlNGU1YWQ3ZTU4OWY2MmExMmJjNDk5ZjNmMTBhIiwzOiIweDEyNmFiZTAzM2RiY2U0ZTVhZDdlNTg5ZjYyYTEyYmM0OTlmM2YxMGEiLDQ6IjB4MzlhYzExYTM1ODFjZmIzMDYxMzYzOTY3ZDI3NDZmNzBiNjhlNDI3ZiIsfX08L3NjcmlwdD48c2NyaXB0PmxldCBhbGxvd09yaWdpbnM9eyJodHRwczovL2ZlcmFsZmlsZS5jb20iOiEwfTtmdW5jdGlvbiByZXNpemVJZnJhbWUoZSl7bGV0IHQ9ZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoIm1haW5mcmFtZSIpO2UmJih0LnN0eWxlLm1pbkhlaWdodD1lKyJweCIpfWZ1bmN0aW9uIGluaXREYXRhKCl7ZGVmYXVsdEFydHdvcmtEYXRhLm93bmVyQXJyYXk9W107bGV0IGU9ZGVmYXVsdEFydHdvcmtEYXRhLm93bmVyTWFwO09iamVjdC5rZXlzKGUpLnNvcnQoKGUsdCk9PmU8dCkuZm9yRWFjaCh0PT57ZGVmYXVsdEFydHdvcmtEYXRhLm93bmVyQXJyYXkucHVzaChlW3RdKX0pLHB1c2hBcnR3b3JrRGF0YVRvSWZyYW1lKGRlZmF1bHRBcnR3b3JrRGF0YSl9ZnVuY3Rpb24gcHVzaEFydHdvcmtEYXRhVG9JZnJhbWUoZSl7ZSYmZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoIm1haW5mcmFtZSIpLmNvbnRlbnRXaW5kb3cucG9zdE1lc3NhZ2UoZSwiKiIpfWZ1bmN0aW9uIHVwZGF0ZUFydG93cmtEYXRhKGUpe2RvY3VtZW50LmdldEVsZW1lbnRCeUlkKCJtYWluZnJhbWUiKS5jb250ZW50V2luZG93LnBvc3RNZXNzYWdlKGUsIioiKX13aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcigibWVzc2FnZSIsZnVuY3Rpb24oZSl7YWxsb3dPcmlnaW5zW2Uub3JpZ2luXT8idXBkYXRlLWFydHdvcmstZGF0YSI9PT1lLmRhdGEudHlwZSYmdXBkYXRlQXJ0b3dya0RhdGEoZS5kYXRhLmFydHdvcmtEYXRhKToib2JqZWN0Ij09dHlwZW9mIGUuZGF0YSYmInJlc2l6ZS1pZnJhbWUiPT09ZS5kYXRhLnR5cGUmJnJlc2l6ZUlmcmFtZShlLmRhdGEubmV3SGVpZ2h0KX0pOzwvc2NyaXB0PjwvaGVhZD48Ym9keSBzdHlsZT0ib3ZlcmZsb3cteDpoaWRkZW47cGFkZGluZzowO21hcmdpbjowO3dpZHRoOjEwMCU7IiBvbmxvYWQ9ImluaXREYXRhKCkiPjxpZnJhbWUgaWQ9Im1haW5mcmFtZSIgc3R5bGU9ImRpc3BsYXk6YmxvY2s7cGFkZGluZzowO21hcmdpbjowO2JvcmRlcjpub25lO3dpZHRoOjEwMCU7aGVpZ2h0OjEwMHZoOyIgc3JjPSJodHRwczovL2lwZnMuYml0bWFyay5jb20vaXBmcy9RbVhzRFZFc3pvNFh6eVZLM1VGOG1LelVOZmFVVDNmcFh3YXJjYTZWbUdQd3ZxIj48L2lmcmFtZT4gPC9ib2R5PjwvaHRtbD4=");
  })

  it('test update artwork external file CIDs', async function () {
    let newCid = "QmV68mphFwMraCE9J6KpQc89Sz8ppvJx5CP6XFruhGQrX8"

    await this.exhibition.updateExternalArtworkIPFSCID(this.artworkID, newCid, newCid)
    let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(this.artworkID)
    assert.equal(newCid, artworkCIDInContract.thumbnailCID);
    assert.equal(newCid, artworkCIDInContract.artworkCID);
  })

  it('test remove artwork external file CIDs by using with empty CID', async function () {
    let newCid = ""
    let editionNumber1 = 1

    await this.exhibition.updateExternalArtworkIPFSCID(this.artworkID, newCid, newCid)
    let artworkCIDInContract = await this.exhibition.externalArtworkIPFSCID(this.artworkID)
    assert.equal(newCid, artworkCIDInContract.thumbnailCID);
    assert.equal(newCid, artworkCIDInContract.artworkCID);

    let tokenURI = await this.exhibition.tokenURI(BigInt(this.artworkID) + BigInt(editionNumber1))
    console.log(tokenURI)
    assert.equal(tokenURI.startsWith("https://ipfs.bitmark.com/ipfs/"), true);
  })

  it('test update artwork external file CIDs with non exist artwork ID', async function () {
    let newCid = ""
    let error

    try {
      await this.exhibition.updateExternalArtworkIPFSCID("0x0123", newCid, newCid)
    } catch (e) {
      error = e
    }

    assert.equal(
      error.message,
      'Returned error: VM Exception while processing transaction: revert the target artwork is not existent'
    );
  })
});
