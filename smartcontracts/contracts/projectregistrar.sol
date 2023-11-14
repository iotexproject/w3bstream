// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract ProjectRegistrar is ERC721, ReentrancyGuard {
    struct Project {
        string uri;
        bytes32 hash;
        bool paused;
        mapping(address => bool) operators;
    }

    uint256 private _nextProjectId;

    mapping(uint256 => Project) public projects;

    constructor() ERC721("ProjectToken", "PTK") {
        _nextProjectId = 1;
    }

    event OperatorAdded(uint256 indexed projectId, address indexed operator);
    event OperatorRemoved(uint256 indexed projectId, address indexed operator);
    event ProjectPaused(uint256 indexed projectId);
    event ProjectUnpaused(uint256 indexed projectId);
    event ProjectUpserted(uint256 indexed projectId, string uri, bytes32 hash);

    modifier onlyProjectOperator(uint256 _projectId) {
        require(canOperateProject(msg.sender, _projectId), "Not authorized to operate this project");
        _;
    }

    modifier onlyProjectOwner(uint256 _projectId) {
        require(ownerOf(_projectId) == msg.sender, "Only the owner can perform this action");
        _;
    }

    function canOperateProject(address _operator, uint256 _projectId) public view returns (bool) {
        return ownerOf(_projectId) == _operator || projects[_projectId].operators[_operator];
    }

    function createProject(string memory _uri, bytes32 _hash) public nonReentrant {
        uint256 projectId = _nextProjectId++;
        Project storage newProject = projects[projectId];
        newProject.uri = _uri;
        newProject.hash = _hash;

        _mint(msg.sender, projectId);
        emit ProjectUpserted(projectId, _uri, _hash);
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
        emit ProjectUpserted(_projectId, _uri, _hash);
    }
}
