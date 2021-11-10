// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Authorizable is Ownable {
    address public trustee;

    constructor() {
        trustee = address(0x0);
    }

    modifier onlyAuthorized() {
        require(msg.sender == trustee || msg.sender == owner());
        _;
    }

    function setTrustee(address _newTrustee) public onlyOwner {
        trustee = _newTrustee;
    }
}
