pragma solidity ^0.8.13;

import "@openzeppelin/contracts/utils/Context.sol";

contract FFV4 {
  function ownerOf(uint256 tokenId) public view returns (address) {}
}

contract OwnerData is Context {
  struct Data {
    address owner; // Address of the owner who submitted the data
    bytes dataHash; // Hash of the actual data stored off-chain (e.g., IPFS CID)
    string metadata; // Any additional metadata associated with the data
  }

  // Mapping to store contract => tokenID => data[]
  mapping(address => mapping(uint256 => Data[])) private tokenData;

  constructor() {}

  function add(address contractAddress, uint256 tokenID, Data calldata data) external {
    // check ownership
    address tokenOwner = FFV4(contractAddress).ownerOf(tokenID);

    require(data.owner == _msgSender() && tokenOwner == _msgSender(), "OwnerData: owner mismatch");

    bool exists = false;
    
    for (uint i = 0; i < tokenData[contractAddress][tokenID].length; i++) {
      if (tokenData[contractAddress][tokenID][i].owner == data.owner) {
        // update existing data
        tokenData[contractAddress][tokenID][i] = data;
        exists = true;
        break;
      }
    }  

    if (!exists) {
      tokenData[contractAddress][tokenID].push(data);
    }
  
    emit DataAdded(contractAddress, tokenID, data);
  }

  function remove(address contractAddress, uint256 tokenID) external {
    // check ownership
    address tokenOwner = FFV4(contractAddress).ownerOf(tokenID);

    require(tokenOwner == _msgSender(), "OwnerData: owner mismatch");

    uint index;
    for (uint i = 0; i < tokenData[contractAddress][tokenID].length; i++) {
      if (tokenData[contractAddress][tokenID][i].owner == _msgSender()) {
        index = i;
        break;
      }
    }

    require(index >= 0 && index < tokenData[contractAddress][tokenID].length, "OwnerData: data not found");

    // remove data from array
    for (uint j = index; j < tokenData[contractAddress][tokenID].length - 1; j++) {
      tokenData[contractAddress][tokenID][j] = tokenData[contractAddress][tokenID][j + 1];
    }
    tokenData[contractAddress][tokenID].pop();
    
    // emit event
    emit DataRemoved(contractAddress, tokenID);
  }

  function get(address contractAddress, uint256 tokenID) external view returns (Data[] memory) {
    return tokenData[contractAddress][tokenID];
  }

  event DataAdded(address contractAddress, uint256 tokenID, Data data);
  event DataRemoved(address contractAddress, uint256 tokenID);
}