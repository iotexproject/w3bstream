pragma solidity ^0.5.0;

contract Store {
    bytes private proof;

    function setProof(bytes memory _proof) public {
        proof = _proof;
    }

    function setProofSnark(bytes calldata proof_snark_seal, bytes calldata proof_snark_post_state_digest, bytes calldata proof_snark_journal) public {
        proof = proof_snark_seal;
    }

    function getProof() public view returns (bytes memory){
        return proof;
    }
}
