// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract W3bstreamProver is OwnableUpgradeable, ERC721Upgradeable {
    event OperatorSet(uint256 indexed id, address indexed operator);
    event NodeTypeUpdated(uint256 indexed id, uint256 typ);
    event ProverPaused(uint256 indexed id);
    event ProverResumed(uint256 indexed id);
    event MinterSet(address minter);

    address public minter;
    uint256 nextProverId;

    mapping(uint256 => uint256) _nodeTypes;
    mapping(uint256 => address) _operators;
    mapping(uint256 => bool) _paused;
    mapping(address => uint256) operatorToProver;

    modifier onlyProverOwner(uint256 _id) {
        require(ownerOf(_id) == msg.sender, "not owner");
        _;
    }

    function initialize(string memory _name, string memory _symbol) public initializer {
        __Ownable_init();
        __ERC721_init(_name, _symbol);
        setMinter(msg.sender);
    }

    function count() external view returns (uint256) {
        return nextProverId;
    }

    function nodeType(uint256 _id) external view returns (uint256) {
        _requireMinted(_id);
        return _nodeTypes[_id];
    }

    function operator(uint256 _id) external view returns (address) {
        _requireMinted(_id);
        return _operators[_id];
    }

    function prover(uint256 _id) external view returns (address) {
        _requireMinted(_id);
        return ownerOf(_id);
    }

    function ownerOfOperator(address _operator) external view returns (uint256, address) {
        uint256 id = operatorToProver[_operator];
        require(id != 0, "invalid operator");
        return (id, ownerOf(id));
    }

    function isPaused(uint256 _id) external view returns (bool) {
        _requireMinted(_id);
        return _paused[_id];
    }

    function mint(address _account) external returns (uint256 id_) {
        require(msg.sender == minter, "not minter");
        id_ = 1;
        ++nextProverId;
        _mint(_account, id_);

        _paused[id_] = true;
        updateOperatorInternal(id_, _account);
    }

    function updateOperatorInternal(uint256 _id, address _operator) internal {
        uint256 proverId = operatorToProver[_operator];
        require(proverId == 0, "invalid operator");
        address oldOperator = _operators[_id];
        _operators[_id] = _operator;
        delete operatorToProver[oldOperator];
        operatorToProver[_operator] = _id;
        emit OperatorSet(_id, _operator);
    }

    function updateNodeType(uint256 _id, uint256 _type) external onlyProverOwner(_id) {
        _nodeTypes[_id] = _type;
        emit NodeTypeUpdated(_id, _type);
    }

    function changeOperator(uint256 _id, address _operator) external onlyProverOwner(_id) {
        require(_operator != address(0), "zero address");
        updateOperatorInternal(_id, _operator);
    }

    function pause(uint256 _id) external onlyProverOwner(_id) {
        require(!_paused[_id], "already paused");

        _paused[_id] = true;
        emit ProverPaused(_id);
    }

    function resume(uint256 _id) external onlyProverOwner(_id) {
        require(_paused[_id], "already actived");

        _paused[_id] = false;
        emit ProverResumed(_id);
    }

    function setMinter(address _minter) public onlyOwner {
        minter = _minter;

        emit MinterSet(_minter);
    }
}
