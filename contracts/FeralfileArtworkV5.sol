// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "./Authorizable.sol";
import "./FeralfileSaleDataV2.sol";
import "./ECDSASigner.sol";

import "./FeralfileVaultV2.sol";

contract FeralfileExhibitionV5 is
    ERC1155,
    Authorizable,
    FeralfileSaleDataV2,
    ECDSASigner
{
    using Strings for uint256;

    struct Artwork {
        uint256 seriesId;
        uint256 tokenId;
        uint256 amount;
    }

    struct MintData {
        uint256 seriesId;
        uint256 tokenId;
        address owner;
        uint256 amount;
    }

    struct BurnData {
        address from;
        uint256 tokenId;
        uint256 amount;
    }

    // version code of contract
    string public constant codeVersion = "FeralfileExhibitionV5";

    // contract URI
    string public contractURI;

    // burnable
    bool public burnable;

    // selling
    bool public selling;

    // mintable
    bool public mintable = true;

    // cost receiver
    address public costReceiver;

    // vault contract instance
    IFeralfileVaultV2 public vault;

    // mapping from series ID to max supply of the series
    mapping(uint256 => uint256) internal _seriesMaxSupplies;

    // mapping from series ID to max supply of artwork belongs to the series
    mapping(uint256 => uint256) internal _seriesArtworkMaxSupplies;

    // mapping from series ID to total supply of the series
    mapping(uint256 => uint256) internal _seriesTotalSupplies;

    // mapping from token ID representing the artwork to total supply of the artwork
    mapping(uint256 => uint256) internal _artworkTotalSupplies;

    // mapping from token ID to Artwork
    mapping(uint256 => Artwork) internal _allArtworks;

    // mapping from owner to an array of token ID
    mapping(address => uint256[]) internal _ownedTokens;

    // mapping from token ID to index of the owner token list
    mapping(uint256 => uint256) private _ownedTokensIndex;

    constructor(
        string memory baseTokenURI_,
        string memory contractURI_,
        address signer_,
        address vault_,
        address costReceiver_,
        bool burnable_,
        uint256[] memory seriesIds_,
        uint256[] memory seriesMaxSupplies_,
        uint256[] memory seriesArtworkMaxSupplies_
    ) ERC1155(baseTokenURI_) ECDSASigner(signer_) {
        require(
            bytes(baseTokenURI_).length > 0,
            "FeralfileExhibitionV5: baseTokenURI_ is empty"
        );
        require(
            bytes(contractURI_).length > 0,
            "FeralfileExhibitionV5: contractURI_ is empty"
        );
        require(
            vault_ != address(0),
            "FeralfileExhibitionV5: vault_ is zero address"
        );
        require(
            costReceiver_ != address(0),
            "FeralfileExhibitionV5: costReceiver_ is zero address"
        );
        require(
            seriesIds_.length > 0,
            "FeralfileExhibitionV5: seriesIds_ is empty"
        );
        require(
            seriesMaxSupplies_.length > 0,
            "FeralfileExhibitionV5: _seriesMaxSupplies is empty"
        );
        require(
            seriesArtworkMaxSupplies_.length > 0,
            "FeralfileExhibitionV5: seriesArtworkMaxSupplies_ is empty"
        );
        require(
            seriesIds_.length == seriesMaxSupplies_.length &&
                seriesIds_.length == seriesArtworkMaxSupplies_.length,
            "FeralfileExhibitionV5: seriesMaxSupplies_ and seriesIds_ lengths are not the same"
        );

        contractURI = contractURI_;
        costReceiver = costReceiver_;
        vault = IFeralfileVaultV2(payable(vault_));
        burnable = burnable_;

        // initialize max supply map
        for (uint256 i = 0; i < seriesIds_.length; i++) {
            // Check invalid series ID
            require(seriesIds_[i] > 0, "FeralfileExhibitionV5: zero seriesId");
            // Check invalid max supply
            require(
                seriesMaxSupplies_[i] > 0,
                "FeralfileExhibitionV5: zero series max supply"
            );
            require(
                seriesArtworkMaxSupplies_[i] > 0,
                "FeralfileExhibitionV5: zero artwork max supply"
            );

            // Check duplicate with others
            require(
                _seriesMaxSupplies[seriesIds_[i]] == 0,
                "FeralfileExhibitionV5: duplicate seriesId"
            );
            require(
                _seriesArtworkMaxSupplies[seriesIds_[i]] == 0,
                "FeralfileExhibitionV5: duplicate seriesId"
            );

            _seriesMaxSupplies[seriesIds_[i]] = seriesMaxSupplies_[i];
            _seriesArtworkMaxSupplies[
                seriesIds_[i]
            ] = seriesArtworkMaxSupplies_[i];
        }
    }

    /// @notice Get all Artworks of an owner
    /// @param owner_ an address of the owner
    function artworkOf(
        address owner_
    ) external view returns (Artwork[] memory) {
        uint256[] memory tokens = _ownedTokens[owner_];
        Artwork[] memory artworks = new Artwork[](tokens.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            artworks[i] = _allArtworks[tokens[i]];
        }
        return artworks;
    }

    /// @notice Get all token IDs of an owner
    /// @param owner_ an address of the owner
    function tokenOf(address owner_) external view returns (uint256[] memory) {
        return _ownedTokens[owner_];
    }

    /// @notice Get series max supply
    /// @param seriesId_ a series ID
    /// @return uint256 the max supply
    function seriesMaxSupply(
        uint256 seriesId_
    ) external view virtual returns (uint256) {
        return _seriesMaxSupplies[seriesId_];
    }

    /// @notice Get series artwork max supply
    /// @param seriesId_ a series ID
    function seriesArtworkMaxSupply(
        uint256 seriesId_
    ) external view virtual returns (uint256) {
        return _seriesArtworkMaxSupplies[seriesId_];
    }

    /// @notice Get series total supply
    /// @param seriesId_ a series ID
    /// @return uint256 the total supply
    function seriesTotalSupply(
        uint256 seriesId_
    ) external view virtual returns (uint256) {
        return _seriesTotalSupplies[seriesId_];
    }

    /// @notice Get artwork total supply
    /// @param tokenId_ a token ID representing the artwork
    /// @return uint256 the total supply
    function artworkTotalSupply(
        uint256 tokenId_
    ) external view virtual returns (uint256) {
        return _artworkTotalSupplies[tokenId_];
    }

    /// @notice Get artwork data
    /// @param tokenId_ a token ID representing the artwork
    /// @return Artwork the Artwork object
    function getArtwork(
        uint256 tokenId_
    ) external view virtual returns (Artwork memory) {
        return _allArtworks[tokenId_];
    }

    /// @notice Set vault contract
    /// @param vault_ - the address of vault contract
    /// @dev don't allow to set vault as zero address
    function setVault(address vault_) external onlyOwner {
        require(
            vault_ != address(0),
            "FeralfileExhibitionV5: vault_ is zero address"
        );
        vault = IFeralfileVaultV2(payable(vault_));
    }

    /// @notice set contract URI
    /// @param contractURI_ contract URI
    function setContractURI(string memory contractURI_) external onlyOwner {
        contractURI = contractURI_;
        emit ContractURIUpdated();
    }

    /// @notice set base token URI
    /// @param uri_ token URI
    function setBaseTokenURI(string memory uri_) external onlyOwner {
        require(bytes(uri_).length > 0, "FeralfileExhibitionV5: uri_ is empty");
        _setURI(uri_);
    }

    /// @notice the cost receiver address
    /// @param costReceiver_ - the address of cost receiver
    function setCostReceiver(address costReceiver_) external onlyOwner {
        require(
            costReceiver_ != address(0),
            "FeralfileExhibitionV5: costReceiver_ is zero address"
        );
        costReceiver = costReceiver_;
    }

    /// @notice Start artwork sale
    function startSale() external onlyOwner {
        mintable = false;
        resumeSale();
    }

    /// @notice Resume artwork sale
    function resumeSale() public onlyOwner {
        require(
            !mintable,
            "FeralfileExhibitionV5: mintable required to be false"
        );
        require(
            !selling,
            "FeralfileExhibitionV5: selling required to be false"
        );
        require(
            _ownedTokens[address(this)].length > 0,
            "FeralfileExhibitionV5: no more artwork to sell"
        );

        selling = true;
    }

    /// @notice Pause artwork sale
    function pauseSale() public onlyOwner {
        require(
            !mintable,
            "FeralfileExhibitionV5: mintable required to be false"
        );
        require(selling, "FeralfileExhibitionV5: selling required to be true");
        selling = false;
    }

    /// @notice burn unsold artworks
    /// @param limit_ the maximum number of unsold artworks to burn
    /// @dev the limit_ is used to prevent the function from running out of gas
    function burnUnsoldArtworks(uint256 limit_) external onlyOwner {
        require(limit_ > 0, "FeralfileExhibitionV5: limit_ is zero");
        require(
            !mintable,
            "FeralfileExhibitionV5: mintable required to be false"
        );
        require(
            !selling,
            "FeralfileExhibitionV5: selling required to be false"
        );

        // get all token IDs of the contract
        uint256[] memory tokenIds = _ownedTokens[address(this)];
        if (tokenIds.length == 0) {
            return;
        }
        if (tokenIds.length < limit_) {
            limit_ = tokenIds.length;
        }

        // burn data
        BurnData[] memory data = new BurnData[](limit_);
        for (uint256 i = 0; i < limit_; i++) {
            data[i] = BurnData({
                from: address(this),
                tokenId: tokenIds[i],
                amount: balanceOf(address(this), tokenIds[i])
            });
        }

        // burn artworks
        for (uint256 i = 0; i < data.length; i++) {
            _burnArtwork(data[i].from, data[i].tokenId, data[i].amount);
        }
    }

    /// @notice transfer unsold artworks to a destination address
    /// @param tokenIds_ an array of token IDs
    /// @param to_ the destination address
    function transferUnsoldArtworks(
        uint256[] calldata tokenIds_,
        address to_
    ) external onlyOwner {
        require(
            to_ != address(0),
            "FeralfileExhibitionV5: to_ is zero address"
        );
        require(
            tokenIds_.length > 0,
            "FeralfileExhibitionV5: tokenIds_ is empty"
        );
        require(
            !mintable,
            "FeralfileExhibitionV5: mintable required to be false"
        );
        require(
            !selling,
            "FeralfileExhibitionV5: selling required to be false"
        );

        uint256[] memory amounts = new uint256[](tokenIds_.length);
        for (uint256 i = 0; i < tokenIds_.length; i++) {
            uint256 amount = balanceOf(address(this), tokenIds_[i]);
            amounts[i] = amount;
        }

        _safeBatchTransferFrom(address(this), to_, tokenIds_, amounts, "");
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
    ) external payable {
        // validation
        require(selling, "FeralfileExhibitionV5: sale is not started");
        require(
            _ownedTokens[address(this)].length > 0,
            "FeralfileExhibitionV5: no artwork to sell"
        );
        validateSaleData(saleData_);

        // payment
        saleData_.payByVaultContract
            ? vault.payForSale(r_, s_, v_, saleData_)
            : require(
                saleData_.price == msg.value,
                "FeralfileExhibitionV5: invalid payment amount"
            );

        // validate signature
        bytes32 message = keccak256(
            abi.encode(block.chainid, address(this), saleData_)
        );
        require(
            isValidSignature(message, r_, s_, v_),
            "FeralfileExhibitionV5: invalid signature"
        );

        // item royalty
        uint256 itemRevenue;
        if (saleData_.price > saleData_.cost) {
            itemRevenue =
                (saleData_.price - saleData_.cost) /
                saleData_.tokenIds.length;
        }

        uint256 distributedRevenue;
        uint256 platformRevenue;
        for (uint256 i = 0; i < saleData_.tokenIds.length; i++) {
            require(
                saleData_.tokenIds[i] != 0,
                "FeralfileExhibitionV5: token ID is zero"
            );

            // send token
            _safeTransferFrom(
                address(this),
                saleData_.destination,
                saleData_.tokenIds[i],
                1,
                ""
            ); // only support 1 token per transaction

            if (itemRevenue > 0) {
                // distribute royalty
                for (
                    uint256 j = 0;
                    j < saleData_.revenueShares[i].length;
                    j++
                ) {
                    uint256 rev = (itemRevenue *
                        saleData_.revenueShares[i][j].bps) / 10000;
                    if (
                        saleData_.revenueShares[i][j].recipient == costReceiver
                    ) {
                        platformRevenue += rev;
                        continue;
                    }
                    distributedRevenue += rev;
                    payable(saleData_.revenueShares[i][j].recipient).transfer(
                        rev
                    );
                }
            }

            emit BuyArtwork(
                saleData_.destination,
                saleData_.tokenIds[i],
                saleData_.biddingUnix
            );
        }

        require(
            saleData_.price - saleData_.cost >=
                distributedRevenue + platformRevenue,
            "FeralfileExhibitionV5: total bps over 10,000"
        );

        // Transfer cost, platform revenue and remaining funds
        uint256 leftOver = saleData_.price - distributedRevenue;
        if (leftOver > 0) {
            payable(costReceiver).transfer(leftOver);
        }
    }

    function _beforeTokenTransfer(
        address operator,
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory amounts,
        bytes memory data
    ) internal virtual override {
        super._beforeTokenTransfer(operator, from, to, ids, amounts, data);

        for (uint256 i = 0; i < ids.length; i++) {
            uint256 tokenId = ids[i];

            if (
                from != address(0) &&
                from != to &&
                balanceOf(from, tokenId) == amounts[i]
            ) {
                // only remove Artwork from Artwork enumeration if the current owned
                // token amount is equal to the amount of token to be transferred
                // otherwise, just update the amount of the token
                _removeTokenFromOwnerEnumeration(from, ids[i]);
            }
            if (
                to != address(0) &&
                to != from &&
                _ownedTokensIndex[tokenId] == 0
            ) {
                // only add Artwork to Artwork enumeration if the token is not owned by the receiver
                // otherwise, just update the amount of the token
                _addTokenToOwnerEnumeration(to, tokenId);
            }
        }
    }

    /// @dev Modify from {ERC721Enumerable}
    function _removeTokenFromOwnerEnumeration(
        address from,
        uint256 tokenId
    ) private {
        // To prevent a gap in from's tokens array, we store the last token in the index of the token to delete, and
        // then delete the last slot (swap and pop).
        uint256 lastTokenIndex = _ownedTokens[from].length - 1;
        uint256 tokenIndex = _ownedTokensIndex[tokenId];

        // When the token to delete is the last token, the swap operation is unnecessary
        if (tokenIndex != lastTokenIndex) {
            uint256 lastTokenId = _ownedTokens[from][lastTokenIndex];

            _ownedTokens[from][tokenIndex] = lastTokenId; // Move the last token to the slot of the to-delete token
            _ownedTokensIndex[lastTokenId] = tokenIndex; // Update the moved token's index
        }

        delete _ownedTokensIndex[tokenId];
        _ownedTokens[from].pop();
    }

    /// @dev Modify from {ERC721Enumerable}
    function _addTokenToOwnerEnumeration(address to, uint256 tokenId) private {
        uint256[] storage tokens = _ownedTokens[to];
        uint256 length = tokens.length;
        tokens.push(tokenId);
        _ownedTokensIndex[tokenId] = length;
    }

    /// @notice Mint new collection of Artwork
    /// @dev the function iterates over the array of MintData to call the internal function _mintArtwork
    /// @param data_ an array of MintData
    function mintArtworks(
        MintData[] calldata data_
    ) external virtual onlyAuthorized {
        require(
            mintable,
            "FeralfileExhibitionV5: contract doesn't allow to mint"
        );
        for (uint256 i = 0; i < data_.length; i++) {
            _mintArtwork(
                data_[i].seriesId,
                data_[i].tokenId,
                data_[i].owner,
                data_[i].amount
            );
        }
    }

    function _mintArtwork(
        uint256 seriesId_,
        uint256 tokenId_,
        address owner_,
        uint256 amount_
    ) internal {
        require(tokenId_ != 0, "FeralfileExhibitionV5: token ID is zero");
        require(amount_ != 0, "FeralfileExhibitionV5: amount is zero");

        Artwork memory artwork = _allArtworks[tokenId_];
        if (artwork.tokenId != 0) {
            // reassign seriesId for existing artwork to avoid bypass the max supply checks
            seriesId_ = artwork.seriesId;
        }

        // check series exists
        require(
            _seriesMaxSupplies[seriesId_] > 0,
            string(
                abi.encodePacked(
                    "FeralfileExhibitionV5: seriesId doesn't exist: ",
                    Strings.toString(seriesId_)
                )
            )
        );

        // check artwork total supplies
        require(
            _artworkTotalSupplies[tokenId_] + amount_ <=
                _seriesArtworkMaxSupplies[seriesId_],
            "FeralfileExhibitionV5: no more artwork slots available"
        );

        if (artwork.tokenId != 0) {
            // if artwork exists before
            // increase artwork amount
            _allArtworks[tokenId_].amount += amount_;
        } else {
            // if artwork doesn't exist before
            // 1.check series total supplies
            require(
                _seriesTotalSupplies[seriesId_] < _seriesMaxSupplies[seriesId_],
                "FeralfileExhibitionV5: no more series slots available"
            );

            // 2. increase series total supplies
            _seriesTotalSupplies[seriesId_] += 1;

            // 3. add artwork to allArtworks
            _allArtworks[tokenId_] = Artwork({
                seriesId: seriesId_,
                tokenId: tokenId_,
                amount: amount_
            });
        }

        // increase artwork total supplies
        _artworkTotalSupplies[tokenId_] += amount_;

        // mint token
        _mint(owner_, tokenId_, amount_, "");

        // emit event
        emit NewArtwork(owner_, seriesId_, tokenId_, amount_);
    }

    /// @notice Burn a collection of artworks
    /// @dev the function iterates over the array of token ID to call the internal function _burnArtwork
    /// @param data_ an array of BurnData
    function burnArtworks(BurnData[] calldata data_) external {
        require(burnable, "FeralfileExhibitionV5: token is not burnable");
        for (uint256 i = 0; i < data_.length; i++) {
            _burnArtwork(data_[i].from, data_[i].tokenId, data_[i].amount);
        }
    }

    function _burnArtwork(
        address from_,
        uint256 tokenId_,
        uint256 amount_
    ) internal {
        require(tokenId_ != 0, "FeralfileExhibitionV5: token ID is zero");
        require(amount_ != 0, "FeralfileExhibitionV5: amount is zero");
        require(
            from_ != address(0),
            "FeralfileExhibitionV5: from is zero address"
        );

        Artwork memory artwork = _allArtworks[tokenId_];
        require(
            artwork.tokenId != 0,
            "FeralfileExhibitionV5: artwork doesn't exist"
        );

        if (artwork.amount <= amount_) {
            // if burn whole token of artwork
            // 1. decrease series total supplies
            // 2. remove artwork from allArtworks
            _seriesTotalSupplies[artwork.seriesId] -= 1;
            delete _allArtworks[tokenId_];
        } else {
            // if burn part of token of artwork
            // just decrease artwork amount
            _allArtworks[tokenId_].amount -= amount_;
        }

        // decrease artwork total supplies
        _artworkTotalSupplies[tokenId_] -= amount_;

        // burn artwork
        _burn(from_, tokenId_, amount_);

        // emit event
        emit BurnArtwork(tokenId_, amount_);
    }

    /// @notice able to receive fund from vault contract
    receive() external payable {
        require(
            msg.sender == address(vault),
            "FeralfileExhibitionV5: only accept fund from vault contract."
        );
    }

    function onERC1155Received(
        address,
        address from_,
        uint256,
        uint256,
        bytes memory
    ) public pure returns (bytes4) {
        require(
            from_ == address(0),
            "FeralfileExhibitionV5: not allowed to send token back"
        );
        return this.onERC1155Received.selector;
    }

    /// @notice Event emitted when new Artwork has been minted
    event NewArtwork(
        address indexed owner_,
        uint256 indexed seriesId_,
        uint256 indexed tokenId_,
        uint256 amount_
    );

    /// @notice Event emitted when Artwork has been burned
    event BurnArtwork(uint256 indexed tokenId_, uint256 amount_);

    /// @notice Event emitted when Artwork has been sold
    event BuyArtwork(
        address indexed buyer_,
        uint256 indexed tokenId_,
        uint256 biddingUnix_
    );

    /// @notice Event emitted when contract URI has been updated
    event ContractURIUpdated();
}
