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
    address public signer;
    address public serviceFeeReceiver;
    uint256 public serviceFee;

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

    // contractAddress => tokenID => Data[]
    mapping(address => mapping(uint256 => Data[])) private _tokenData;
    // contractAddress => tokenID => owner => bool
    mapping(address => mapping(uint256 => mapping(address => bool))) private _tokenDataOwner;
    // contractAddress => tokenID => bool
    mapping(address => mapping(uint256 => bool)) private _publicTokens;

    event DataAdded(address indexed contractAddress, uint256 indexed tokenID, Data data);
    event DataRemoved(address indexed contractAddress, uint256 indexed tokenID, uint256[] indexes);

    error TrusteeIsZeroAddress();
    error EmptyServiceFeeReceiver();
    error PaymentRequiredForPublicToken();
    error InvalidParameters();
    error OwnerAndSenderMismatch();
    error SenderIsNotTheOwner();
    error OwnerDataAlreadyAdded();
    error InvalidSignature();


    constructor(address signer_, address serviceFeeReceiver_, uint256 serviceFee_) {
        if (signer_ == address(0)) {
            revert TrusteeIsZeroAddress();
        }
        if (serviceFeeReceiver_ == address(0)) {
            revert EmptyServiceFeeReceiver();
        }
        signer = signer_;
        serviceFeeReceiver = serviceFeeReceiver_;
        serviceFee = serviceFee_;
    }

    /// @notice add data to the token
    /// @param contractAddress_ - the address of the contract
    /// @param tokenID_ - the token ID
    /// @param data_ - the data to add
    function add(address contractAddress_, uint256 tokenID_, Data calldata data_) external payable {
        if (_publicTokens[contractAddress_][tokenID_] && msg.value < serviceFee) {
            revert PaymentRequiredForPublicToken();
        }
        _addData(_msgSender(), contractAddress_, tokenID_, data_);
        if (msg.value > 0) {
            payable(serviceFeeReceiver).transfer(msg.value);
        }
    }

    /// @notice get data by contract address and token ID
    /// @param contractAddress_ - the address of the contract
    /// @param tokenID_ - the token ID
    /// @param startIndex - the start index
    /// @param count - the count of data
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

    /// @notice get data by contract address, token ID and owner
    /// @param contractAddress_ - the address of the contract
    /// @param tokenID_ - the token ID
    /// @param owner_ - the owner address
    function getByOwner(address contractAddress_, uint256 tokenID_, address owner_) public view returns (Data[] memory) {
        Data[] memory data = _tokenData[contractAddress_][tokenID_];
        Data[] memory temp = new Data[](data.length);
        uint256 count = 0;
        for (uint256 i = 0; i < data.length; i++) {
            if (data[i].owner == owner_) {
                temp[count] = data[i];
                count++;
            }
        }
        Data[] memory result = new Data[](count);
        for (uint256 i = 0; i < count; i++) {
            result[i] = temp[i];
        }
        return result;
    }

    /// @notice remove data by indexes
    /// @param contractAddress_ - the address of the contract
    /// @param tokenID_ - the token ID
    /// @param indexes_ - the indexes of the data to remove
    function remove(address contractAddress_, uint256 tokenID_, uint256[] calldata indexes_) external onlyOwner {
        Data[] storage data = _tokenData[contractAddress_][tokenID_];
        for (uint256 i = 0; i < indexes_.length; i++) {
            data[indexes_[i]].dataHash = new bytes(0);
        }
        emit DataRemoved(contractAddress_, tokenID_, indexes_);
    }

    /// @notice the service fee of adding data
    /// @param serviceFee_ - the service fee of adding data
    function setServiceFee(uint256 serviceFee_) external onlyOwner {
        serviceFee = serviceFee_;
    }

    /// @notice the service fee receiver address
    /// @param serviceFeeReceiver_ - the address of service fee receiver
    function setServiceFeeReceiver(address serviceFeeReceiver_) external onlyOwner {
        if (serviceFeeReceiver_ == address(0)) {
            revert EmptyServiceFeeReceiver();
        }
        serviceFeeReceiver = serviceFeeReceiver_;
    }

    /// @notice the address of the signer
    /// @param signer_ - the address of the signer
    function setSigner(address signer_) external onlyOwner {
        if (signer_ == address(0)) {
            revert TrusteeIsZeroAddress();
        }
        signer = signer_;
    }

    /// @notice set public tokens
    /// @param contractAddresses_ - the addresses of the contracts
    /// @param tokenIDs_ - the token IDs
    /// @param isPublic_ - the flag of public token
    function setPublicTokens(address[] memory contractAddresses_, uint256[] memory tokenIDs_, bool isPublic_) external onlyOwner {
        if (contractAddresses_.length == 0 || contractAddresses_.length != tokenIDs_.length) {
            revert InvalidParameters();
        }
        for (uint256 i = 0; i < contractAddresses_.length; i++) {
            _publicTokens[contractAddresses_[i]][tokenIDs_[i]] = isPublic_;
        }
    }

    /// @notice add data with signature
    /// @param contractAddress_ - the address of the contract
    /// @param tokenID_ - the token ID
    /// @param data_ - the data to add
    /// @param signature_ - the signature
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
        if (reqSigner != signer) {
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