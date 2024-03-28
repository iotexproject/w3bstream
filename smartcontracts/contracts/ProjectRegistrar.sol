// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProjectRegistrar.sol";

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract ProjectRegistrar is IProjectRegistrar, OwnableUpgradeable, ERC721Upgradeable {
    event FeeWithdrawn(address indexed account, uint256 amount);

    uint256 public override fee;

    mapping(uint256 => Project) projects;
    uint256 nextProjectId;

    modifier onlyProjectOwner(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender, "Only project owner to operate");
        _;
    }

    function _requireOperator(address _operator) internal view virtual {
        require(_operator == msg.sender, "Only project operator to operate");
    }

    function initialize(uint256 _fee, string calldata _name, string calldata _symbol) public initializer {
        __Ownable_init();
        __ERC721_init(_name, _symbol);
        fee = _fee;
    }

    function isPaused(uint256 _projectId) external view override returns (bool) {
        return projects[_projectId].paused;
    }

    function operator(uint256 _projectId) external view override returns (address) {
        return projects[_projectId].operator;
    }

    function hash(uint256 _projectId) external view override returns (bytes32) {
        return projects[_projectId].hash;
    }

    function uri(uint256 _projectId) external view override returns (string memory) {
        return projects[_projectId].uri;
    }

    function metadata(uint256 _projectId, bytes32 _name) external view override returns (bytes memory) {
        return projects[_projectId].metadata[_name];
    }

    function register(string calldata _uri, bytes32 _hash) external payable override returns (uint256 _projectId) {
        require(msg.value > fee, "Register fee too low");
        _projectId = ++nextProjectId;

        Project storage _project = projects[_projectId];
        _project.paused = false;
        _project.operator = msg.sender;
        _project.hash = _hash;
        _project.uri = _uri;

        _mint(msg.sender, _projectId);
        emit ProjectCreated(msg.sender, _projectId);
    }

    function changeOperator(uint256 _projectId, address _operator) external override onlyProjectOwner(_projectId) {
        require(_operator != address(0), "zero address");
        projects[_projectId].operator = _operator;

        emit ProjectOperatorChanged(_projectId, _operator);
    }

    function pause(uint256 _projectId) external override {
        _requireMinted(_projectId);
        Project storage _project = projects[_projectId];
        _requireOperator(_project.operator);
        require(!_project.paused, "project already paused");

        _project.paused = true;

        emit ProjectPaused(_projectId);
    }

    function resume(uint256 _projectId) external override {
        _requireMinted(_projectId);
        Project storage _project = projects[_projectId];
        _requireOperator(_project.operator);
        require(_project.paused, "project already actived");

        _project.paused = false;

        emit ProjectResumed(_projectId);
    }

    function updateConfig(uint256 _projectId, string calldata _uri, bytes32 _hash) external override {
        _requireMinted(_projectId);
        Project storage _project = projects[_projectId];
        _requireOperator(_project.operator);
        _project.hash = _hash;
        _project.uri = _uri;

        emit ProjectConfigUpdated(_projectId);
    }

    function updateMetadata(uint256 _projectId, bytes32 _name, bytes calldata _value) external override {
        _requireMinted(_projectId);
        Project storage _project = projects[_projectId];
        _requireOperator(_project.operator);
        _project.metadata[_name] = _value;

        emit ProjectMetadataUpdated(_projectId, _name);
    }

    function removeMetadata(uint256 _projectId, bytes32 _name) external override {
        _requireMinted(_projectId);
        Project storage _project = projects[_projectId];
        _requireOperator(_project.operator);

        delete _project.metadata[_name];

        emit ProjectMetadataRemoved(_projectId, _name);
    }

    function withdrawFee(address _account, uint256 _amount) external onlyOwner {
        (bool success, ) = payable(_account).call{value: _amount}("");
        require(success, "withdraw fee fail");

        emit FeeWithdrawn(_account, _amount);
    }
}
