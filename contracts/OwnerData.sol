// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

/**
 * @title OwnerData
 * @dev Manages addition data for ERC721 or ERC1155 token.
 */
contract OwnerData is Context, Ownable {
    string private constant SIGNED_MESSAGE = "Authorize to write your data to the contract";
    address private immutable _signer;
    address private immutable _costReceiver;
    uint256 private immutable _cost;

    struct Data {
        address owner;
        bytes dataHash;
        string metadata;
    }

    struct Signature {
        bytes ownerSign;
        uint256 expiryBlock;
        bytes32 r;
        bytes32 s;
        uint8 v;
    }

    struct SignedAddParams {
        address contractAddress;
        uint256 tokenID;
        Data data;
        Signature signature;
    }

    // contractAddress => tokenID => Data[]
    mapping(address => mapping(uint256 => Data[])) private _tokenData;
    // contractAddress => tokenID => owner => bool
    mapping(address => mapping(uint256 => mapping(address => bool))) private _tokenDataOwner;
    // contractAddress => tokenID => bool
    mapping(address => mapping(uint256 => bool)) private _publicTokens;

    event DataAdded(address indexed contractAddress, uint256 indexed tokenID, Data data);

    constructor(address signer_, address costReceiver_, uint256 cost_) {
        require(signer_ != address(0), "OwnerData: Trustee is the zero address");
        require(costReceiver_ != address(0), "OwnerData: Cost receiver is the zero address");
        require(cost_ > 0, "OwnerData: Cost is zero");
        _signer = signer_;
        _costReceiver = costReceiver_;
        _cost = cost_;
    }

    function add(address contractAddress_, uint256 tokenID_, Data calldata data_) external payable {
        require(!_publicTokens[contractAddress_][tokenID_] || msg.value == _cost, "OwnerData: Payment required for public token");
        _addData(_msgSender(), contractAddress_, tokenID_, data_);
        if (msg.value > 0) {
            payable(_costReceiver).transfer(msg.value);
        }
    }

    function get(address contractAddress_, uint256 tokenID_, uint256 startIndex, uint256 count) public view returns (Data[] memory) {
        Data[] memory data = _tokenData[contractAddress_][tokenID_];
        require(startIndex >= 0 && count > 0 && startIndex < data.length, "OwnerData: Invalid parameters");
        if (count > data.length - startIndex) {
            count = data.length - startIndex;
        }
        Data[] memory result = new Data[](count);
        for (uint256 i = 0; i < count; i++) {
            result[i] = data[startIndex + i];
        }
        return result;
    }

    function remove(address contractAddress_, uint256 tokenID_, uint256[] calldata indexes_) external {
        Data[] storage data = _tokenData[contractAddress_][tokenID_];
        for (uint256 i = 0; i < indexes_.length; i++) {
            require(indexes_[i] < data.length, "Index out of bounds");
            if (indexes_[i] != data.length - 1) {
                data[indexes_[i]] = data[data.length - 1];
            }
            data.pop();
        }
    }

    function setPublicTokens(address[] memory contractAddresses_, uint256[] memory tokenIDs_, bool isPublic_) external onlyOwner {
        require(contractAddresses_.length == tokenIDs_.length, "OwnerData: Arrays length mismatch");
        for (uint256 i = 0; i < contractAddresses_.length; i++) {
            _publicTokens[contractAddresses_[i]][tokenIDs_[i]] = isPublic_;
        }
    }

    function signedAdd(SignedAddParams[] calldata params_) external {
        for (uint256 i = 0; i < params_.length; i++) {
            _signedAdd(params_[i]);
        }
    }


    function _signedAdd(SignedAddParams calldata params_) private {
        _validateSignature(params_.signature);
        if (_publicTokens[params_.contractAddress][params_.tokenID]) {
            _addData(params_.data.owner, params_.contractAddress, params_.tokenID, params_.data);
        } else {
            address account = _recoverOwnerSignature(params_.signature.ownerSign);
            _addData(account, params_.contractAddress, params_.tokenID, params_.data);
        }
    }

    function _addData(
        address sender_,
        address contractAddress_,
        uint256 tokenID_,
        Data calldata data_
    ) private {
        require(data_.owner == sender_, "OwnerData: data owner and sender mismatch");
        require(data_.dataHash.length > 0, "OwnerData: dataHash is empty");

        if (!_publicTokens[contractAddress_][tokenID_]) {
            require(_isOwner(contractAddress_, tokenID_, data_.owner), "OwnerData: sender is not the owner");
            require(!_tokenDataOwner[contractAddress_][tokenID_][data_.owner], "OwnerData: data already added");
            _tokenDataOwner[contractAddress_][tokenID_][data_.owner] = true;
        }

        _tokenData[contractAddress_][tokenID_].push(data_);

        emit DataAdded(contractAddress_, tokenID_, data_);
    }

    function _validateSignature(Signature calldata signature_) private view {
        require(block.number < signature_.expiryBlock, "OwnerData: signature expired");
        bytes32 message = keccak256(
            abi.encode(block.chainid, address(this), signature_.ownerSign, signature_.expiryBlock)
        );
        address reqSigner = ECDSA.recover(
            ECDSA.toEthSignedMessageHash(message),
            signature_.v,
            signature_.r,
            signature_.s
        );
        require(reqSigner == _signer, "OwnerData: Invalid signature");
    }

    function _recoverOwnerSignature(bytes memory signature_) private view returns (address) {
        bytes memory message = abi.encodePacked(SIGNED_MESSAGE, " ", Strings.toHexString(address(this)), ".");
        return ECDSA.recover(ECDSA.toEthSignedMessageHash(message), signature_);
    }

    function _isOwner(address contractAddress, uint256 tokenID, address account) private view returns (bool) {
        if (IERC165(contractAddress).supportsInterface(type(IERC1155).interfaceId)) {
            return IERC1155(contractAddress).balanceOf(account, tokenID) > 0;
        }
        if (IERC165(contractAddress).supportsInterface(type(IERC721).interfaceId)) {
            return IERC721(contractAddress).ownerOf(tokenID) == account;
        }
        return false;
    }
}