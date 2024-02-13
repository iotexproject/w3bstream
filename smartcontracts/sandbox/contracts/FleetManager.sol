// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IFleetManager} from "./interfaces/IFleetManager.sol";
import {IOperatorRegistry} from "./interfaces/IOperatorRegistry.sol";
import {IProjectRegistry} from "./interfaces/IProjectRegistry.sol";

contract FleetManager is IFleetManager {
    address public projectRegistry;
    address public operatorRegistry;

    mapping(uint256 => address[]) public operators;

    event OperatorAdded(uint256 indexed projectId, address indexed operator);
    event OperatorRemoved(uint256 indexed projectId, address indexed operator);

    constructor(address _projectRegistry, address _operatorRegistry) {
        projectRegistry = _projectRegistry;
        operatorRegistry = _operatorRegistry;
    }

    modifier onlyProjectOwner(uint256 _projectId) {
        require(
            IProjectRegistry(projectRegistry).isProjectOwner(msg.sender, _projectId),
            "FleetManager: not project owner"
        );
        _;
    }

    modifier onlyValidOperator(address _operator) {
        address node = IOperatorRegistry(operatorRegistry).getOperator(_operator).node;
        require(node != address(0), "FleetManager: operator not registered");

        // NEED TO CHECK IF THE OPERATOR HAS ENOUGH STAKE
        _;
    }

    modifier notAlreadyAllowed(uint256 _projectId, address _operator) {
        address[] memory projectOperators = operators[_projectId];

        for (uint256 i = 0; i < projectOperators.length; i++) {
            require(projectOperators[i] != _operator, "FleetManager: operator already allowed");
        }
        _;
    }

    function allow(
        uint256 _projectId,
        address _operator
    ) external onlyProjectOwner(_projectId) onlyValidOperator(_operator) notAlreadyAllowed(_projectId, _operator) {
        operators[_projectId].push(_operator);
        emit OperatorAdded(_projectId, _operator);
    }

    function disallow(uint256 _projectId, address _operator) external onlyProjectOwner(_projectId) {
        address[] storage projectOperators = operators[_projectId];

        bool found;

        for (uint256 i = 0; i < projectOperators.length; i++) {
            if (projectOperators[i] == _operator) {
                projectOperators[i] = projectOperators[projectOperators.length - 1];
                projectOperators.pop();
                found = true;
                break;
            }
        }

        require(found, "FleetManager: operator not found");

        emit OperatorRemoved(_projectId, _operator);
    }

    function isAllowed(address _node, uint256 _projectId) external view returns (bool) {
        require(_node != address(0), "FleetManager: invalid node");

        address[] memory projectOperators = operators[_projectId];

        for (uint256 i = 0; i < projectOperators.length; i++) {
            if (_getOperatorNode(projectOperators[i]) == _node) {
                return true;
            }
        }

        return false;
    }

    function _getOperatorNode(address _operator) internal view returns (address) {
        return IOperatorRegistry(operatorRegistry).getOperator(_operator).node;
    }
}
