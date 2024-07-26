package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var (
	operatorSetTopic   = crypto.Keccak256Hash([]byte("OperatorSet(uint256,address)"))
	vmTypeAddedTopic   = crypto.Keccak256Hash([]byte("VMTypeAdded(uint256,uint256)"))
	vmTypeDeletedTopic = crypto.Keccak256Hash([]byte("VMTypeDeleted(uint256,uint256)"))
	proverPausedTopic  = crypto.Keccak256Hash([]byte("ProverPaused(uint256)"))
	proverResumedTopic = crypto.Keccak256Hash([]byte("ProverResumed(uint256)"))
)

type Prover struct {
	ID              uint64
	OperatorAddress common.Address
	Paused          bool
	VMTypes         map[uint64]bool
}

type vmTypeUpdated struct {
	isAdded bool
	typ     uint64
}

type proverDiff struct {
	id              uint64
	operatorAddress *common.Address
	paused          *bool
	vmTypesUpdated  []vmTypeUpdated
}

type blockProver struct {
	Provers map[uint64]*Prover
}

type blockProverDiff struct {
	diffs map[uint64]*proverDiff
}

func newProver() *Prover {
	return &Prover{
		Paused:  true,
		VMTypes: map[uint64]bool{},
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
	for _, u := range diff.vmTypesUpdated {
		if u.isAdded {
			p.VMTypes[u.typ] = true
		} else {
			delete(p.VMTypes, u.typ)
		}
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
		case operatorSetTopic:
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

		case vmTypeAddedTopic:
			e, err := c.proverInstance.ParseVMTypeAdded(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse prover vm type added event")
			}

			p, ok := ps.diffs[e.Id.Uint64()]
			if !ok {
				p = &proverDiff{id: e.Id.Uint64()}
			}
			nt := e.Typ.Uint64()
			p.vmTypesUpdated = append(p.vmTypesUpdated, vmTypeUpdated{isAdded: true, typ: nt})
			ps.diffs[e.Id.Uint64()] = p

		case vmTypeDeletedTopic:
			e, err := c.proverInstance.ParseVMTypeDeleted(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse prover vm type deleted event")
			}

			p, ok := ps.diffs[e.Id.Uint64()]
			if !ok {
				p = &proverDiff{id: e.Id.Uint64()}
			}
			nt := e.Typ.Uint64()
			p.vmTypesUpdated = append(p.vmTypesUpdated, vmTypeUpdated{isAdded: false, typ: nt})
			ps.diffs[e.Id.Uint64()] = p

		case proverPausedTopic:
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

		case proverResumedTopic:
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
