// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "./Authorizable.sol";
import "./UpdateableOperatorFilterer.sol";

contract Vault {
    function pay(uint256 weiAmount) external {}
}

contract FeralfileExhibitionV4 is
    IERC165,
    ERC721Enumerable,
    Authorizable,
    UpdateableOperatorFilterer
{
    using Strings for uint256;

    // artwork id multiple for each series
    uint256 public constant ARTWORK_ID_MULTIPLE = 1000000;

    // version code of contract
    string public constant codeVersion = "FeralfileExhibitionV4";

    // signer of buy artwork
    address public signer = address(0);

    // cost receiver
    address public costReceiver = address(0);

    // vault contract
    address public vault = address(0);

    // burnable
    bool public isBurnable;

    // bridgeable
    bool public isBridgeable;

    // token base URI
    string internal _tokenBaseURI;

    // contract URI
    string private _contractURI;

    // default true and set to false when the sale starts
    bool private _canMint = true;

    // default false and set to true when the sale starts
    bool private _isSelling = false;

    struct Artwork {
        uint256 seriesIndex;
        uint256 artworkIndex;
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

    mapping(uint256 => Artwork) public artworks; //  => tokenID => Artwork

    constructor(
        string memory name_,
        string memory symbol_,
        string memory contractURI_,
        string memory tokenBaseURI_,
        address signer_,
        address vault_,
        address costReceiver_,
        bool isBurnable_,
        bool isBridgeable_
    ) ERC721(name_, symbol_) {
        isBurnable = isBurnable_;
        isBridgeable = isBridgeable_;
        _contractURI = contractURI_;
        _tokenBaseURI = tokenBaseURI_;
        signer = signer_;
        vault = vault_;
        costReceiver = costReceiver_;
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual override(ERC721Enumerable, IERC165) returns (bool) {
        return
            interfaceId == type(IERC721Enumerable).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /// @notice Return a count of series registered in this exhibition
    function totalSeries() public view virtual returns (uint256) {}

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    function tokenURI(
        uint256 tokenId_
    ) public view virtual override returns (string memory) {
        require(
            _exists(tokenId_),
            "ERC721Metadata: URI query for nonexistent token"
        );

        string memory baseURI = _tokenBaseURI;
        return
            bytes(baseURI).length > 0
                ? string.concat(baseURI, "/", Strings.toString(tokenId_))
                : "";
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
    function isValidRequest(
        bytes32 message_,
        address owner_,
        bytes32 r_,
        bytes32 s_,
        uint8 v_
    ) internal pure returns (bool) {
        address reqSigner = ECDSA.recover(
            ECDSA.toEthSignedMessageHash(message_),
            v_,
            r_,
            s_
        );
        return reqSigner == owner_;
    }

    /// @notice batchMint is function mint array of tokens
    /// @param mintParams_ - the array of transfer parameters
    function mintArtworks(
        Artwork[] memory mintParams_
    ) external virtual onlyAuthorized {
        for (uint256 i = 0; i < mintParams_.length; i++) {
            _mintArtwork(
                mintParams_[i].seriesIndex,
                mintParams_[i].artworkIndex
            );
        }
    }

    /// @notice mint artworks
    /// @param seriesIndex_ - the index of series
    /// @param artworkIndex_ - the index of artwork in series
    function _mintArtwork(
        uint256 seriesIndex_,
        uint256 artworkIndex_
    ) internal {
        require(_canMint, "FeralfileExhibitionV4: not in minting stage");
        require(
            seriesIndex_ > 0,
            "FeralfileExhibitionV4: invalid series index"
        );

        uint256 artworkID = seriesIndex_ * ARTWORK_ID_MULTIPLE + artworkIndex_;
        _mint(address(this), artworkID);
        artworks[artworkID] = Artwork(seriesIndex_, artworkIndex_);

        emit NewArtwork(address(this), artworkIndex_, artworkID);
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
            isValidRequest(requestHash, signer, r_, s_, v_),
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

    /// @notice burn artworks
    /// @param artworkIDs_ - the list of artwork id will be burned
    function burnArtworks(uint256[] memory artworkIDs_) public {
        require(isBurnable, "FeralfileExhibitionV4: not burnable");
        for (uint256 i = 0; i < artworkIDs_.length; i++) {
            require(
                _exists(artworkIDs_[i]),
                "FeralfileExhibitionV4: artwork not found"
            );
            require(
                _isApprovedOrOwner(_msgSender(), artworkIDs_[i]),
                "FeralfileExhibitionV4: caller is not owner nor approved"
            );
            delete artworks[artworkIDs_[i]];

            _burn(artworkIDs_[i]);

            emit BurnArtwork(artworkIDs_[i]);
        }
    }

    // @notice able to recieve funds from vault contract
    receive() external payable {
        require(msg.sender == vault, "only accept fund from vault contract.");
    }

    event NewArtwork(
        address indexed owner,
        uint256 indexed artworkIndex,
        uint256 indexed artworkID
    );
    event BuyArtwork(address indexed buyer, uint256 indexed tokenId);
    event BurnArtwork(uint256 indexed artworkID);
}
