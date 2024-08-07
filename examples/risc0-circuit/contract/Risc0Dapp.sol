// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

interface IRisc0Receiver {
    function verify(bytes calldata seal, bytes32 imageId, bytes32 journalDigest) external view;
}

contract Risc0Dapp {

    bytes public proof;
    bytes public proof_seal;
    bytes32 public proof_journal;
    uint256 public projectId;
    uint256 public proverId;
    string public clientId;
    // risc0 verification contract
    address private risc0Verifier;

    mapping(uint256 => bytes32) private projectIdToImageId;


    function process(uint256 _projectId, uint256 _proverId, string memory _clientId, bytes calldata _data) public {
        projectId = _projectId;
        proverId = _proverId;
        clientId = _clientId;
        proof = _data;
        (bytes memory proof_snark_seal, bytes memory proof_snark_journal) = abi.decode(_data, (bytes, bytes));
        proof_seal = proof_snark_seal;
        proof_journal = sha256(proof_snark_journal);
        // verify zk proof
        IRisc0Receiver(risc0Verifier).verify(proof_seal, projectIdToImageId[projectId], proof_journal);
        // TODO
    }

    function setProjectIdToImageId(uint256 _projectId, bytes32 _imageId) public {
        projectIdToImageId[_projectId] = _imageId;
    }

    function getImageIdByProjectId(uint256 _projectId) public view returns (bytes32) {
        return projectIdToImageId[_projectId];
    }
    
    function setReceiver(address _receiver) public {
        risc0Verifier = _receiver;
    }

    function getReceiver() public view returns (address ){
        return risc0Verifier;
    }
}