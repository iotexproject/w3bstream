// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract W3bstreamProject is OwnableUpgradeable, ERC721Upgradeable {
    struct ProjectConfig {
        string uri;
        bytes32 hash;
    }

    event AttributeSet(uint256 indexed projectId, bytes32 indexed key, bytes value);
    event ProjectPaused(uint256 indexed projectId);
    event ProjectResumed(uint256 indexed projectId);
    event ProjectConfigUpdated(uint256 indexed projectId, string uri, bytes32 hash);
    event MinterSet(address indexed minter);

    mapping(uint256 => ProjectConfig) projectConfigs;
    mapping(uint256 => bool) paused;
    mapping(uint256 => mapping(bytes32 => bytes)) public attributes;

    address public minter;
    uint256 nextProjectId;

    modifier onlyProjectOwner(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender, "not owner");
        _;
    }

    function initialize(string calldata _name, string calldata _symbol) public initializer {
        __Ownable_init();
        __ERC721_init(_name, _symbol);
        setMinter(msg.sender);
    }

    function isPaused(uint256 _projectId) external view returns (bool) {
        _requireMinted(_projectId);
        return paused[_projectId];
    }

    function config(uint256 _projectId) external view returns (ProjectConfig memory) {
        _requireMinted(_projectId);
        return projectConfigs[_projectId];
    }

    function attribute(uint256 _projectId, bytes32 _name) external view returns (bytes memory) {
        return attributes[_projectId][_name];
    }

    function attributesOf(uint256 _projectId, bytes32[] memory _keys) external view returns (bytes[] memory values_) {
        _requireMinted(_projectId);

        values_ = new bytes[](_keys.length);
        mapping(bytes32 => bytes) storage attrs = attributes[_projectId];
        for (uint i = 0; i < _keys.length; i++) {
            values_[i] = attrs[_keys[i]];
        }
    }

    function mint(address _owner) external returns (uint256 projectId_) {
        require(msg.sender == minter, "not minter");

        projectId_ = ++nextProjectId;

        _mint(_owner, projectId_);
        paused[projectId_] = true;
    }

    function setAttributes(
        uint256 _projectId,
        bytes32[] memory _keys,
        bytes[] memory _values
    ) external onlyProjectOwner(_projectId) {
        require(_keys.length == _values.length, "invalid input");

        mapping(bytes32 => bytes) storage _attributes = attributes[_projectId];
        for (uint i = 0; i < _keys.length; i++) {
            _attributes[_keys[i]] = _values[i];
            emit AttributeSet(_projectId, _keys[i], _values[i]);
        }
    }

    function count() external view returns (uint256) {
        return nextProjectId;
    }

    function updateConfig(uint256 _projectId, string memory _uri, bytes32 _hash) external onlyProjectOwner(_projectId) {
        require(bytes(_uri).length != 0, "empty uri");
        ProjectConfig storage c = projectConfigs[_projectId];
        c.uri = _uri;
        c.hash = _hash;

        emit ProjectConfigUpdated(_projectId, _uri, _hash);
    }

    function pause(uint256 _projectId) external onlyProjectOwner(_projectId) {
        require(!paused[_projectId], "project already paused");
        paused[_projectId] = true;

        emit ProjectPaused(_projectId);
    }

    function resume(uint256 _projectId) external onlyProjectOwner(_projectId) {
        require(paused[_projectId], "project already actived");
        paused[_projectId] = false;

        emit ProjectResumed(_projectId);
    }

    function setMinter(address _minter) public onlyOwner {
        minter = _minter;

        emit MinterSet(_minter);
    }
}
