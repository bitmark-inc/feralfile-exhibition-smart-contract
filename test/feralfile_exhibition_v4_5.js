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

  it("should calculate artwork index correctly for a token", async () => {
    const seriesId = 1;
    const artworkIndex = 5;
    const tokenId = generateTokenId(seriesId, artworkIndex);
    
    const mintData = [
      {
        seriesId: seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      },
    ];

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test getArtworkIndex
    const calculatedIndex = await exhibition.getArtworkIndex(tokenId.toString());
    assert.equal(
      calculatedIndex.toString(),
      artworkIndex.toString(),
      `Artwork index should be ${artworkIndex}`
    );
  });

  it("should return artwork indexes for all tokens in same series owned by same owner", async () => {
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

    // Test seriesArtworkIndexesOfOwner using the first token as reference
    const returnedIndexes = await exhibition.seriesArtworkIndexesOfOwner(tokenIds[0]);
    
    assert.equal(
      returnedIndexes.length,
      artworkIndexes.length,
      "Should return correct number of indexes"
    );
    
    // Sort both arrays for comparison (since order may not be guaranteed)
    const sortedReturned = returnedIndexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    const sortedExpected = artworkIndexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    
    // Verify all indexes match
    for (let i = 0; i < sortedExpected.length; i++) {
      assert.equal(
        sortedReturned[i],
        sortedExpected[i],
        `Artwork index ${i} should match`
      );
    }
  });

  it("should only return indexes for tokens in the same series", async () => {
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
      tokenIds.push(tokenId.toString());
      mintData.push({
        seriesId: testCase.seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test seriesArtworkIndexesOfOwner for series 1 (should return indexes 5 and 50)
    const series1Indexes = await exhibition.seriesArtworkIndexesOfOwner(tokenIds[0]);
    assert.equal(series1Indexes.length, 2, "Should return 2 indexes for series 1");
    
    const sortedSeries1 = series1Indexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    assert.equal(sortedSeries1[0], "5", "First index should be 5");
    assert.equal(sortedSeries1[1], "50", "Second index should be 50");

    // Test seriesArtworkIndexesOfOwner for series 2 (should return indexes 10 and 100)
    const series2Indexes = await exhibition.seriesArtworkIndexesOfOwner(tokenIds[1]);
    assert.equal(series2Indexes.length, 2, "Should return 2 indexes for series 2");
    
    const sortedSeries2 = series2Indexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    assert.equal(sortedSeries2[0], "10", "First index should be 10");
    assert.equal(sortedSeries2[1], "100", "Second index should be 100");
  });

  it("should revert when querying non-existent token", async () => {
    const nonExistentTokenId = "999999999999999999999999999";
    
    try {
      await exhibition.getArtworkIndex(nonExistentTokenId);
      assert.fail("Should have thrown an error");
    } catch (error) {
      assert(
        error.message.includes("token does not exist"),
        "Should revert for non-existent token"
      );
    }
  });

  it("should only return indexes for tokens owned by the same owner", async () => {
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

    // Test seriesArtworkIndexesOfOwner for owner1 (should only return owner1's indexes)
    const returnedIndexes = await exhibition.seriesArtworkIndexesOfOwner(owner1TokenIds[0]);
    assert.equal(returnedIndexes.length, owner1Indexes.length, "Should return only owner1's indexes");
    
    const sortedReturned = returnedIndexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    const sortedExpected = owner1Indexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    
    for (let i = 0; i < sortedExpected.length; i++) {
      assert.equal(sortedReturned[i], sortedExpected[i], `Index ${i} should match`);
    }

    // Test seriesArtworkIndexesOfOwner for owner2 (should only return owner2's indexes)
    const owner2ReturnedIndexes = await exhibition.seriesArtworkIndexesOfOwner(owner2TokenIds[0]);
    assert.equal(owner2ReturnedIndexes.length, owner2Indexes.length, "Should return only owner2's indexes");
    
    const sortedOwner2Returned = owner2ReturnedIndexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    const sortedOwner2Expected = owner2Indexes.map(idx => idx.toString()).sort((a, b) => Number(a) - Number(b));
    
    for (let i = 0; i < sortedOwner2Expected.length; i++) {
      assert.equal(sortedOwner2Returned[i], sortedOwner2Expected[i], `Owner2 index ${i} should match`);
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

    // Test seriesArtworkIndexesOfOwner with single token
    const indexes = await exhibition.seriesArtworkIndexesOfOwner(tokenId.toString());
    assert.equal(indexes.length, 1, "Should return 1 index");
    assert.equal(indexes[0].toString(), artworkIndex.toString(), "Index should match");

    // Test getArtworkIndex
    const index = await exhibition.getArtworkIndex(tokenId.toString());
    assert.equal(index.toString(), artworkIndex.toString(), "getArtworkIndex should return correct index");
  });

  it("should calculate correct indexes for edge values", async () => {
    const testCases = [
      { seriesId: 1, artworkIndex: 0 }, // Minimum index
      { seriesId: 1, artworkIndex: 999999 }, // Near maximum for multiplier
      { seriesId: 2, artworkIndex: 0 },
      { seriesId: 2, artworkIndex: 150 },
    ];
    
    const mintData = [];

    for (const testCase of testCases) {
      const tokenId = generateTokenId(testCase.seriesId, testCase.artworkIndex);
      mintData.push({
        seriesId: testCase.seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Verify each token's index
    for (const testCase of testCases) {
      const tokenId = generateTokenId(testCase.seriesId, testCase.artworkIndex);
      
      const calculatedIndex = await exhibition.getArtworkIndex(tokenId.toString());
      assert.equal(
        calculatedIndex.toString(),
        testCase.artworkIndex.toString(),
        `Index for series ${testCase.seriesId} should be ${testCase.artworkIndex}`
      );
    }
  });
});

