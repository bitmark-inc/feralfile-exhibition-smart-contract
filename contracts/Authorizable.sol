// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Authorizable is Ownable {
    address[] public trustees;

    constructor() {}

    modifier onlyAuthorized() {
        bool isAuthorized = false;
        for (uint256 i = 0; i < trustees.length; i++) {
            if (msg.sender == trustees[i]) {
                isAuthorized = true;
                break;
            }
        }
        require(isAuthorized || msg.sender == owner());
        _;
    }

    function setTrustees(address[] memory _newTrustees) public onlyOwner {
        trustees = _newTrustees;
    }
}
