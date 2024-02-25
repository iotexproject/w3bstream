pragma solidity ^0.8.0;

contract Store {
    bytes private proof;
    bytes private proof_seal;
    bytes private proof_post_state_digest;
    bytes private proof_journal;
    uint256 private projectId;
    address private receiver;


    function setProof(bytes memory _proof) public {
        proof = _proof;
    }

    function submit(uint256 _projectId, address _receiver, bytes calldata _data_snark) public {
        projectId = _projectId;
        receiver = _receiver;
        (bytes memory proof_snark_seal, bytes memory proof_snark_post_state_digest, bytes memory proof_snark_journal) = abi.decode(_data_snark, (bytes, bytes, bytes));
        proof_seal = proof_snark_seal;
        proof_post_state_digest = proof_snark_post_state_digest;
        proof_journal = proof_snark_journal;
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

    function getReceiver() public view returns (address ){
        return receiver;
    }
}
