const FeralfileExhibitionV4_5 = artifacts.require("FeralfileExhibitionV4_5");
const FeralfileVault = artifacts.require("FeralfileVault");

contract("FeralfileExhibitionV4_5", async (accounts) => {
  let exhibition;
  let vault;
  const owner = accounts[0];
  const owner2 = accounts[2];
  const signer = accounts[1];
  const costReceiver = accounts[3];
  
  // Helper function to generate tokenId from seriesId and artworkIndex
  // TokenId structure: any prefix in upper 128 bits + (seriesId * 1000000 + artworkIndex) in lower 128 bits
  const generateTokenId = (seriesId, artworkIndex) => {
    // For simplicity, we use 0 as prefix since the contract only uses lower 128 bits
    return BigInt(seriesId * 1000000 + artworkIndex);
  };
  
  beforeEach(async () => {
    // Deploy vault first
    vault = await FeralfileVault.new(signer);
    
    const seriesIds = [1, 2];
    const seriesMaxSupplies = [100, 200];
    
    exhibition = await FeralfileExhibitionV4_5.new(
      "Test Exhibition",
      "TEST",
      true, // burnable
      true, // bridgeable
      signer,
      vault.address,
      costReceiver,
      "ipfs://test-contract-uri",
      seriesIds,
      seriesMaxSupplies
    );
  });

  it("should return token IDs for all tokens in same series owned by same owner", async () => {
    const seriesId = 1;
    const artworkIndexes = [0, 1, 2, 10, 99];
    
    const tokenIds = [];
    const mintData = [];

    for (const artworkIndex of artworkIndexes) {
      const tokenId = generateTokenId(seriesId, artworkIndex);
      tokenIds.push(tokenId.toString());
      mintData.push({
        seriesId: seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test seriesArtworksOfOwner using the first token as reference
    const returnedTokenIds = await exhibition.seriesArtworksOfOwner(tokenIds[0]);
    
    assert.equal(
      returnedTokenIds.length,
      tokenIds.length,
      "Should return correct number of token IDs"
    );
    
    // Sort both arrays for comparison (since order may not be guaranteed)
    const sortedReturned = returnedTokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    const sortedExpected = tokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    
    // Verify all token IDs match
    for (let i = 0; i < sortedExpected.length; i++) {
      assert.equal(
        sortedReturned[i],
        sortedExpected[i],
        `Token ID ${i} should match`
      );
    }
  });

  it("should only return token IDs for tokens in the same series", async () => {
    const testCases = [
      { seriesId: 1, artworkIndex: 5 },
      { seriesId: 2, artworkIndex: 10 },
      { seriesId: 1, artworkIndex: 50 },
      { seriesId: 2, artworkIndex: 100 },
    ];
    
    const tokenIds = [];
    const mintData = [];

    for (const testCase of testCases) {
      const tokenId = generateTokenId(testCase.seriesId, testCase.artworkIndex);
      tokenIds.push({ tokenId: tokenId.toString(), ...testCase });
      mintData.push({
        seriesId: testCase.seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test seriesArtworksOfOwner for series 1 (should return token IDs for series 1)
    const series1TokenIds = await exhibition.seriesArtworksOfOwner(tokenIds[0].tokenId);
    assert.equal(series1TokenIds.length, 2, "Should return 2 token IDs for series 1");
    
    const sortedSeries1 = series1TokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    const expectedSeries1 = [tokenIds[0].tokenId, tokenIds[2].tokenId].sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    assert.equal(sortedSeries1[0], expectedSeries1[0], "First token ID should match");
    assert.equal(sortedSeries1[1], expectedSeries1[1], "Second token ID should match");

    // Test seriesArtworksOfOwner for series 2 (should return token IDs for series 2)
    const series2TokenIds = await exhibition.seriesArtworksOfOwner(tokenIds[1].tokenId);
    assert.equal(series2TokenIds.length, 2, "Should return 2 token IDs for series 2");
    
    const sortedSeries2 = series2TokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    const expectedSeries2 = [tokenIds[1].tokenId, tokenIds[3].tokenId].sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    assert.equal(sortedSeries2[0], expectedSeries2[0], "First token ID should match");
    assert.equal(sortedSeries2[1], expectedSeries2[1], "Second token ID should match");
  });

  it("should revert when querying non-existent token with seriesArtworksOfOwner", async () => {
    const nonExistentTokenId = "999999999999999999999999999";
    
    try {
      await exhibition.seriesArtworksOfOwner(nonExistentTokenId);
      assert.fail("Should have thrown an error");
    } catch (error) {
      assert(
        error.message.includes("token does not exist"),
        "Should revert for non-existent token"
      );
    }
  });

  it("should only return token IDs for tokens owned by the same owner", async () => {
    const seriesId = 1;
    const owner1Indexes = [0, 5, 10];
    const owner2Indexes = [15, 20];
    
    const owner1TokenIds = [];
    const owner2TokenIds = [];
    const mintData = [];

    // Mint tokens for owner1
    for (const artworkIndex of owner1Indexes) {
      const tokenId = generateTokenId(seriesId, artworkIndex);
      owner1TokenIds.push(tokenId.toString());
      mintData.push({
        seriesId: seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    // Mint tokens for owner2
    for (const artworkIndex of owner2Indexes) {
      const tokenId = generateTokenId(seriesId, artworkIndex);
      owner2TokenIds.push(tokenId.toString());
      mintData.push({
        seriesId: seriesId,
        tokenId: tokenId.toString(),
        owner: owner2,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test seriesArtworksOfOwner for owner1 (should only return owner1's token IDs)
    const returnedTokenIds = await exhibition.seriesArtworksOfOwner(owner1TokenIds[0]);
    assert.equal(returnedTokenIds.length, owner1TokenIds.length, "Should return only owner1's token IDs");
    
    const sortedReturned = returnedTokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    const sortedExpected = owner1TokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    
    for (let i = 0; i < sortedExpected.length; i++) {
      assert.equal(sortedReturned[i], sortedExpected[i], `Token ID ${i} should match`);
    }

    // Test seriesArtworksOfOwner for owner2 (should only return owner2's token IDs)
    const owner2ReturnedTokenIds = await exhibition.seriesArtworksOfOwner(owner2TokenIds[0]);
    assert.equal(owner2ReturnedTokenIds.length, owner2TokenIds.length, "Should return only owner2's token IDs");
    
    const sortedOwner2Returned = owner2ReturnedTokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    const sortedOwner2Expected = owner2TokenIds.map(id => id.toString()).sort((a, b) => BigInt(a) > BigInt(b) ? 1 : -1);
    
    for (let i = 0; i < sortedOwner2Expected.length; i++) {
      assert.equal(sortedOwner2Returned[i], sortedOwner2Expected[i], `Owner2 token ID ${i} should match`);
    }
  });

  it("should handle edge case with single token", async () => {
    const seriesId = 1;
    const artworkIndex = 42;
    const tokenId = generateTokenId(seriesId, artworkIndex);
    
    const mintData = [{
      seriesId: seriesId,
      tokenId: tokenId.toString(),
      owner: owner,
    }];

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test seriesArtworksOfOwner with single token
    const returnedTokenIds = await exhibition.seriesArtworksOfOwner(tokenId.toString());
    assert.equal(returnedTokenIds.length, 1, "Should return 1 token ID");
    assert.equal(returnedTokenIds[0].toString(), tokenId.toString(), "Token ID should match");
  });

  it("should handle edge values for token IDs", async () => {
    const testCases = [
      { seriesId: 1, artworkIndex: 0 }, // Minimum index
      { seriesId: 1, artworkIndex: 999999 }, // Near maximum for multiplier
      { seriesId: 2, artworkIndex: 0 },
      { seriesId: 2, artworkIndex: 150 },
    ];
    
    const mintData = [];
    const tokenIds = [];

    for (const testCase of testCases) {
      const tokenId = generateTokenId(testCase.seriesId, testCase.artworkIndex);
      tokenIds.push({ tokenId: tokenId.toString(), ...testCase });
      mintData.push({
        seriesId: testCase.seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Verify tokens can be queried successfully for each series
    const series1TokenIds = await exhibition.seriesArtworksOfOwner(tokenIds[0].tokenId);
    assert.equal(series1TokenIds.length, 2, "Should return 2 token IDs for series 1");
    
    const series2TokenIds = await exhibition.seriesArtworksOfOwner(tokenIds[2].tokenId);
    assert.equal(series2TokenIds.length, 2, "Should return 2 token IDs for series 2");
  });
});

