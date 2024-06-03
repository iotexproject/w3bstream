package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var (
	operatorSetTopicHash     = crypto.Keccak256Hash([]byte("OperatorSet(uint256,address)"))
	nodeTypeUpdatedTopicHash = crypto.Keccak256Hash([]byte("NodeTypeUpdated(uint256,uint256)"))
	proverPausedTopicHash    = crypto.Keccak256Hash([]byte("ProverPaused(uint256)"))
	proverResumedTopicHash   = crypto.Keccak256Hash([]byte("ProverResumed(uint256)"))

	emptyAddress = common.Address{}
)

type Prover struct {
	ID              uint64
	OperatorAddress common.Address
	Paused          bool
	NodeTypes       uint64
}

type proverDiff struct {
	id              uint64
	operatorAddress *common.Address
	paused          *bool
	nodeTypes       *uint64
}

type blockProver struct {
	Provers map[uint64]*Prover
}

type blockProverDiff struct {
	diffs map[uint64]*proverDiff
}

func newProver() *Prover {
	return &Prover{
		Paused: true,
	}
}

func (p *Prover) merge(diff *proverDiff) {
	if diff.id != 0 {
		p.ID = diff.id
	}
	if diff.operatorAddress != nil {
		p.OperatorAddress = *diff.operatorAddress
	}
	if diff.paused != nil {
		p.Paused = *diff.paused
	}
	if diff.nodeTypes != nil {
		p.NodeTypes = *diff.nodeTypes
	}
}

func (ps *blockProver) merge(diff *blockProverDiff) {
	for id, p := range ps.Provers {
		diffP, ok := diff.diffs[id]
		if ok {
			p.merge(diffP)
		}
	}
	for id, p := range diff.diffs {
		if _, ok := ps.Provers[id]; !ok {
			np := newProver()
			np.merge(p)
			ps.Provers[id] = np
		}
	}
}

// return blockNumber -> *blockProverDiff
func (c *Contract) processProverLogs(logs []types.Log) (map[uint64]*blockProverDiff, error) {
	psMap := map[uint64]*blockProverDiff{}

	for _, l := range logs {
		ps, ok := psMap[l.BlockNumber]
		if !ok {
			ps = &blockProverDiff{
				diffs: map[uint64]*proverDiff{},
			}
		}
		switch l.Topics[0] {
		case operatorSetTopicHash:
			e, err := c.proverInstance.ParseOperatorSet(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse prover operator set event")
			}

			p, ok := ps.diffs[e.Id.Uint64()]
			if !ok {
				p = &proverDiff{id: e.Id.Uint64()}
			}
			p.operatorAddress = &e.Operator
			ps.diffs[e.Id.Uint64()] = p

		case nodeTypeUpdatedTopicHash:
			e, err := c.proverInstance.ParseNodeTypeUpdated(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse prover node type updated event")
			}

			p, ok := ps.diffs[e.Id.Uint64()]
			if !ok {
				p = &proverDiff{id: e.Id.Uint64()}
			}
			nt := e.Typ.Uint64()
			p.nodeTypes = &nt
			ps.diffs[e.Id.Uint64()] = p

		case proverPausedTopicHash:
			e, err := c.proverInstance.ParseProverPaused(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse prover paused event")
			}

			p, ok := ps.diffs[e.Id.Uint64()]
			if !ok {
				p = &proverDiff{id: e.Id.Uint64()}
			}
			paused := true
			p.paused = &paused
			ps.diffs[e.Id.Uint64()] = p

		case proverResumedTopicHash:
			e, err := c.proverInstance.ParseProverResumed(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse prover resumed event")
			}

			p, ok := ps.diffs[e.Id.Uint64()]
			if !ok {
				p = &proverDiff{id: e.Id.Uint64()}
			}
			paused := false
			p.paused = &paused
			ps.diffs[e.Id.Uint64()] = p
		}
		psMap[l.BlockNumber] = ps
	}

	return psMap, nil
}
