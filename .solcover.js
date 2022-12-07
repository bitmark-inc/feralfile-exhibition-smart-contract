module.exports = {
  client: require('ganache-cli'),
  skipFiles: ["FeralfileArtwork.sol", "FeralfileArtworkV2.sol", "FeralfileArtworkV3_1.sol", "Migrations.sol", "mocks/MockDecentraland.sol"],
};
