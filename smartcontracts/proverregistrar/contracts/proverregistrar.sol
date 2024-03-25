// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract ProverRegistrar is ERC721, ReentrancyGuard {
    struct Prover {
        string id;
        bool paused;
        mapping(address => bool) operators;
    }

    uint64 private _nextProverID;

    mapping(uint64 => Prover) public provers;
    mapping(string => uint64) public proverIDs;

    constructor() ERC721("ProverToken", "PTK") {
        _nextProverID = 1;
    }

    event OperatorAdded(string indexed proverID, address indexed operator);
    event OperatorRemoved(string indexed proverID, address indexed operator);
    event ProverPaused(string indexed proverID);
    event ProverUnpaused(string indexed proverID);
    event ProverUpserted(string indexed proverID);

    modifier onlyProverOperator(string memory _proverID) {
        require(canOperateProver(msg.sender, _proverID), "Not authorized to operate this prover");
        _;
    }

    modifier onlyProverOwner(string memory _proverID) {
        require(ownerOf(proverIDs[_proverID]) == msg.sender, "Only the owner can perform this action");
        _;
    }

    function canOperateProver(address _operator, string memory _proverID) public view returns (bool) {
        return ownerOf(proverIDs[_proverID]) == _operator || provers[proverIDs[_proverID]].operators[_operator];
    }

    function createProver(string memory _id) public nonReentrant {
        require(bytes(_id).length != 0, "Empty ID value");
        require(proverIDs[_id] == uint64(0), "The ID value has already been registered");

        uint64 proverID = _nextProverID++;
        Prover storage newProver = provers[proverID];
        newProver.id = _id;

        proverIDs[_id] = proverID;

        _mint(msg.sender, proverID);
        emit ProverUpserted(_id);
    }

    function addOperator(string memory _id, address _operator) public onlyProverOwner(_id) {
        provers[proverIDs[_id]].operators[_operator] = true;
        emit OperatorAdded(_id, _operator);
    }

    function removeOperator(string memory _id, address _operator) public onlyProverOwner(_id) {
        provers[proverIDs[_id]].operators[_operator] = false;
        emit OperatorRemoved(_id, _operator);
    }

    function pauseProver(string memory _id) public onlyProverOperator(_id) {
        Prover storage prover = provers[proverIDs[_id]];
        require(!prover.paused, "Prover is already paused");
        prover.paused = true;
        emit ProverPaused(_id);
    }

    function unpauseProver(string memory _id) public onlyProverOperator(_id) {
        Prover storage prover = provers[proverIDs[_id]];
        require(prover.paused, "Prover is not paused");
        prover.paused = false;
        emit ProverUnpaused(_id);
    }
}
