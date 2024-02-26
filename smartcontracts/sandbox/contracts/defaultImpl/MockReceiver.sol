// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract MockRisc0SnarkReceiver is Initializable {
    address public _verify;

    function initialize(address _risc0Verify) public initializer {
        _verify = _risc0Verify;
    }

    bytes private proof_seal;
    bytes private proof_post_state_digest;
    bytes private proof_journal;

    function receiveData(bytes calldata _data) external {
        (
            bytes memory proof_snark_seal,
            bytes memory proof_snark_post_state_digest,
            bytes memory proof_snark_journal
        ) = abi.decode(_data, (bytes, bytes, bytes));
        proof_seal = proof_snark_seal;
        proof_post_state_digest = proof_snark_post_state_digest;
        proof_journal = proof_snark_journal;
    }

    function getSeal() public view returns (bytes memory) {
        return proof_seal;
    }

    function getPostStateDigest() public view returns (bytes memory) {
        return proof_post_state_digest;
    }

    function getJournal() public view returns (bytes memory) {
        return proof_journal;
    }

    function setVerify(address _risc0Verify) external {
      _verify = _risc0Verify;
    }
}
