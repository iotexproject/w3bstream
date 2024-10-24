// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract W3bstreamDebits is OwnableUpgradeable {
    event OperatorSet(address indexed operator);
    event Deposited(address indexed token, address indexed owner, uint256 amount);
    event Withheld(address indexed token, address indexed owner, uint256 amount);
    event Distributed(address indexed token, address indexed owner, address[] recipients, uint256[] amounts);
    event Redeemed(address indexed token, address indexed owner, uint256 amount);
    event Withdrawn(address indexed token, address indexed owner, uint256 amount);

    address public operator;

    mapping(address => mapping(address => uint256)) balances;
    mapping(address => mapping(address => uint256)) withholdings;

    modifier onlyOperator() {
        require(msg.sender == operator, "not debits operator");
        _;
    }

    function initialize() public initializer {
        __Ownable_init();
    }

    function setOperator(address _operator) external onlyOwner {
        operator = _operator;
        emit OperatorSet(_operator);
    }

    function deposit(address token, uint256 amount) external {
        require(token != address(0), "invalid token");
        bool success = IERC20(token).transferFrom(msg.sender, address(this), amount);
        require(success, "transfer failed");
        balances[token][msg.sender] += amount;
        emit Deposited(token, msg.sender, amount);
    }

    function withhold(address token, address owner, uint256 amount) external onlyOperator {
        require(token != address(0), "reward token not set");
        require(balances[token][owner] >= amount, "insufficient balance");
        withholdings[token][owner] += amount;
        balances[token][owner] -= amount;
        emit Withheld(token, owner, amount);
    }

    function distribute(
        address token,
        address _owner,
        address[] calldata _recipients,
        uint256[] calldata _amounts
    ) external onlyOperator {
        require(token != address(0), "reward token not set");
        require(_recipients.length == _amounts.length, "length mismatch");
        for (uint256 i = 0; i < _recipients.length; i++) {
            withholdings[token][_owner] -= _amounts[i]; // overflow protected
            balances[token][_recipients[i]] += _amounts[i];
        }
        emit Distributed(token, _owner, _recipients, _amounts);
    }

    function redeem(address token, address _owner, uint256 _amount) external onlyOperator {
        require(token != address(0), "invalid token");
        require(withholdings[token][_owner] >= _amount, "insufficient balance");
        withholdings[token][_owner] -= _amount;
        balances[token][_owner] += _amount;
        emit Redeemed(token, _owner, _amount);
    }

    function withdraw(address _token, uint256 _amount) external {
        address sender = msg.sender;
        require(_token != address(0), "zero address");
        require(balances[_token][sender] >= _amount, "insufficient balance");
        balances[_token][sender] -= _amount;
        bool success = IERC20(_token).transfer(sender, _amount);
        require(success, "transfer failed");
        emit Withdrawn(_token, sender, _amount);
    }

    function balanceOf(address token, address owner) external view returns (uint256) {
        return balances[token][owner];
    }

    function withholdingOf(address token, address owner) external view returns (uint256) {
        return withholdings[token][owner];
    }
}
