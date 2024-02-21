// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract FF {
    function ownerOf(uint256 tokenId) public view returns (address) {}
}

string constant SIGNED_MESSAGE = "Feral File is requesting authorization to write your sound piece to contract";

contract OwnerData is Context {
    struct Data {
        address owner;
        bytes dataHash;
        string metadata;
    }

    mapping(address => mapping(uint256 => Data[])) private _tokenData;

    function add(address contractAddress, uint256 tokenID, Data calldata data) external {
        require(_isOwner(contractAddress, tokenID, _msgSender()), "OwnerData: caller is not the owner");
        require(data.owner == _msgSender(), "OwnerData: data owner mismatch");
        _updateData(contractAddress, tokenID, data);
    }

    function signedAdd(
        address contractAddress,
        uint256 tokenID,
        bytes memory signature,
        Data calldata data
    ) external {
        address signer = _verifySignature(signature);
        require(_isOwner(contractAddress, tokenID, signer), "OwnerData: signer is not the owner");
        require(data.owner == signer, "OwnerData: data owner mismatch");
        _updateData(contractAddress, tokenID, data);
    }

    function remove(address contractAddress, uint256 tokenID) external {
        require(_isOwner(contractAddress, tokenID, _msgSender()), "OwnerData: caller is not the owner");
        _removeData(contractAddress, tokenID);
    }

    function get(address contractAddress, uint256 tokenID) external view returns (Data[] memory) {
        return _tokenData[contractAddress][tokenID];
    }

    function _updateData(
        address contractAddress,
        uint256 tokenID,
        Data calldata data
    ) private {
        Data[] storage datas = _tokenData[contractAddress][tokenID];
        bool exists = false;
        for (uint256 i = 0; i < datas.length; ++i) {
            if (datas[i].owner == data.owner) {
                datas[i] = data;
                exists = true;
                break;
            }
        }
        if (!exists) {
            datas.push(data);
        }
        emit DataAdded(contractAddress, tokenID, data);
    }

    function _removeData(address contractAddress, uint256 tokenID) private {
        Data[] storage datas = _tokenData[contractAddress][tokenID];
        for (uint256 i = 0; i < datas.length; ++i) {
            if (datas[i].owner == _msgSender()) {
                datas[i] = datas[datas.length - 1];
                datas.pop();
                emit DataRemoved(contractAddress, tokenID);
                return;
            }
        }
        revert("OwnerData: data not found");
    }

    function _verifySignature(bytes memory signature) private view returns (address) {
        bytes memory message = abi.encodePacked(SIGNED_MESSAGE, " ", Strings.toHexString(address(this)), ".");
        return ECDSA.recover(ECDSA.toEthSignedMessageHash(message), signature);
    }

    function _isOwner(address contractAddress, uint256 tokenID, address owner) private view returns (bool) {
        return FF(contractAddress).ownerOf(tokenID) == owner;
    }

    event DataAdded(address indexed contractAddress, uint256 indexed tokenID, Data data);
    event DataRemoved(address indexed contractAddress, uint256 indexed tokenID);
}