package message

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/output/chain/eth"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/test/contract"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
)

type Processor struct {
	vmHandler          *vm.Handler
	projectManager     *project.Manager
	chainEndpoint      string
	operatorPrivateKey string
	ps                 *p2p.PubSubs
}

func NewProcessor(vmHandler *vm.Handler, projectManager *project.Manager, chainEndpoint, operatorPrivateKey, bootNodeMultiaddr string, iotexChainID int) (*Processor, error) {
	p := &Processor{
		vmHandler:          vmHandler,
		chainEndpoint:      chainEndpoint,
		operatorPrivateKey: operatorPrivateKey,
		projectManager:     projectManager,
	}

	ps, err := p2p.NewPubSubs(p.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	p.ps = ps
	return p, nil
}

func (r *Processor) Run() {
	go r.runMessageRequest()
}

// TODO support batch message fetch & fetch frequency define
func (r *Processor) runMessageRequest() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	projectIDs := r.projectManager.GetAllProjectID()

	for range ticker.C {
		for _, projectID := range projectIDs {
			if err := r.ps.Add(projectID); err != nil {
				slog.Error("add project pubsub failed", "error", err)
				continue
			}

			d := &p2p.Data{
				Request: &p2p.RequestData{
					ProjectID: projectID,
				},
			}
			if err := r.ps.Publish(projectID, d); err != nil {
				slog.Error("publish data failed", "error", err)
			}
		}
	}
}

func (r *Processor) handleP2PData(d *p2p.Data, topic *pubsub.Topic) {
	if d.Message == nil || len(d.Message.Messages) == 0 {
		return
	}
	ms := d.Message.Messages
	mids := r.getMessageIDs(ms)
	slog.Debug("get new messages", "message_ids", mids)
	r.reportSuccess(mids, types.MessageStateFetched, "", topic)

	project, err := r.projectManager.Get(ms[0].ProjectID)
	if err != nil {
		slog.Error("get project failed", "error", err)
		r.reportFail(mids, err, topic)
		return
	}

	r.reportSuccess(mids, types.MessageStateProving, "", topic)
	res, err := r.vmHandler.Handle(ms, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed", "error", err)
		r.reportFail(mids, err, topic)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	r.reportSuccess(mids, types.MessageStateProved, string(res), topic)

	if r.operatorPrivateKey == "" {
		info := "missing operator private key, will not write to chain"
		slog.Debug(info)
		r.reportSuccess(mids, types.MessageStateOutputted, info, topic)
		return
	}

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail(mids, err, topic)
		return
	}

	slog.Debug("writing proof to chain")

	r.reportSuccess(mids, types.MessageStateOutputting, "writing proof to chain", topic)
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x6e30b42554DDA34bAFca9cB00Ec4B464f452a671", data)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail(mids, err, topic)
		return
	}
	r.reportSuccess(mids, types.MessageStateOutputted, txHash, topic)
	slog.Debug("transaction hash", "tx_hash", txHash)
}

func (r *Processor) reportFail(messageIDs []string, err error, topic *pubsub.Topic) {
	d := p2p.Data{
		Response: &p2p.ResponseData{
			MessageIDs: messageIDs,
			State:      types.MessageStateFailed,
			Comment:    err.Error(),
			CreatedAt:  time.Now(),
		},
	}
	j, err := json.Marshal(&d)
	if err != nil {
		slog.Error("json marshal p2p response data failed", "error", err, "messageIDs", messageIDs)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish response data to p2p network failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) reportSuccess(messageIDs []string, state types.MessageState, comment string, topic *pubsub.Topic) {
	d := p2p.Data{
		Response: &p2p.ResponseData{
			MessageIDs: messageIDs,
			State:      state,
			Comment:    comment,
			CreatedAt:  time.Now(),
		},
	}
	j, err := json.Marshal(&d)
	if err != nil {
		slog.Error("json marshal p2p response data failed", "error", err, "messageIDs", messageIDs)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish response data to p2p network failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) getMessageIDs(ms []*types.Message) []string {
	ids := []string{}
	for _, m := range ms {
		ids = append(ids, m.MessageID)
	}
	return ids
}
