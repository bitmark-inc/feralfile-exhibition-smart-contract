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
    uint256 public cost;

    struct Data {
        address owner;
        bytes dataHash;
        uint256 blockNumber;
        string metadata;
    }

    struct Signature {
        bytes ownerSign;
        uint256 expiryBlock;
        bytes32 r;
        bytes32 s;
        uint8 v;
    }

    // contractAddress => tokenID => Data[]
    mapping(address => mapping(uint256 => Data[])) private _tokenData;
    // contractAddress => tokenID => owner => bool
    mapping(address => mapping(uint256 => mapping(address => bool))) private _tokenDataOwner;
    // contractAddress => tokenID => bool
    mapping(address => mapping(uint256 => bool)) private _publicTokens;

    event DataAdded(address indexed contractAddress, uint256 indexed tokenID, Data data);

    error TrusteeIsZeroAddress();
    error CostReceiverIsZeroAddress();
    error CostIsZero();
    error PaymentRequiredForPublicToken();
    error IndexOutOfBounds();
    error InvalidParameters();
    error OwnerAndSenderMismatch();
    error SenderIsNotTheOwner();
    error OwnerDataAlreadyAdded();
    error InvalidSignature();


    constructor(address signer_, address costReceiver_, uint256 cost_) {
        if (signer_ == address(0)) {
            revert TrusteeIsZeroAddress();
        }
        if (costReceiver_ == address(0)) {
            revert CostReceiverIsZeroAddress();
        }
        if (cost_ == 0) {
            revert CostIsZero();
        }
        _signer = signer_;
        _costReceiver = costReceiver_;
        cost = cost_;
    }

    function add(address contractAddress_, uint256 tokenID_, Data calldata data_) external payable {
        if (_publicTokens[contractAddress_][tokenID_] &&  msg.value != cost) {
            revert PaymentRequiredForPublicToken();
        }
        _addData(_msgSender(), contractAddress_, tokenID_, data_);
        if (msg.value > 0) {
            payable(_costReceiver).transfer(msg.value);
        }
    }

    function get(address contractAddress_, uint256 tokenID_, uint256 startIndex, uint256 count) public view returns (Data[] memory) {
        Data[] memory data = _tokenData[contractAddress_][tokenID_];
        if (startIndex > data.length) {
            return new Data[](0);
        }
        if (startIndex + count > data.length) {
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
            if (indexes_[i] >= data.length) {
                revert IndexOutOfBounds();
            }
            if (indexes_[i] != data.length - 1) {
                data[indexes_[i]] = data[data.length - 1];
            }
            data.pop();
        }
    }

    function setCost(uint256 cost_) external onlyOwner {
        if (cost_ == 0) {
            revert CostIsZero();
        }
        cost = cost_;
    }

    function setPublicTokens(address[] memory contractAddresses_, uint256[] memory tokenIDs_, bool isPublic_) external onlyOwner {
        if (contractAddresses_.length == 0 || contractAddresses_.length != tokenIDs_.length) {
            revert InvalidParameters();
        }
        for (uint256 i = 0; i < contractAddresses_.length; i++) {
            _publicTokens[contractAddresses_[i]][tokenIDs_[i]] = isPublic_;
        }
    }

    function signedAdd(address contractAddress_, uint256 tokenID_, Data calldata data_, Signature calldata signature_) external {
        _validateSignature(signature_);
        address account = data_.owner;
        if (!_publicTokens[contractAddress_][tokenID_]) {
            account = _recoverOwnerSignature(signature_.ownerSign);
        }
        _addData(account, contractAddress_, tokenID_, data_);
    }

    function _addData(address sender_, address contractAddress_, uint256 tokenID_, Data memory data_) private {
        if (data_.owner != sender_) {
            revert OwnerAndSenderMismatch();
        }
        if (data_.dataHash.length == 0) {
            revert InvalidParameters();
        }

        if (!_publicTokens[contractAddress_][tokenID_]) {
            if (_tokenDataOwner[contractAddress_][tokenID_][data_.owner]) {
                revert OwnerDataAlreadyAdded();
            }
            if (!_isOwner(contractAddress_, tokenID_, data_.owner)) {
                revert SenderIsNotTheOwner();
            }
            _tokenDataOwner[contractAddress_][tokenID_][data_.owner] = true;
        }

        data_.blockNumber = block.number;

        _tokenData[contractAddress_][tokenID_].push(data_);

        emit DataAdded(contractAddress_, tokenID_, data_);
    }

    function _validateSignature(Signature calldata signature_) private view {
        if (block.number > signature_.expiryBlock) {
            revert InvalidSignature();
        }
        bytes32 message = keccak256(
            abi.encode(block.chainid, address(this), signature_.ownerSign, signature_.expiryBlock)
        );
        address reqSigner = ECDSA.recover(
            ECDSA.toEthSignedMessageHash(message),
            signature_.v,
            signature_.r,
            signature_.s
        );
        if (reqSigner != _signer) {
            revert InvalidSignature();
        }
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