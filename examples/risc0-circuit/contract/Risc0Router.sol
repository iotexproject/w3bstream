// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

interface IRisc0Receiver {
    function verifySnark(bytes calldata seal, bytes32 imageId, bytes32 postStateDigest, bytes memory journal) external view returns (bool);
}

contract Risc0Router {
    bytes private proof;
    bytes private proof_seal;
    bytes private proof_post_state_digest;
    bytes private proof_journal;
    uint256 private projectId;
    uint256 private taskId;
    address private receiver;

    IRisc0Receiver public risc0Receiver;

    // TODO move to depin contract
    mapping(uint256 => bytes32) private projectIdToImageId;

    function setProof(bytes memory _proof) public {
        proof = _proof;
    }

    function submit(address _receiver, uint256 _projectId, uint256 _taskId, bytes calldata _data_snark) public {
        projectId = _projectId;
        receiver = _receiver;
        taskId = _taskId;
        (bytes memory proof_snark_seal, bytes memory proof_snark_post_state_digest, bytes memory proof_snark_journal) = abi.decode(_data_snark, (bytes, bytes, bytes));
        proof_seal = proof_snark_seal;
        proof_post_state_digest = proof_snark_post_state_digest;
        proof_journal = proof_snark_journal;
        // TODO receiver is risc0 verification contract
        // risc0Receiver = IRisc0Receiver(receiver);
        // bool success = risc0Receiver.verifySnark(proof_seal, projectIdToImageId[projectId], proof_post_state_digest, proof_journal);
        // require(success, "Failed to verify proof");
    }

    function setProjectIdToImageId(uint256 _projectId, bytes32 _imageId) public {
        projectIdToImageId[_projectId] = _imageId;
    }

    function getImageIdByProjectId(uint256 _projectId) public view returns (bytes32) {
        return projectIdToImageId[_projectId];
    }

    function getSeal() public view returns (bytes memory){
        return proof_seal;
    }

    function getPostStateDigest() public view returns (bytes memory){
        return proof_post_state_digest;
    }

    function getJournal() public view returns (bytes memory){
        return proof_journal;
    }

    function getProjectId() public view returns (uint256){
        return projectId;
    }

    function getTaskId() public view returns (uint256){
        return taskId;
    }

    function getReceiver() public view returns (address ){
        return receiver;
    }
}
