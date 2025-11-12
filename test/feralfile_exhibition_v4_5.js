const FeralfileExhibitionV4_5 = artifacts.require("FeralfileExhibitionV4_5");
const FeralfileVault = artifacts.require("FeralfileVault");

contract("FeralfileExhibitionV4_5", async (accounts) => {
  let exhibition;
  let vault;
  const owner = accounts[0];
  const signer = accounts[1];
  const costReceiver = accounts[3];
  
  // Exhibition ID without dashes (32 hex characters = 128 bits)
  const exhibitionId = "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6"; // 32 chars
  
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
      seriesMaxSupplies,
      exhibitionId
    );
  });

  it("should store the token ID prefix shard correctly", async () => {
    const storedShard = await exhibition.tokenIdPrefixShard();
    assert.equal(storedShard, exhibitionId, "Token ID prefix shard should match");
  });

  it("should calculate artwork index correctly for a token", async () => {
    // Simulate a token ID calculation
    // For seriesId = 1, artworkIndex = 5
    // part2 = 1 * 1000000 + 5 = 1000005 = 0xF4245
    // tokenId = (exhibitionId as uint128 << 128) | part2
    
    const seriesId = 1;
    const artworkIndex = 5;
    
    // Convert exhibition ID hex string to BigInt
    const exhibitionIdInt = BigInt("0x" + exhibitionId);
    const part2 = BigInt(seriesId * 1000000 + artworkIndex);
    
    // Construct tokenId: prefix in upper 128 bits, part2 in lower 128 bits
    const tokenId = (exhibitionIdInt << BigInt(128)) | part2;
    
    // First, we need to mint this token
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

  it("should calculate artwork indexes for multiple tokens", async () => {
    const seriesId = 1;
    const artworkIndexes = [0, 1, 2, 10, 99];
    
    const exhibitionIdInt = BigInt("0x" + exhibitionId);
    const tokenIds = [];
    const mintData = [];

    for (const artworkIndex of artworkIndexes) {
      const part2 = BigInt(seriesId * 1000000 + artworkIndex);
      const tokenId = (exhibitionIdInt << BigInt(128)) | part2;
      tokenIds.push(tokenId.toString());
      mintData.push({
        seriesId: seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test tokenIndexesByOwner - now returns TokenIndex[] struct array
    const tokenIndexMappings = await exhibition.tokenIndexesByOwner(tokenIds);
    
    for (let i = 0; i < artworkIndexes.length; i++) {
      assert.equal(
        tokenIndexMappings[i].tokenId.toString(),
        tokenIds[i],
        `Token ID ${i} should match`
      );
      assert.equal(
        tokenIndexMappings[i].index.toString(),
        artworkIndexes[i].toString(),
        `Artwork index ${i} should be ${artworkIndexes[i]}`
      );
    }
  });

  it("should handle different series IDs correctly", async () => {
    const testCases = [
      { seriesId: 1, artworkIndex: 5 },
      { seriesId: 2, artworkIndex: 10 },
      { seriesId: 1, artworkIndex: 50 },
      { seriesId: 2, artworkIndex: 100 },
    ];
    
    const exhibitionIdInt = BigInt("0x" + exhibitionId);
    const tokenIds = [];
    const mintData = [];

    for (const testCase of testCases) {
      const part2 = BigInt(testCase.seriesId * 1000000 + testCase.artworkIndex);
      const tokenId = (exhibitionIdInt << BigInt(128)) | part2;
      tokenIds.push(tokenId.toString());
      mintData.push({
        seriesId: testCase.seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test tokenIndexesByOwner - now returns TokenIndex[] struct array
    const tokenIndexMappings = await exhibition.tokenIndexesByOwner(tokenIds);
    
    for (let i = 0; i < testCases.length; i++) {
      assert.equal(
        tokenIndexMappings[i].tokenId.toString(),
        tokenIds[i],
        `Token ID ${i} should match`
      );
      assert.equal(
        tokenIndexMappings[i].index.toString(),
        testCases[i].artworkIndex.toString(),
        `Artwork index for series ${testCases[i].seriesId} should be ${testCases[i].artworkIndex}`
      );
    }
  });

  it("should reject invalid prefix shard length", async () => {
    const seriesIds = [1, 2];
    const seriesMaxSupplies = [100, 200];
    
    try {
      await FeralfileExhibitionV4_5.new(
        "Test Exhibition",
        "TEST",
        true,
        true,
        signer,
        vault.address,
        costReceiver,
        "ipfs://test-contract-uri",
        seriesIds,
        seriesMaxSupplies,
        "invalidlength" // Too short
      );
      assert.fail("Should have thrown an error");
    } catch (error) {
      assert(
        error.message.includes("invalid prefix shard length"),
        "Should reject invalid prefix shard length"
      );
    }
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

  it("should return correct tokenId to index mapping structure", async () => {
    const seriesId = 1;
    const artworkIndexes = [0, 5, 10];
    
    const exhibitionIdInt = BigInt("0x" + exhibitionId);
    const tokenIds = [];
    const mintData = [];

    for (const artworkIndex of artworkIndexes) {
      const part2 = BigInt(seriesId * 1000000 + artworkIndex);
      const tokenId = (exhibitionIdInt << BigInt(128)) | part2;
      tokenIds.push(tokenId.toString());
      mintData.push({
        seriesId: seriesId,
        tokenId: tokenId.toString(),
        owner: owner,
      });
    }

    await exhibition.mintArtworks(mintData, { from: owner });

    // Test that tokenIndexesByOwner returns proper mapping structure
    const mappings = await exhibition.tokenIndexesByOwner(tokenIds);
    
    assert.equal(mappings.length, tokenIds.length, "Should return same number of mappings as input tokens");
    
    // Verify each mapping contains both tokenId and index
    for (let i = 0; i < mappings.length; i++) {
      assert.isDefined(mappings[i].tokenId, `Mapping ${i} should have tokenId`);
      assert.isDefined(mappings[i].index, `Mapping ${i} should have index`);
      assert.equal(
        mappings[i].tokenId.toString(),
        tokenIds[i],
        `Mapping ${i} tokenId should match input`
      );
      assert.equal(
        mappings[i].index.toString(),
        artworkIndexes[i].toString(),
        `Mapping ${i} index should match expected artwork index`
      );
    }
  });
});

