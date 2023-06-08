// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "./Authorizable.sol";
import "./UpdateableOperatorFilterer.sol";

contract Vault {
    function pay(uint256 weiAmount) external {}
}

contract FeralfileExhibitionV4 is
    ERC721,
    Authorizable,
    UpdateableOperatorFilterer
{
    using Strings for uint256;

    struct Artwork {
        uint256 seriesId;
        uint256 tokenId;
    }

    struct MintData {
        uint256 seriesId;
        uint256 tokenId;
        address owner;
    }

    struct Royalty {
        address recipient;
        uint256 bps;
    }

    struct SaleData {
        uint256 price; // in wei
        uint256 cost; // in wei
        uint256 expiryTime;
        address destination;
        uint256[] tokenIds;
        Royalty[][] royalties; // address and royalty bps (500 means 5%)
        bool payByVaultContract; // get eth from vault contract, used by credit card pay that proxy by ITX
    }

    // version code of contract
    string public constant codeVersion = "FeralfileExhibitionV4";

    // cost receiver
    address public costReceiver;

    // burnable
    bool public burnable;

    // bridgeable
    bool public bridgeable;

    // total supply
    uint256 public totalSupply;

    // mintable
    bool public mintable = true;

    // token base URI
    string public tokenBaseURI;

    // contract URI
    string public contractURI;

    // vault contract address
    address public vault;

    // signer
    address internal _signer;

    // series max supplies
    mapping(uint256 => uint256) internal _seriesMaxSupplies;

    // series total supplies
    mapping(uint256 => uint256) internal _seriesTotalSupplies;

    // all artworks
    mapping(uint256 => Artwork) internal _allArtworks;

    // selling
    bool private _selling;

    constructor(
        string memory name_,
        string memory symbol_,
        address signer_,
        bool burnable_,
        bool bridgeable_,
        address vault_,
        address costReceiver_,
        string memory tokenBaseURI_,
        string memory contractURI_,
        uint256[] memory seriesIds_,
        uint256[] memory seriesMaxSupplies_
    ) ERC721(name_, symbol_) {
        // validations
        require(
            bytes(name_).length > 0,
            "FeralfileExhibitionV4: name_ is empty"
        );
        require(
            bytes(symbol_).length > 0,
            "FeralfileExhibitionV4: symbol_ is empty"
        );
        require(
            vault_ != address(0),
            "FeralfileExhibitionV4: vaultAddress_ is zero address"
        );
        require(
            costReceiver_ != address(0),
            "FeralfileExhibitionV4: costReceiver_ is zero address"
        );
        require(
            bytes(tokenBaseURI_).length > 0,
            "FeralfileExhibitionV4: tokenBaseURI_ is empty"
        );
        require(
            bytes(contractURI_).length > 0,
            "FeralfileExhibitionV4: contractURI_ is empty"
        );
        require(
            signer_ != address(0),
            "FeralfileExhibitionV4: signer_ is zero address"
        );
        require(
            seriesIds_.length > 0,
            "FeralfileExhibitionV4: seriesIds_ is empty"
        );
        require(
            seriesMaxSupplies_.length > 0,
            "FeralfileExhibitionV4: _seriesMaxSupplies is empty"
        );
        require(
            seriesIds_.length == seriesMaxSupplies_.length,
            "FeralfileExhibitionV4: seriesMaxSupplies_ and seriesIds_ lengths are not the same"
        );

        burnable = burnable_;
        bridgeable = bridgeable_;
        costReceiver = costReceiver_;
        vault = vault_;
        tokenBaseURI = tokenBaseURI_;
        contractURI = contractURI_;
        _signer = signer_;

        // initialize max supply map
        for (uint256 i = 0; i < seriesIds_.length; i++) {
            _seriesMaxSupplies[seriesIds_[i]] = seriesMaxSupplies_[i];
            require(
                seriesMaxSupplies_[i] > 0,
                "FeralfileExhibitionV4: zero max supply"
            );
        }
    }

    /// @notice Set signer
    function setSigner(address signer_) external onlyOwner {
        require(
            signer_ != address(0),
            "FeralfileExhibitionV4: signer_ is zero address"
        );
        _signer = signer_;
    }

    function signer() external view returns (address) {
        return _signer;
    }

    /// @notice Get series max supply
    /// @param seriesId a series ID
    /// @return uint256 the max supply
    function seriesMaxSupply(
        uint256 seriesId
    ) public view virtual returns (uint256) {
        return _seriesMaxSupplies[seriesId];
    }

    /// @notice Get series total supply
    /// @param seriesId a series ID
    /// @return uint256 the total supply
    function seriesTotalSupply(
        uint256 seriesId
    ) public view virtual returns (uint256) {
        return _seriesTotalSupplies[seriesId];
    }

    /// @notice Get artwork data
    /// @param tokenId a token ID representing the artwork
    /// @return Artwork the Artwork object
    function getArtwork(
        uint256 tokenId
    ) public view virtual returns (Artwork memory) {
        require(_exists(tokenId), "ERC721: token doesn't exist");
        return _allArtworks[tokenId];
    }

    /// @notice Set vault contract
    /// @dev don't allow to set signer as zero address
    function setVault(address vault_) external onlyOwner {
        vault = vault_;
    }

    /// @notice Turn on token sale
    function startSale() external onlyOwner {
        mintable = false;
        _selling = true;
    }

    /// @notice Turn off token sale
    function stopSale() external onlyOwner {
        _selling = false;
    }

    /// @dev override for OperatorFilterRegistry
    function setApprovalForAll(
        address operator,
        bool approved
    ) public override(ERC721) onlyAllowedOperatorApproval(operator) {
        super.setApprovalForAll(operator, approved);
    }

    /// @dev override for OperatorFilterRegistry
    function approve(
        address operator,
        uint256 tokenId
    ) public override(ERC721) onlyAllowedOperatorApproval(operator) {
        super.approve(operator, tokenId);
    }

    /// @dev override for OperatorFilterRegistry
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override(ERC721) onlyAllowedOperator(from) {
        super.transferFrom(from, to, tokenId);
    }

    /// @dev override for OperatorFilterRegistry
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override(ERC721) onlyAllowedOperator(from) {
        super.safeTransferFrom(from, to, tokenId);
    }

    /// @dev override for OperatorFilterRegistry
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes memory data
    ) public override(ERC721) onlyAllowedOperator(from) {
        super.safeTransferFrom(from, to, tokenId, data);
    }

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    function tokenURI(
        uint256 tokenId
    ) public view virtual override returns (string memory) {
        require(
            bytes(tokenBaseURI).length > 0,
            "ERC721Metadata: _tokenBaseURI is empty"
        );
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );

        return string(abi.encodePacked(tokenBaseURI, "/", tokenId.toString()));
    }

    /// @notice Update the base URI for all tokens
    function setTokenBaseURI(string memory baseURI_) external onlyOwner {
        tokenBaseURI = baseURI_;
    }

    /// @notice the cost receiver address
    /// @param costReceiver_ - the address of cost receiver
    function setCostReceiver(address costReceiver_) external onlyOwner {
        costReceiver = costReceiver_;
    }

    /// @notice isValidRequest validates a message by ecrecover to ensure
    //          it is signed by owner of token.
    /// @param message_ - the raw message for signing
    /// @param owner_ - owner address of token
    /// @param r_ - part of signature for validating parameters integrity
    /// @param s_ - part of signature for validating parameters integrity
    /// @param v_ - part of signature for validating parameters integrity
    function verifySignature(
        bytes32 message_,
        address owner_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) internal pure returns (bool) {
        return
            ECDSA.recover(ECDSA.toEthSignedMessageHash(message_), v_, r_, s_) ==
            owner_;
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
        require(_selling, "FeralfileExhibitionV4: sale is not started");
        require(
            saleData_.tokenIds.length > 0,
            "FeralfileExhibitionV4: tokenIds is empty"
        );
        require(
            saleData_.tokenIds.length == saleData_.royalties.length,
            "FeralfileExhibitionV4: tokenIds and royalties length mismatch"
        );
        require(
            saleData_.expiryTime > block.timestamp,
            "FeralfileExhibitionV4: sale is expired"
        );

        saleData_.payByVaultContract
            ? Vault(payable(vault)).pay(saleData_.price)
            : require(
                saleData_.price == msg.value,
                "FeralfileExhibitionV4: invalid payment amount"
            );

        bytes32 requestHash = keccak256(
            abi.encode(block.chainid, address(this), saleData_)
        );

        require(
            verifySignature(requestHash, _signer, r_, s_, v_),
            "FeralfileExhibitionV4: invalid signature"
        );

        uint256 itemRevenue;
        if (saleData_.price > saleData_.cost) {
            itemRevenue =
                (saleData_.price - saleData_.cost) /
                saleData_.tokenIds.length;
        }

        uint256 distributedRevenue;
        for (uint256 i = 0; i < saleData_.tokenIds.length; i++) {
            // send NFT
            _safeTransfer(
                address(this),
                saleData_.destination,
                saleData_.tokenIds[i],
                ""
            );
            if (itemRevenue > 0) {
                // distribute royalty
                for (uint256 j = 0; j < saleData_.royalties[i].length; j++) {
                    uint256 rev = (itemRevenue *
                        saleData_.royalties[i][j].bps) / 10000;
                    distributedRevenue += rev;
                    payable(saleData_.royalties[i][j].recipient).transfer(rev);
                }
            }

            emit BuyArtwork(saleData_.destination, saleData_.tokenIds[i]);
        }

        // Transfer cost and remaining funds
        uint256 cost = saleData_.price - distributedRevenue;
        if (cost > 0) {
            payable(costReceiver).transfer(cost);
        }
    }

    /// @notice able to recieve funds from vault contract
    receive() external payable {
        require(
            msg.sender == vault,
            "FeralfileExhibitionV4: only accept fund from vault contract."
        );
    }

    /// @notice utility function for checking the series exists
    function _seriesExists(uint256 seriesId) internal view returns (bool) {
        return _seriesMaxSupplies[seriesId] > 0;
    }

    /// @notice Mint new collection of Artwork
    /// @dev the function iterates over the array of MintData to call the internal function _mintArtwork
    /// @param data an array of MintData
    function mintArtworks(
        MintData[] calldata data
    ) external virtual onlyAuthorized {
        for (uint256 i = 0; i < data.length; i++) {
            _mintArtwork(data[i].seriesId, data[i].tokenId, data[i].owner);
        }
    }

    function _mintArtwork(
        uint256 seriesId,
        uint256 tokenId,
        address owner
    ) internal {
        // pre-condition checks
        require(
            mintable,
            "FeralfileExhibitionV4: contract doesn't allow to mint"
        );
        require(
            _seriesExists(seriesId),
            string(
                abi.encodePacked(
                    "FeralfileExhibitionV4: seriesId doesn't exist: ",
                    Strings.toString(seriesId)
                )
            )
        );

        // mint
        totalSupply += 1;
        _seriesTotalSupplies[seriesId] += 1;
        _allArtworks[tokenId] = Artwork(seriesId, tokenId);
        _mint(owner, tokenId);
        require(
            _seriesTotalSupplies[seriesId] <= _seriesMaxSupplies[seriesId],
            "FeralfileExhibitionV4: no slots available"
        );

        // emit event
        emit NewArtwork(owner, seriesId, tokenId);
    }

    /// @notice Burn a collection of artworks
    /// @dev the function iterates over the array of token ID to call the internal function _burnArtwork
    /// @param tokenIds an array of token ID
    function burnArtworks(uint256[] memory tokenIds) external {
        for (uint256 i = 0; i < tokenIds.length; i++) {
            _burnArtwork(tokenIds[i]);
        }
    }

    function _burnArtwork(uint256 tokenId) internal {
        require(burnable, "FeralfileExhibitionV4: token is not burnable");
        require(_exists(tokenId), "ERC721: token doesn't exist");

        // burn artwork
        Artwork memory artwork = _allArtworks[tokenId];
        _seriesTotalSupplies[artwork.seriesId] -= 1;
        totalSupply -= 1;
        delete _allArtworks[tokenId];
        _burn(tokenId);

        // emit event
        emit BurnArtwork(tokenId);
    }

    /// @notice Event emitted when new Artwork has been minted
    event NewArtwork(
        address indexed owner,
        uint256 indexed seriesId,
        uint256 indexed tokenId
    );

    /// @notice Event emitted when Artwork has been burned
    event BurnArtwork(uint256 indexed tokenId);

    /// @notice Event emitted when Artwork has been sold
    event BuyArtwork(address indexed buyer, uint256 indexed tokenId);
}
