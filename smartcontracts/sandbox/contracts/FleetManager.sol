// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {IFleetManager} from "./interfaces/IFleetManager.sol";
import {IOperatorRegistry} from "./interfaces/IOperatorRegistry.sol";

contract FleetManager is IFleetManager, Initializable {
    address public projectRegistry;
    address public operatorRegistry;

    mapping(uint256 => address[]) public operators;

    event OperatorAdded(uint256 indexed projectId, address indexed operator);
    event OperatorRemoved(uint256 indexed projectId, address indexed operator);

    function initialize(address _projectRegistry, address _operatorRegistry) public initializer {
        projectRegistry = _projectRegistry;
        operatorRegistry = _operatorRegistry;
    }

    modifier onlyProjectOwner(uint256 _projectId) {
        if (IERC721(projectRegistry).ownerOf(_projectId) != msg.sender) {
            revert NotProjectOwner();
        }
        _;
    }

    modifier onlyValidOperator(address _operator) {
        address node = IOperatorRegistry(operatorRegistry).getOperator(_operator).node;
        if (node == address(0)) {
            revert OperatorNotRegistered();
        }

        // NEED TO CHECK IF THE OPERATOR HAS ENOUGH STAKE
        _;
    }

    modifier notAlreadyAllowed(uint256 _projectId, address _operator) {
        address[] memory projectOperators = operators[_projectId];

        for (uint256 i = 0; i < projectOperators.length; i++) {
            if (projectOperators[i] == _operator) {
                revert OperatorAlreadyAllowed();
            }
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

        if (!found) {
            revert OperatorNotFound();
        }

        emit OperatorRemoved(_projectId, _operator);
    }

    function isAllowed(address _node, uint256 _projectId) external view returns (bool) {
        if (_node == address(0)) {
            revert InvalidNodeAddress();
        }

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
