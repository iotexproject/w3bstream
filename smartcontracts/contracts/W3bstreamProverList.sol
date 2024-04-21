// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./W3bstreamProver.sol";

contract W3bstreamProverList {

    struct Prover {
        uint256 id;
        uint256 nodeType;
        address operator;
        bool isPaused;
    }

    function list(address _proverContract) external view returns (uint256 blockNumber_, Prover[] memory provers_) {
       W3bstreamProver w3bstreamProver = W3bstreamProver(_proverContract);
       uint256 count = w3bstreamProver.count();
       provers_ = new Prover[](count);
       blockNumber_ = block.number;

        for (uint256 i = 1; i <= count; i++) {
            Prover memory prover;

            prover.id = i;
            prover.nodeType = w3bstreamProver.nodeType(i);
            prover.operator = w3bstreamProver.operator(i);
            prover.isPaused = w3bstreamProver.isPaused(i);

            provers_[i-1] = prover;
        }
    }
}
