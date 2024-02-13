// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract ProjectRegistry is ERC721, ReentrancyGuard {
    struct Project {
        string uri;
        bytes32 hash;
        bool paused;
        address[] operators;
    }

    uint64 private _nextProjectId;

    mapping(uint64 => Project) public projects;

    constructor() ERC721("ProjectToken", "PTK") {
        _nextProjectId = 1;
    }

    event OperatorAdded(uint64 indexed projectId, address indexed operator);
    event OperatorRemoved(uint64 indexed projectId, address indexed operator);
    event ProjectPaused(uint64 indexed projectId);
    event ProjectUnpaused(uint64 indexed projectId);
    event ProjectUpserted(uint64 indexed projectId, string uri, bytes32 hash);

    modifier onlyProjectOperator(uint64 _projectId) {
        require(canOperateProject(msg.sender, _projectId), "ProjectRegistry: Not authorized to operate this project");
        _;
    }

    modifier onlyProjectOwner(uint64 _projectId) {
        require(ownerOf(_projectId) == msg.sender, "ProjectRegistry: Only the owner can perform this action");
        _;
    }

    function canOperateProject(address _operator, uint64 _projectId) public view returns (bool) {
        if (ownerOf(_projectId) == _operator) {
            return true;
        }

        address[] memory operators = projects[_projectId].operators;
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == _operator) {
                return true;
            }
        }

        return false;
    }

    function createProject(string memory _uri, bytes32 _hash) public nonReentrant {
        require(bytes(_uri).length != 0, "Empty uri value");

        uint64 projectId = _nextProjectId++;
        Project storage newProject = projects[projectId];
        newProject.uri = _uri;
        newProject.hash = _hash;

        _mint(msg.sender, projectId);
        emit ProjectUpserted(projectId, _uri, _hash);
    }

    function addOperator(uint64 _projectId, address _operator) public onlyProjectOwner(_projectId) {
        for (uint256 i = 0; i < projects[_projectId].operators.length; i++) {
            require(projects[_projectId].operators[i] != _operator, "ProjectRegistry: Operator already added");
        }

        projects[_projectId].operators.push(_operator);
        emit OperatorAdded(_projectId, _operator);
    }

    function getOperators(uint64 _projectId) public view returns (address[] memory) {
        return projects[_projectId].operators;
    }

    function removeOperator(uint64 _projectId, address _operator) public onlyProjectOwner(_projectId) {
        address[] storage operators = projects[_projectId].operators;

        bool isFound;

        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == _operator) {
                operators[i] = operators[operators.length - 1];
                operators.pop();
                isFound = true;
                break;
            }
        }

        require(isFound, "ProjectRegistry: Operator not found");

        emit OperatorRemoved(_projectId, _operator);
    }

    function pauseProject(uint64 _projectId) public onlyProjectOperator(_projectId) {
        Project storage project = projects[_projectId];

        require(!project.paused, "ProjectRegistry: Project already paused");
        project.paused = true;
        emit ProjectPaused(_projectId);
    }

    function unpauseProject(uint64 _projectId) public onlyProjectOperator(_projectId) {
        Project storage project = projects[_projectId];
        require(project.paused, "ProjectRegistry: Project is not paused");
        project.paused = false;
        emit ProjectUnpaused(_projectId);
    }

    function updateProject(
        uint64 _projectId,
        string memory _uri,
        bytes32 _hash
    ) public onlyProjectOperator(_projectId) {
        require(bytes(_uri).length != 0, "ProjectRegistry: Invalid URI");

        projects[_projectId].uri = _uri;
        projects[_projectId].hash = _hash;
        emit ProjectUpserted(_projectId, _uri, _hash);
    }
}
