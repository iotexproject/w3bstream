// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProjectRegistrar.sol";

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract ProjectRegistrar is IProjectRegistrar, OwnableUpgradeable, ERC721Upgradeable {
    event FeeWithdrawn(address indexed account, uint256 amount);

    uint256 public override registrationFee;
    mapping(uint256 => mapping(address => bool)) public override operators;
    mapping(uint256 => mapping(bytes32 => bytes)) public override attributes;

    mapping(uint256 => Project) projects;
    uint256 nextProjectId;

    modifier onlyProjectOwner(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender, "not owner");
        _;
    }

    modifier onlyProjectOperator(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender || operators[_projectId][msg.sender], "not operator");
        _;
    }

    function initialize(uint256 _registrationFee, string calldata _name, string calldata _symbol) public initializer {
        __Ownable_init();
        __ERC721_init(_name, _symbol);
        registrationFee = _registrationFee;
        emit RegistrationFeeSet(_registrationFee);
    }

    function isPaused(uint256 _projectId) external view override returns (bool) {
        _requireMinted(_projectId);
        return projects[_projectId].paused;
    }

    function hash(uint256 _projectId) external view override returns (bytes32) {
        _requireMinted(_projectId);
        return projects[_projectId].hash;
    }

    function uri(uint256 _projectId) external view override returns (string memory) {
        _requireMinted(_projectId);
        return projects[_projectId].uri;
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

    function register(string calldata _uri, bytes32 _hash) external payable override returns (uint256 _projectId) {
        require(msg.value > registrationFee, "insufficient fee");
        _projectId = ++nextProjectId;

        Project storage _project = projects[_projectId];
        _project.paused = false;
        _project.hash = _hash;
        _project.uri = _uri;

        _mint(msg.sender, _projectId);
        emit ProjectCreated(msg.sender, _projectId, _uri, _hash);
    }

    function setAttributes(
        uint64 _projectId,
        bytes32[] memory _keys,
        bytes[] memory _values
    ) external onlyProjectOperator(_projectId) {
        require(_keys.length == _values.length, "invalid input");

        mapping(bytes32 => bytes) storage _attributes = attributes[_projectId];
        for (uint i = 0; i < _keys.length; i++) {
            _attributes[_keys[i]] = _values[i];
            emit AttributeSet(_projectId, _keys[i], _values[i]);
        }
    }

    function addOperator(uint256 _projectId, address _operator) external override onlyProjectOwner(_projectId) {
        operators[_projectId][_operator] = true;
        emit OperatorAdded(_projectId, _operator);
    }

    function removeOperator(uint256 _projectId, address _operator) external override onlyProjectOwner(_projectId) {
        operators[_projectId][_operator] = false;
        emit OperatorRemoved(_projectId, _operator);
    }

    function updateProject(
        uint256 _projectId,
        string memory _uri,
        bytes32 _hash
    ) external override onlyProjectOwner(_projectId) {
        require(bytes(_uri).length != 0, "empty uri");
        Project storage project = projects[_projectId];
        project.uri = _uri;
        project.hash = _hash;

        emit ProjectUpdated(_projectId, _uri, _hash);
    }

    function pause(uint256 _projectId) external override onlyProjectOperator(_projectId) {
        Project storage _project = projects[_projectId];
        require(!_project.paused, "project already paused");

        _project.paused = true;

        emit ProjectPaused(_projectId);
    }

    function resume(uint256 _projectId) external override onlyProjectOperator(_projectId) {
        Project storage _project = projects[_projectId];
        require(_project.paused, "project already actived");

        _project.paused = false;

        emit ProjectResumed(_projectId);
    }

    function withdrawFee(address _account, uint256 _amount) external onlyOwner {
        (bool success, ) = payable(_account).call{value: _amount}("");
        require(success, "withdraw fee fail");

        emit FeeWithdrawn(_account, _amount);
    }
}
