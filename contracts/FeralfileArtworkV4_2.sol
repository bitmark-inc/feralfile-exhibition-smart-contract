// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import { Strings } from "@openzeppelin/contracts/utils/Strings.sol"; 
import { Base64 } from "@openzeppelin/contracts/utils/Base64.sol";
import { FeralfileExhibitionV4_1 } from "./FeralfileArtworkV4_1.sol";
import { FeralfileToken } from "./FeralfileToken.sol";

contract FeralfileExhibitionV4_2 is FeralfileExhibitionV4_1 {

    error TokenIDNotFound();
    error MintNotEnabled();
    error FunctionNotSupported();
    error EmptyURI();

    struct TokenInfo {
        string imageURI;
        bytes parameters;
        uint256 coinAmount;
    }

    struct MintDataWithIndex {
        uint256 seriesId;
        uint256 tokenId;
        address owner;
        uint256 tokenIndex;
    }

    FeralfileToken private _erc20Token;

    mapping(uint256 => TokenInfo) private _tokenInfos; // tokenID => tokenInfo
    mapping(uint256 => uint256) private _tokenIndexes; // tokenID => tokenIndex

    string public tokenPlaceholderAnimationURI;
    string public tokenPlaceholderImageURI;

    string public artworkFileURI;

    constructor(
        string memory name_,
        string memory symbol_,
        bool burnable_,
        bool bridgeable_,
        address signer_,
        address vault_,
        address costReceiver_,
        string memory contractURI_,
        uint256[] memory seriesIds_,
        uint256[] memory seriesMaxSupplies_,
        address erc20ContractAddress_
    )
        FeralfileExhibitionV4_1(
            name_,
            symbol_,
            burnable_,
            bridgeable_,
            signer_,
            vault_,
            costReceiver_,
            contractURI_,
            seriesIds_,
            seriesMaxSupplies_
        )
    {
        _erc20Token = FeralfileToken(erc20ContractAddress_);
    }
    
    /// @notice Mint new collection of Artwork
    /// @dev the function iterates over the array of MintDataWithIndex to call the internal function _mintArtwork
    /// @param data an array of MintDataWithIndex
    function mintArtworksWithIndex(
        MintDataWithIndex[] calldata data
    ) external onlyAuthorized {
        if (!mintable) {
            revert MintNotEnabled();
        }

        for (uint256 i = 0; i < data.length; i++) {
            _tokenIndexes[data[i].tokenId] = data[i].tokenIndex;
            _mintArtwork(data[i].seriesId, data[i].tokenId, data[i].owner);
        }
    }

    /// @notice override revert mint new collection of Artwork
    function mintArtworks(MintData[] calldata) external override view onlyAuthorized {
        revert FunctionNotSupported();
    }

    /// @notice Update the imageURI & parameters of an edition to a new value
    function updateTokenInformation(uint256 tokenId, string calldata imageURI, bytes calldata parameters, uint256 coinAmount)
        external
        onlyAuthorized
    {
        if (!_exists(tokenId)) {
            revert TokenIDNotFound();
        }

        TokenInfo memory tokenInfo = _tokenInfos[tokenId];
        _tokenInfos[tokenId] = TokenInfo(imageURI, parameters, coinAmount);

        // mint token for the owner
        uint256 mintAmount = coinAmount - tokenInfo.coinAmount;
        if (mintAmount > 0) {
            address owner = ownerOf(tokenId);
            _erc20Token.mint(owner, mintAmount * 10**18);
        }
    }

    /// @notice pay to get artworks to a destination address. The pricing, costs and other details is included in the saleData
    /// @param r_ - part of signature for validating parameters integrity
    /// @param s_ - part of signature for validating parameters integrity
    /// @param v_ - part of signature for validating parameters integrity
    /// @param saleData_ - the sale data
    function buyArtworks(
        bytes32 r_,
        bytes32 s_,
        uint8 v_,
        SaleData calldata saleData_
    ) external payable override {
        require(_selling, "FeralfileExhibitionV4: sale is not started");
        super._checkContractOwnedToken();
        validateSaleData(saleData_);

        saleData_.payByVaultContract
            ? vault.payForSale(r_, s_, v_, saleData_)
            : require(
                saleData_.price == msg.value,
                "FeralfileExhibitionV4: invalid payment amount"
            );

        bytes32 message = keccak256(
            abi.encode(block.chainid, address(this), saleData_)
        );

        if (!isValidSignature(message, r_, s_, v_)) {
            revert InvalidSignature();
        }

        uint256 itemRevenue;
        if (saleData_.price > saleData_.cost) {
            itemRevenue =
                (saleData_.price - saleData_.cost) /
                saleData_.tokenIds.length;
        }

        uint256 distributedRevenue;
        uint256 platformRevenue;
        uint256 coinsTransferAmount = 0;
        for (uint256 i = 0; i < saleData_.tokenIds.length; i++) {
            // send NFT
            _safeTransfer(
                address(this),
                saleData_.destination,
                saleData_.tokenIds[i],
                ""
            );
            // distribute royalty
            RevenueShare[] memory revenueShares = saleData_.revenueShares[i];
            uint256 remainingRev = itemRevenue;

            // deduct advances payment from revenue
            for (
                uint256 j = 0;
                j < revenueShares.length && remainingRev > 0;
                j++
            ) {
                uint256 remainingAdvanceAmount = advances[
                    revenueShares[j].recipient
                ];
                uint256 rev = remainingAdvanceAmount >= remainingRev
                    ? remainingRev
                    : remainingAdvanceAmount;
                platformRevenue += rev;
                advances[revenueShares[j].recipient] -= rev;
                remainingRev -= rev;
            }

            // distribute revenue
            if (remainingRev > 0) {
                for (uint256 j = 0; j < revenueShares.length; j++) {
                    address recipient = revenueShares[j].recipient;
                    uint256 rev = (remainingRev * revenueShares[j].bps) / 10000;
                    if (recipient == costReceiver) {
                        platformRevenue += rev;
                        continue;
                    }
                    distributedRevenue += rev;
                    payable(recipient).transfer(rev);
                }
            }

            TokenInfo memory tokenInfo = _tokenInfos[saleData_.tokenIds[i]];
            coinsTransferAmount += tokenInfo.coinAmount;

            emit BuyArtwork(saleData_.destination, saleData_.tokenIds[i]);
        }

        // Call the transfer function of the ERC20 contract
        if (coinsTransferAmount > 0 && _erc20Token.balanceOf(address(this)) > coinsTransferAmount * 10**18) {
            bool success = _erc20Token.transfer(saleData_.destination, coinsTransferAmount * 10**18);
            require(success, "FeralfileExhibitionV4: FeralfileToken transfer failed");
        }

        require(
            saleData_.price - saleData_.cost >=
                distributedRevenue + platformRevenue,
            "FeralfileExhibitionV4: total bps over 10,000"
        );

        // Transfer cost, platform revenue and remaining funds
        uint256 leftOver = saleData_.price - distributedRevenue;
        if (leftOver > 0) {
            payable(costReceiver).transfer(leftOver);
        }        
    }

    /// @notice Update the base URI for all tokens
    function setArtworkFileURI(string calldata uri) external virtual onlyOwner {
        if (bytes(uri).length == 0) {
            revert EmptyURI();
        }

        artworkFileURI = uri;
    }

    /// @notice Update tokenPlaceholderImageURI
    function setTokenPlaceholderImageURI(string calldata uri) external virtual onlyOwner {
        if (bytes(uri).length == 0) {
            revert EmptyURI();
        }

        tokenPlaceholderImageURI = uri;
    }

    /// @notice Update tokenPlaceholderAnimationURI
    function setTokenPlaceholderAnimationURI(string calldata uri) external virtual onlyOwner {
        if (bytes(uri).length == 0) {
            revert EmptyURI();
        }

        tokenPlaceholderAnimationURI = uri;
    }

    /// @notice _buildAnimationURI returns the animation url
    /// @param parameters - the token paramters
    function _buildAnimationURI(bytes memory parameters)
        private
        view
        returns (string memory)
    {
        if (parameters.length == 0) {
            return tokenPlaceholderAnimationURI;
        }

        return string(
            abi.encodePacked(
                "data:text/html;base64,",
                 _buildIframe(Base64.encode(parameters), artworkFileURI)
            )
        );
    }

    /// @notice _buildImageURI returns the image url
    /// @param imageURI - the token imageURI
    function _buildImageURI(string memory imageURI)
        private
        view
        returns (string memory)
    {
        if (bytes(imageURI).length == 0) {
            return tokenPlaceholderImageURI;
        }

        return imageURI;
    }

    /// @notice _buildIframe returns a base64 encoded data for ff-frame
    /// @param tokenParams - the paramteres which would bring into the artwork
    /// @param artworkURI - the artwork file URI to be loaded into the iframe
    function _buildIframe(string memory tokenParams, string memory artworkURI)
        private
        pure
        returns (string memory)
    {
        return 
            Base64.encode(
                abi.encodePacked(
                    "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\">",
                    "<script>var defaultArtworkData=\"",
                    tokenParams,
                    "\";let allowOrigins = {\"https://feralfile.com\": !0};function initData(){document.getElementById(\"mainframe\").contentWindow.postMessage(defaultArtworkData, \"*\");}</script>",
                    "</head><body style=\"overflow-x:hidden;padding:0;margin:0;width: 100%;\" onload=\"initData()\"><iframe id=\"mainframe\" style=\"display:block;padding:0;margin:0;border:none;width:100%;height:100vh;\" src=\"",
                    artworkURI,
                    "\"></iframe></body></html>"
                )
            );
    }

    /// @notice _buildDataURL returns a base64 encoded data for ff-frame
    /// @param tokenID - the token ID for building artwork data
    function _buildDataURL(uint256 tokenID) private view returns (string memory) {
        TokenInfo memory tokenInfo = _tokenInfos[tokenID];
        uint256 tokenIndex = _tokenIndexes[tokenID];

        string memory tokenName = string(
            abi.encodePacked(name(), " #", Strings.toString(tokenIndex + 1))
        );

        string memory json = string(
            abi.encodePacked(
                "{\"name\":\"", tokenName,
                "\", \"external_url\":\"https://feralfile.com\", \"image\":\"", _buildImageURI(tokenInfo.imageURI),
                "\", \"animation_url\":\"", _buildAnimationURI(tokenInfo.parameters),
                "\"}"
            )
        );

        return string(
            abi.encodePacked(
                "data:application/json;base64,", Base64.encode(bytes(json))
            )
        );
    }

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    function tokenURI(uint256 tokenId)
        public
        view
        virtual
        override
        returns (string memory)
    {
        if (!_exists(tokenId)) {
            revert TokenIDNotFound();
        }

        return _buildDataURL(tokenId);
    }
}
