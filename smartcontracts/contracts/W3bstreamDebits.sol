// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract W3bstreamDebits is OwnableUpgradeable {
    event OperatorSet(uint256 indexed id, address indexed operator);
    event RewardTokenSet(uint256 indexed id, address indexed rewardToken);
    event Deposited(address indexed token, address indexed owner, uint256 amount);
    event Withheld(address indexed token, address indexed owner, uint256 amount);
    event Distributed(address indexed token, address indexed owner, address[] recipients, uint256[] amounts);
    event Redeemed(address indexed token, address indexed owner, uint256 amount);
    event Withdrawn(address indexed token, address indexed owner, uint256 amount);

    address public operator;

    mapping(uint256 => address) _rewardTokens;
    mapping(address => mapping(address => uint256)) _balances;
    mapping(address => mapping(address => uint256)) _withholdings;

    modifier onlyOperator() {
        require(msg.sender == operator, "not operator");
        _;
    } 

    function initialize() public initializer {
        __Ownable_init();
    }

    function rewardToken(uint256 _id) external view returns (address) {
        return _rewardTokens[_id];
    }

    function setRewardToken(uint256 _id, address _rewardToken) external onlyOwner {
        require(_rewardToken != address(0), "zero address");
        require(_rewardTokens[_id] == address(0), "already set");
        _rewardTokens[_id] = _rewardToken;
        emit RewardTokenSet(_id, _rewardToken);
    }

    function deposit(uint256 _id, uint256 _amount) external {
        address token = _rewardTokens[_id];
        require(token != address(0), "reward token not set");
        bool success = IERC20(token).transferFrom(msg.sender, address(this), _amount);
        require(success, "transfer failed");
        _balances[token][msg.sender] += _amount;
        emit Deposited(token, msg.sender, _amount);
    }

    function withhold(uint256 _id, address _owner, uint256 _amount) external {
        address token = _rewardTokens[_id];
        require(token != address(0), "reward token not set");
        require(_balances[token][_owner] - _withholdings[token][_owner] >= _amount, "insufficient balance");
        _withholdings[token][_owner] += _amount;
        _balances[token][_owner] -= _amount;
        emit Withheld(token, _owner, _amount);
    }

    function distribute(uint256 _id, address _owner, address[] calldata _recipients, uint256[] calldata _amounts) external onlyOperator {
        address token = _rewardTokens[_id];
        require(token != address(0), "reward token not set");
        require(_recipients.length == _amounts.length, "length mismatch");
        for (uint256 i = 0; i < _recipients.length; i++) {
            require(_withholdings[token][_owner] >= _amounts[i], "insufficient balance");
            _withholdings[token][_owner] -= _amounts[i];
            bool success = IERC20(token).transfer(_recipients[i], _amounts[i]);
            require(success, "transfer failed");
        }
        emit Distributed(token, _owner, _recipients, _amounts);
    }

    function redeem(uint256 _id, address _owner, uint256 _amount) external onlyOperator {
        address token = _rewardTokens[_id];
        require(token != address(0), "reward token not set");
        require(_withholdings[token][_owner] >= _amount, "insufficient balance");
        _withholdings[token][_owner] -= _amount;
        _balances[token][_owner] += _amount;
        emit Redeemed(token, _owner, _amount);
    }

    function withdraw(address _token, uint256 _amount) external {
        address sender = msg.sender;
        require(_token != address(0), "zero address");
        require(_balances[_token][sender] >= _amount, "insufficient balance");
        _balances[_token][sender] -= _amount;
        bool success = IERC20(_token).transfer(sender, _amount);
        require(success, "transfer failed");
        emit Withdrawn(_token, sender, _amount);
    }
}
