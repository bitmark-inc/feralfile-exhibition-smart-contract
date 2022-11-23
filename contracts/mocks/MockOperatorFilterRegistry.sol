// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MockOperatorFilterRegistry {
    address private allowedExchange;

    constructor(address allowedExchange_) {
        allowedExchange = allowedExchange_;
    }

    function isOperatorAllowed(address registrant, address operator)
        external
        view
        returns (bool)
    {
        // allow for all regular address
        if (operator.code.length == 0) {
            return true;
        }

        if (operator == allowedExchange) {
            return true;
        }

        return false;
    }
}
