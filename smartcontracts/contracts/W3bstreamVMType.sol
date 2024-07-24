// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract W3bstreamVMType is OwnableUpgradeable, ERC721Upgradeable {
    event VMTypeSet(uint256 indexed id);
    event VMTypePaused(uint256 indexed id);
    event VMTypeResumed(uint256 indexed id);

    uint256 nextTypeId;

    mapping(uint256 => string) _types;
    mapping(uint256 => bool) _paused;

    modifier onlyVMTypeOwner(uint256 _id) {
        require(ownerOf(_id) == msg.sender, "not owner");
        _;
    }

    function initialize(string memory _name, string memory _symbol) public initializer {
        __Ownable_init();
        __ERC721_init(_name, _symbol);
    }

    function count() external view returns (uint256) {
        return nextTypeId;
    }

    function vmTypeName(uint256 _id) external view returns (string memory) {
        _requireMinted(_id);
        return _types[_id];
    }

    function isPaused(uint256 _id) external view returns (bool) {
        _requireMinted(_id);
        return _paused[_id];
    }

    function mint(string memory _name) external returns (uint256 id_) {
        id_ = ++nextTypeId;
        _mint(msg.sender, id_);

        _types[id_] = _name;
        _paused[id_] = false;
        emit VMTypeSet(id_);
    }

    function pause(uint256 _id) external onlyVMTypeOwner(_id) {
        require(!_paused[_id], "already paused");

        _paused[_id] = true;
        emit VMTypePaused(_id);
    }

    function resume(uint256 _id) external onlyVMTypeOwner(_id) {
        require(_paused[_id], "already actived");

        _paused[_id] = false;
        emit VMTypeResumed(_id);
    }
}
