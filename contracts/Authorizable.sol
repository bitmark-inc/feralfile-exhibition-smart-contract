// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Authorizable is Ownable {
    mapping(address => bool) private trustees;

    constructor() {}

    modifier onlyAuthorized() {
        require(trustees[msg.sender] || msg.sender == owner());
        _;
    }

    function isTrustee(address _trustee) public view returns (bool) {
        return trustees[_trustee];
    }

    function addTrustee(address _trustee) public onlyOwner {
        trustees[_trustee] = true;
    }

    function removeTrustee(address _trustee) public onlyOwner {
        delete trustees[_trustee];
    }
}
