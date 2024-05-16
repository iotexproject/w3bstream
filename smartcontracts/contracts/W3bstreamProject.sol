// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract W3bstreamProject is OwnableUpgradeable {
    struct ProjectConfig {
        string uri;
        bytes32 hash;
    }

    event ProjectBinded(uint256 indexed projectId);
    event AttributeSet(uint256 indexed projectId, bytes32 indexed key, bytes value);
    event ProjectPaused(uint256 indexed projectId);
    event ProjectResumed(uint256 indexed projectId);
    event ProjectConfigUpdated(uint256 indexed projectId, string uri, bytes32 hash);
    event BinderSet(address indexed binder);

    mapping(uint256 => bool) projects;
    mapping(uint256 => ProjectConfig) projectConfigs;
    mapping(uint256 => bool) paused;
    mapping(uint256 => mapping(bytes32 => bytes)) public attributes;

    IERC721 public project;
    address public binder;
    uint256 public count;

    modifier onlyProjectOwner(uint256 _projectId) {
        require(project.ownerOf(_projectId) == msg.sender, "not owner");
        _;
    }

    function requireProjectRegister(uint256 _projectId) internal view virtual {
        require(project.ownerOf(_projectId) != address(0), "invalid project");
    }

    function ownerOf(uint256 _projectId) external view returns (address) {
        return project.ownerOf(_projectId);
    }

    function initialize(address _project) public initializer {
        __Ownable_init();
        project = IERC721(_project);
        setBinder(msg.sender);
    }

    function isPaused(uint256 _projectId) external view returns (bool) {
        requireProjectRegister(_projectId);
        return paused[_projectId];
    }

    function config(uint256 _projectId) external view returns (ProjectConfig memory) {
        requireProjectRegister(_projectId);
        return projectConfigs[_projectId];
    }

    function attribute(uint256 _projectId, bytes32 _name) external view returns (bytes memory) {
        return attributes[_projectId][_name];
    }

    function attributesOf(uint256 _projectId, bytes32[] memory _keys) external view returns (bytes[] memory values_) {
        requireProjectRegister(_projectId);

        values_ = new bytes[](_keys.length);
        mapping(bytes32 => bytes) storage attrs = attributes[_projectId];
        for (uint i = 0; i < _keys.length; i++) {
            values_[i] = attrs[_keys[i]];
        }
    }

    function bind(uint256 _projectId) external {
        require(msg.sender == binder, "not binder");
        require(!projects[_projectId], "already bind");

        count++;
        paused[_projectId] = true;
        projects[_projectId] = true;
        emit ProjectBinded(_projectId);
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

    function setBinder(address _binder) public onlyOwner {
        binder = _binder;

        emit BinderSet(_binder);
    }
}
