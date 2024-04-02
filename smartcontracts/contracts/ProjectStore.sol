// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProjectStore.sol";

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract ProjectStore is IProjectStore, OwnableUpgradeable, ERC721Upgradeable {
    event MinterChanged(address indexed minter);

    mapping(uint256 => Project) projects;
    mapping(uint256 => bool) paused;
    mapping(uint256 => mapping(bytes32 => bytes)) public override attributes;

    address public minter;
    uint256 nextProjectId;

    modifier onlyProjectOwner(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender, "not owner");
        _;
    }

    function initialize(address _minter, string calldata _name, string calldata _symbol) public initializer {
        __Ownable_init();
        __ERC721_init(_name, _symbol);
        minter = _minter;
        emit MinterChanged(_minter);
    }

    function isPaused(uint256 _projectId) external view override returns (bool) {
        _requireMinted(_projectId);
        return paused[_projectId];
    }

    function project(uint256 _projectId) external view override returns (Project memory) {
        _requireMinted(_projectId);
        Project memory project = projects[_projectId];
        return (project.uri, project.hash);
    }

    function attribute(uint256 _projectId, bytes32 _name) external view returns (bytes calldata) {
        return attributes[_projectId][_name];
    }

    function attributesOf(
        uint256 _projectId,
        bytes32[] memory _keys
    ) external view override returns (bytes[] memory values_) {
        _requireMinted(_projectId);

        values_ = new bytes[](_keys.length);
        mapping(bytes32 => bytes) storage project = attributes[_projectId];
        for (uint i = 0; i < _keys.length; i++) {
            values_[i] = project[_keys[i]];
        }
    }

    function mint(address _owner, string calldata _uri, bytes32 _hash) external override returns (uint256 _projectId) {
        require(msg.sender == minter, "not minter");

        projectId = ++nextProjectId;

        _mint(_owner, projectId);
        paused[projectId] = true;
        updateProjectInternal(projectId, _uri, _hash);
    }

    function setAttributes(
        uint64 _projectId,
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
        return nextProjectId + 1;
    }
    function updateProjectInternal(
        uint256 _projectId,
        string memory _uri,
        bytes32 _hash
    ) internal {
        require(bytes(_uri).length != 0, "empty uri");
        Project storage project = projects[_projectId];
        project.uri = _uri;
        project.hash = _hash;

        emit ProjectUpdated(_projectId, _uri, _hash);
    }

    function updateProject(
        uint256 _projectId,
        string memory _uri,
        bytes32 _hash
    ) external override onlyProjectOwner(_projectId) {
        updateProjectInternal(_projectId, _uri, _hash);
    }

    function pause(uint256 _projectId) external override onlyProjectOwner(_projectId) {
        require(!paused[_projectId], "project already paused");
        paused[_projectId] = true;

        emit ProjectPaused(_projectId);
    }

    function resume(uint256 _projectId) external override onlyProjectOwner(_projectId) {
        require(paused[_projectId], "project already actived");
        paused[_projectId] = false;

        emit ProjectResumed(_projectId);
    }

    function changeMinter(address _minter) external onlyOwner {
        minter = _minter;

        emit MinterChanged(_minter);
    }
}
