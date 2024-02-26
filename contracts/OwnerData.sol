// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract IFF {
    function ownerOf(uint256 tokenId) public view returns (address) {}
}

string constant SIGNED_MESSAGE = "Authorize to write your data to the contract";

contract OwnerData is Context {
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

    address private _trustee;

    // contractAddress => tokenID => Data[]
    mapping(address => mapping(uint256 => Data[])) private _tokenData;
    // contractAddress => tokenID => owner => bool
    mapping(address => mapping(uint256 => mapping(address => bool))) private _tokenDataOwner;

    constructor(address trustee_) {
        require(trustee_ != address(0), "OwnerData: Trustee is the zero address");
        _trustee = trustee_;
    }

    function add(address contractAddress, uint256 tokenID, Data calldata data) external {
        _addData(_msgSender(), contractAddress, tokenID, data);
    }

    function signedAdd(
        address contractAddress,
        uint256 tokenID,
        Data calldata data,
        Signature calldata signature
    ) external {
        _validateSignature(signature);
        address account = _recoverOwnerSignature(signature.ownerSign);
        _addData(account, contractAddress, tokenID, data);
    }


    function get(address contractAddress, uint256 tokenID) external view returns (Data[] memory) {
        return _tokenData[contractAddress][tokenID];
    }

    function _addData(
        address sender,
        address contractAddress,
        uint256 tokenID,
        Data calldata data
    ) private {
        require(_isOwner(contractAddress, tokenID, sender), "OwnerData: sender is not the owner");
        require(data.owner == sender, "OwnerData: data owner mismatch");
        require(data.dataHash.length > 0, "OwnerData: dataHash is empty");
        require(!_tokenDataOwner[contractAddress][tokenID][data.owner], "OwnerData: data already added");

        _tokenData[contractAddress][tokenID].push(data);
        _tokenDataOwner[contractAddress][tokenID][data.owner] = true;

        emit DataAdded(contractAddress, tokenID, data);
    }

    function _validateSignature(Signature calldata signature) private view {
        require(signature.expiryBlock >= block.number, "OwnerData: signature expired");

        bytes32 message = keccak256(
            abi.encode(block.chainid, address(this), signature.ownerSign, signature.expiryBlock)
        );
        address reqSigner = ECDSA.recover(
            ECDSA.toEthSignedMessageHash(message),
            signature.v,
            signature.r,
            signature.s
        );
        require(reqSigner == _trustee, "OwnerData: invalid signature");
    }

    function _recoverOwnerSignature(bytes memory signature) private view returns (address) {
        bytes memory message = abi.encodePacked(SIGNED_MESSAGE, " ", Strings.toHexString(address(this)), ".");
        return ECDSA.recover(ECDSA.toEthSignedMessageHash(message), signature);
    }

    function _isOwner(address contractAddress, uint256 tokenID, address account) private view returns (bool) {
        return IFF(contractAddress).ownerOf(tokenID) == account;
    }

    event DataAdded(address indexed contractAddress, uint256 indexed tokenID, Data data);
}