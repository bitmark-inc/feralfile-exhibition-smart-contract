// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract FFV5 {
    function balanceOf(address account, uint256 id) public view returns (uint256) {}
}

string constant SIGNED_MESSAGE = "Authorize to write your data to the contract";

contract OwnerData is Context {
    struct Data {
        address owner;
        bytes dataHash;
        string metadata;
    }

    // contractAddress => tokenID => Data[]
    mapping(address => mapping(uint256 => Data[])) private _tokenData;
    // contractAddress => tokenID => owner => bool
    mapping(address => mapping(uint256 => mapping(address => bool))) private _tokenDataOwner;

    function add(address contractAddress, uint256 tokenID, Data calldata data) external {
        _addData(_msgSender(), contractAddress, tokenID, data);
    }

    function signedAdd(
        address contractAddress,
        uint256 tokenID,
        bytes memory signature,
        Data calldata data
    ) external {
        address signer = _verifySignature(signature);
        _addData(signer, contractAddress, tokenID, data);
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
        if (!_isOwner(contractAddress, tokenID, sender)) {
            revert NotOwner(sender, contractAddress, tokenID);
        }
        if (data.owner != sender) {
            revert OwnerMismatch(data.owner, sender);
        }
        require(data.dataHash.length > 0, "OwnerData: dataHash is empty");
        require(!_tokenDataOwner[contractAddress][tokenID][data.owner], "OwnerData: data already added");

        _tokenData[contractAddress][tokenID].push(data);
        _tokenDataOwner[contractAddress][tokenID][data.owner] = true;

        emit DataAdded(contractAddress, tokenID, data);
    }

    function _verifySignature(bytes memory signature) private view returns (address) {
        bytes memory message = abi.encodePacked(SIGNED_MESSAGE, " ", Strings.toHexString(address(this)), ".");
        return ECDSA.recover(ECDSA.toEthSignedMessageHash(message), signature);
    }

    function _isOwner(address contractAddress, uint256 tokenID, address account) private view returns (bool) {
        return FFV5(contractAddress).balanceOf(account, tokenID) > 0;
    }

    event DataAdded(address indexed contractAddress, uint256 indexed tokenID, Data data);

    error NotOwner(address caller, address contractAddress, uint256 tokenID); 
    error OwnerMismatch(address dataOwner, address caller);
}