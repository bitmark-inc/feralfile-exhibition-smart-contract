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
        address minter;
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
    bool private _burnable;

    // bridgeable
    bool private _bridgeable;

    // mintable
    bool private _mintable = true;

    // token base URI
    string internal _tokenBaseURI;

    // contract URI
    string private _contractURI;

    // vault contract address
    address private _vaultAddress;

    // default false and set to true when the sale starts
    bool private _isSelling = false;

    // signer
    address private _signer;

    // series max supplies
    mapping(uint256 => uint256) internal _seriesMaxSupplies;

    // series total supplies
    mapping(uint256 => uint256) internal _seriesTotalSupplies;

    // total supply
    uint256 private _totalSupply;

    // all artworks
    mapping(uint256 => Artwork) internal _allArtworks;

    constructor(
        string memory name_,
        string memory symbol_,
        address signer_,
        bool burnable_,
        bool bridgeable_,
        address vaultAddress_,
        address costReceiver_,
        string memory tokenBaseURI_,
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
            vaultAddress_ != address(0),
            "FeralfileExhibitionV4: vaultAddress_ is zero address"
        );
        require(
            bytes(tokenBaseURI_).length > 0,
            "FeralfileExhibitionV4: tokenBaseURI_ is empty"
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

        _signer = signer_;
        _burnable = burnable_;
        _bridgeable = bridgeable_;
        _vaultAddress = vaultAddress_;
        _tokenBaseURI = tokenBaseURI_;
        costReceiver = costReceiver_;

        // initialize max supply map
        for (uint256 i = 0; i < seriesIds_.length; i++) {
            _seriesMaxSupplies[seriesIds_[i]] = seriesMaxSupplies_[i];
            require(
                seriesMaxSupplies_[i] > 0,
                "FeralfileExhibitionV4: zero max supply"
            );
        }
    }

    function mintable() public view virtual returns (bool) {
        return _mintable;
    }

    function setMintable(bool mintable_) external onlyAuthorized {
        _mintable = mintable_;
    }

    function burnable() public view virtual returns (bool) {
        return _burnable;
    }

    function setBurnable(bool burnable_) external onlyAuthorized {
        _burnable = burnable_;
    }

    function bridgeable() public view virtual returns (bool) {
        return _bridgeable;
    }

    function setBridgeable(bool bridgeable_) external onlyAuthorized {
        _bridgeable = bridgeable_;
    }

    function signer() public view virtual returns (address) {
        return _signer;
    }

    function totalSupply() public view virtual returns (uint256) {
        return _totalSupply;
    }

    function seriesMaxSupply(
        uint256 seriesId
    ) public view virtual returns (uint256) {
        return _seriesMaxSupplies[seriesId];
    }

    function seriesTotalSupply(
        uint256 seriesId
    ) public view virtual returns (uint256) {
        return _seriesTotalSupplies[seriesId];
    }

    function getArtwork(
        uint256 tokenId
    ) public view virtual returns (Artwork memory) {
        require(_exists(tokenId), "ERC721: token doesn't exist");
        return _allArtworks[tokenId];
    }

    function setVaultContract(address vaultAddress_) external onlyOwner {
        _vaultAddress = vaultAddress_;
    }

    function startSale() external onlyOwner {
        _mintable = false;
        _isSelling = true;
    }

    function endSale() external onlyOwner {
        _isSelling = false;
    }

    function setApprovalForAll(
        address operator,
        bool approved
    ) public override(ERC721) onlyAllowedOperatorApproval(operator) {
        super.setApprovalForAll(operator, approved);
    }

    function approve(
        address operator,
        uint256 tokenId
    ) public override(ERC721) onlyAllowedOperatorApproval(operator) {
        super.approve(operator, tokenId);
    }

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override(ERC721) onlyAllowedOperator(from) {
        super.transferFrom(from, to, tokenId);
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override(ERC721) onlyAllowedOperator(from) {
        super.safeTransferFrom(from, to, tokenId);
    }

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
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );

        string memory baseURI = _tokenBaseURI;
        if (bytes(baseURI).length == 0) {
            baseURI = "ipfs://";
        }

        return string(abi.encodePacked(baseURI, "/", tokenId.toString()));
    }

    /// @notice Update the base URI for all tokens
    function setTokenBaseURI(string memory baseURI_) external onlyAuthorized {
        _tokenBaseURI = baseURI_;
    }

    /// @notice A URL for the opensea storefront-level metadata
    function contractURI() public view returns (string memory) {
        return _contractURI;
    }

    /// @notice the signer would sign the data of
    /// @param signer_ - the address of signer
    function setSigner(address signer_) external onlyOwner {
        signer = signer_;
    }

    /// @notice the vault contract address
    /// @param vault_ - the address of vault contract
    function setVaultContract(address vault_) external onlyOwner {
        vault = vault_;
    }

    /// @notice the cost receiver address
    /// @param costReceiver_ - the address of cost receiver
    function setCostReceiver(address costReceiver_) external onlyOwner {
        costReceiver = costReceiver_;
    }

    // @notice to start the sale
    function startSale() external onlyOwner {
        _canMint = false;
        _isSelling = true;
    }

    // @notice to end the sale
    function endSale() external onlyOwner {
        _isSelling = false;
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
        require(_isSelling, "FeralfileExhibitionV4: sale is not started");
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
            verifySignature(requestHash, signer, r_, s_, v_),
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

    // @notice able to recieve funds from vault contract
    receive() external payable {
        require(msg.sender == vault, "only accept fund from vault contract.");
    }

    function _seriesExists(uint256 seriesId) internal view returns (bool) {
        return _seriesMaxSupplies[seriesId] > 0;
    }

    function mintArtworks(
        MintData[] memory data
    ) external virtual onlyAuthorized {
        for (uint256 i = 0; i < data.length; i++) {
            _mintArtwork(data[i].seriesId, data[i].tokenId, data[i].minter);
        }
    }

    function _mintArtwork(
        uint256 seriesId,
        uint256 tokenId,
        address owner
    ) internal {
        // pre-condition checks
        require(
            _mintable,
            "FeralfileExhibitionV4: contract doesn't allow to mint"
        );
        require(
            owner != address(0),
            "FeralfileExhibitionV4: mint to zero address"
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
        require(!_exists(tokenId), "ERC721: token already minted");
        require(
            _seriesTotalSupplies[seriesId] < _seriesMaxSupplies[seriesId],
            "FeralfileExhibitionV4: no slots available"
        );

        // mint
        _totalSupply += 1;
        _seriesTotalSupplies[seriesId] += 1;
        _allArtworks[tokenId] = Artwork(seriesId, tokenId);
        _safeMint(owner, tokenId);

        // emit event
        emit NewArtwork(owner, seriesId, tokenId);
    }

    function burnArtworks(uint256[] memory tokenIds) external {
        for (uint256 i = 0; i < tokenIds.length; i++) {
            _burnArtwork(tokenIds[i]);
        }
    }

    function _burnArtwork(uint256 tokenId) internal {
        require(burnable(), "FeralfileExhibitionV4: token is not burnable");
        require(_exists(tokenId), "ERC721: token doesn't exist");

        // burn artwork
        Artwork memory artwork = _allArtworks[tokenId];
        _seriesTotalSupplies[artwork.seriesId] -= 1;
        _totalSupply -= 1;
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

    event BuyArtwork(address indexed buyer, uint256 indexed tokenId);
}
