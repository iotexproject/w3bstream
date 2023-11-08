// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract ProjectRegistrar is ERC721, ReentrancyGuard {
    struct Project {
        string uri;
        bytes32 hash;
        bool paused;
        mapping(address => bool) operators;
    }

    mapping(uint256 => Project) public projects;

    constructor() ERC721("ProjectToken", "PTK") {}

    event OperatorAdded(uint256 indexed projectId, address indexed operator);
    event OperatorRemoved(uint256 indexed projectId, address indexed operator);
    event ProjectPaused(uint256 indexed projectId);
    event ProjectUnpaused(uint256 indexed projectId);
    event ProjectUpdated(uint256 indexed projectId, string uri, bytes32 hash);

    modifier onlyProjectOperator(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender || projects[_projectId].operators[msg.sender], "Only the owner or operators can perform this action");
        _;
    }

    modifier onlyProjectOwner(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender, "Only the owner can perform this action");
        _;
    }

    function createProject(string memory _uri, bytes32 _hash) public nonReentrant {
        uint256 projectId = totalSupply();
        Project storage newProject = projects[projectId];
        newProject.uri = _uri;
        newProject.hash = _hash;

        _mint(msg.sender, projectId);
        newProject.operators[msg.sender] = true;
    }

    function addOperator(uint256 _projectId, address _operator) public onlyProjectOwner(_projectId) {
        projects[_projectId].operators[_operator] = true;
        emit OperatorAdded(_projectId, _operator);
    }

    function removeOperator(uint256 _projectId, address _operator) public onlyProjectOwner(_projectId) {
        projects[_projectId].operators[_operator] = false;
        emit OperatorRemoved(_projectId, _operator);
    }

    function pauseProject(uint256 _projectId) public onlyProjectOperator(_projectId) {
        Project storage project = projects[_projectId];
        require(!project.paused, "Project is already paused");
        project.paused = true;
        emit ProjectPaused(_projectId);
    }

    function unpauseProject(uint256 _projectId) public onlyProjectOperator(_projectId) {
        Project storage project = projects[_projectId];
        require(project.paused, "Project is not paused");
        project.paused = false;
        emit ProjectUnpaused(_projectId);
    }

    function updateProject(uint256 _projectId, string memory _uri, bytes32 _hash) public onlyProjectOperator(_projectId) {
        projects[_projectId].uri = _uri;
        projects[_projectId].hash = _hash;
        emit ProjectUpdated(_projectId, _uri, _hash);
    }
}
