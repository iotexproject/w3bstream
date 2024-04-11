// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

contract W3bstreamCredit is OwnableUpgradeable, ERC20Upgradeable {
    address minter;
    event MinterSet(address minter);

    function initialize(string memory _name, string memory _symbol) public initializer {
        __Ownable_init();
        __ERC20_init(_name, _symbol);
    }

    function setMinter(address _minter) external onlyOwner {
        minter = _minter;
        emit MinterSet(_minter);
    }

    function grant(address _prover, uint256 _amount) external {
        require(msg.sender == minter, "not minter");
        _mint(_prover, _amount);
    }
}
