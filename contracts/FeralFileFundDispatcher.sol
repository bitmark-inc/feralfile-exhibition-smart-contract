// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract FeralFileFundDispatcher is Ownable {
    using SafeMath for uint256;

    address public operator; // Address operator

    // Modifiers
    modifier onlyOperator() {
        require(_msgSender() == operator, "Only operator can call this function.");
        _;
    }

    constructor() {
        operator = msg.sender;
    }

    /// @notice update the operator address
    /// @param _newOperator - new operator address
    function updateOperator(address _newOperator) public onlyOwner {
        require(_newOperator != address(0), "Invalid operator address");

        operator = _newOperator;

        emit NewOperator(_newOperator);
    }

    /// @notice To withdraw tokens from contract, to deposit directly transfer to the contract
    /// @param _tokenAddress - erc20 token address
    /// @param _value - token amount to withdraw
    function withdrawToken(address _tokenAddress, uint256 _value) public onlyOperator {
        // Check if contract is having required balance
        IERC20 token = IERC20(_tokenAddress);
        require(
            token.balanceOf(address(this)) >= _value,
            "insufficient balance"
        );
        require(
            token.transfer(msg.sender, _value),
            "Unable to transfer token to the owner account"
        );

        emit WithdrawToken(msg.sender, _value);
    }

    /// @notice To withdraw eth from contract, to deposit directly transfer to the contract
    /// @param _value - token amount to withdraw
    function withdrawETH(uint256 _value) public onlyOperator {
        require(
            address(this).balance >= _value,
            "insufficient balance"
        );
        payable(msg.sender).transfer(_value);

        emit WithdrawETH(msg.sender, _value);
    }

    /// @notice To transfer tokens from this contract to the provided list of token holders with respective amount
    /// @param _tokenAddress - erc20 token address
    /// @param _receivers - receiver addresses
    /// @param _amounts - receiver amounts
    function batchTransferToken(
        address _tokenAddress,
        address[] calldata _receivers,
        uint256[] calldata _amounts
    ) external onlyOperator {
        require(
            _receivers.length == _amounts.length,
            "Invalid input parameters"
        );

        IERC20 token = IERC20(_tokenAddress);
        uint256 total = token.balanceOf(address(this));

        for (uint256 i = 0; i < _receivers.length; i++) {
            require(total >= _amounts[i], "insufficient balance");

            require(token.transfer(_receivers[i], _amounts[i]),
                "Unable to transfer token to the account"
            );

            total = total - _amounts[i];
        }
    }

    /// @notice To transfer tokens from tokenOwner to the provided list of token holders with respective amount
    /// @param _tokenAddress - erc20 token address
    /// @param _tokenOwner - token owner address
    /// @param _receivers - receiver addresses
    /// @param _amounts - receiver amounts
    function batchTransferTokenFrom(
        address _tokenAddress,
        address _tokenOwner,
        address[] calldata _receivers,
        uint256[] calldata _amounts
    ) external onlyOperator {
        require(
            _receivers.length == _amounts.length,
            "Invalid input parameters"
        );

        IERC20 token = IERC20(_tokenAddress);
        uint256 total = token.balanceOf(_tokenOwner);

        for (uint256 i = 0; i < _receivers.length; i++) {
            require(total >= _amounts[i], "insufficient balance");

            require(token.transferFrom(_tokenOwner, _receivers[i], _amounts[i]),
                "Unable to transfer token to the account"
            );

            total = total - _amounts[i];
        }
    }

    /// @notice To transfer ethereum to the provided list of token holders with respective amount
    /// @param _receivers - receiver addresses
    /// @param _amounts - receiver amounts
    function batchTransferETH(
        address[] calldata _receivers,
        uint256[] calldata _amounts
    ) external onlyOperator payable {
        require(
            _receivers.length == _amounts.length,
            "Invalid input parameters"
        );

        uint256 total = msg.value;

        for (uint256 i = 0; i < _receivers.length; i++) {
            require(total >= _amounts[i], "insufficient balance");

            payable(_receivers[i]).transfer(_amounts[i]);

            total = total - _amounts[i];
        }
    }

    // Events
    event NewOperator(address operator);
    event WithdrawToken(address indexed receiver, uint256 amount);
    event WithdrawETH(address indexed receiver, uint256 amount);
}
