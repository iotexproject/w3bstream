// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract ZNodeRegistrar is ERC721, ReentrancyGuard {
    struct ZNode {
        string did;
        bool paused;
        mapping(address => bool) operators;
    }

    uint64 private _nextZNodeID;

    mapping(uint64 => ZNode) public znodes;
    mapping(string => uint64) public znodeDIDs;

    constructor() ERC721("ZNodeToken", "ZTK") {
        _nextZNodeID = 1;
    }

    event OperatorAdded(string indexed znodeDID, address indexed operator);
    event OperatorRemoved(string indexed znodeDID, address indexed operator);
    event ZNodePaused(string indexed znodeDID);
    event ZNodeUnpaused(string indexed znodeDID);
    event ZNodeUpserted(string indexed znodeDID);

    modifier onlyZNodeOperator(string memory _znodeDID) {
        require(canOperateZNode(msg.sender, _znodeDID), "Not authorized to operate this znode");
        _;
    }

    modifier onlyZNodeOwner(string memory _znodeDID) {
        require(ownerOf(znodeDIDs[_znodeDID]) == msg.sender, "Only the owner can perform this action");
        _;
    }

    function canOperateZNode(address _operator, string memory _znodeDID) public view returns (bool) {
        return ownerOf(znodeDIDs[_znodeDID]) == _operator || znodes[znodeDIDs[_znodeDID]].operators[_operator];
    }

    function createZNode(string memory _did) public nonReentrant {
        require(bytes(_did).length != 0, "Empty DID value");
        require(znodeDIDs[_did] == uint64(0), "The DID value has already been registered");

        uint64 znodeID = _nextZNodeID++;
        ZNode storage newZNode = znodes[znodeID];
        newZNode.did = _did;

        znodeDIDs[_did] = znodeID;

        _mint(msg.sender, znodeID);
        emit ZNodeUpserted(_did);
    }

    function addOperator(string memory _did, address _operator) public onlyZNodeOwner(_did) {
        znodes[znodeDIDs[_did]].operators[_operator] = true;
        emit OperatorAdded(_did, _operator);
    }

    function removeOperator(string memory _did, address _operator) public onlyZNodeOwner(_did) {
        znodes[znodeDIDs[_did]].operators[_operator] = false;
        emit OperatorRemoved(_did, _operator);
    }

    function pauseZNode(string memory _did) public onlyZNodeOperator(_did) {
        ZNode storage znode = znodes[znodeDIDs[_did]];
        require(!znode.paused, "ZNode is already paused");
        znode.paused = true;
        emit ZNodePaused(_did);
    }

    function unpauseZNode(string memory _did) public onlyZNodeOperator(_did) {
        ZNode storage znode = znodes[znodeDIDs[_did]];
        require(znode.paused, "ZNode is not paused");
        znode.paused = false;
        emit ZNodeUnpaused(_did);
    }
}
