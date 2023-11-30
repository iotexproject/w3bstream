package message

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/output/chain/eth"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/test/contract"
	"github.com/machinefi/sprout/vm"
	"github.com/pkg/errors"
)

type Processor struct {
	vmHandler          *vm.Handler
	projectManager     *project.Manager
	chainEndpoint      string
	operatorPrivateKey string
	topic              *pubsub.Topic
	sub                *pubsub.Subscription
}

func NewProcessor(vmHandler *vm.Handler, projectManager *project.Manager, chainEndpoint, p2pMultiaddr, operatorPrivateKey string) (*Processor, error) {
	h, err := libp2p.New(libp2p.ListenAddrStrings(p2pMultiaddr))
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host failed")
	}
	ctx := context.Background()
	go p2p.DiscoverPeers(ctx, h, p2p.Topic)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, errors.Wrap(err, "new gossip subscription failed")
	}
	topic, err := ps.Join(p2p.Topic)
	if err != nil {
		return nil, errors.Wrap(err, "join topic failed")
	}
	//go streamConsoleTo(ctx, topic)

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, errors.Wrap(err, "topic subscription failed")
	}
	//printMessagesFrom(ctx, sub)

	p := &Processor{
		vmHandler:          vmHandler,
		chainEndpoint:      chainEndpoint,
		operatorPrivateKey: operatorPrivateKey,
		projectManager:     projectManager,
		topic:              topic,
		sub:                sub,
	}
	return p, nil
}

func (r *Processor) Run() {
	go r.runMessageRequest()
	go r.runMessageProcess()
}

// TODO support batch message fetch & fetch frequency define
func (r *Processor) runMessageRequest() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	projectIDs := r.projectManager.GetAllProjectID()

	for range ticker.C {
		for _, projectID := range projectIDs {
			d := p2p.Data{
				Type:      p2p.Request,
				ProjectID: projectID,
			}
			j, err := json.Marshal(&d)
			if err != nil {
				slog.Error("json marshal p2p data failed", "error", err)
				continue
			}

			if err := r.topic.Publish(context.Background(), j); err != nil {
				slog.Error("publish data to p2p network failed", "error", err)
			}
		}
	}
}

func (r *Processor) runMessageProcess() {
	for {
		m, err := r.sub.Next(context.Background())
		if err != nil {
			slog.Error("get p2p data failed", "error", err)
			continue
		}
		d := p2p.Data{}
		if err := json.Unmarshal(m.Message.Data, &d); err != nil {
			slog.Error("json unmarshal p2p data failed", "error", err)
			continue
		}
		if d.Type != p2p.Message {
			continue
		}
		if len(d.Messages) != 0 {
			r.process(d.Messages)
		}
	}
}

func (r *Processor) process(ms []*proto.Message) {
	mids := r.getMessageIDs(ms)
	slog.Debug("get new messages", "message_ids", mids)
	r.reportSuccess(mids, proto.MessageState_FETCHED, "")

	project, err := r.projectManager.Get(ms[0].ProjectID)
	if err != nil {
		slog.Error("get project failed", "error", err)
		r.reportFail(mids, err)
		return
	}

	r.reportSuccess(mids, proto.MessageState_PROVING, "")
	res, err := r.vmHandler.Handle(ms, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed", "error", err)
		r.reportFail(mids, err)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	r.reportSuccess(mids, proto.MessageState_PROVED, string(res))

	if r.operatorPrivateKey == "" {
		info := "missing operator private key, will not write to chain"
		slog.Debug(info)
		r.reportSuccess(mids, proto.MessageState_OUTPUTTED, info)
		return
	}

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail(mids, err)
		return
	}

	slog.Debug("writing proof to chain")

	r.reportSuccess(mids, proto.MessageState_OUTPUTTING, "writing proof to chain")
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x6e30b42554DDA34bAFca9cB00Ec4B464f452a671", data)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail(mids, err)
		return
	}
	r.reportSuccess(mids, proto.MessageState_OUTPUTTED, txHash)
	slog.Debug("transaction hash", "tx_hash", txHash)
}

func (r *Processor) reportFail(messageIDs []string, err error) {
	d := p2p.Data{
		Type: p2p.Response,
		Report: &proto.ReportRequest{
			MessageIDs: messageIDs,
			State:      proto.MessageState_FAILED,
			Comment:    err.Error(),
		},
	}
	j, err := json.Marshal(&d)
	if err != nil {
		slog.Error("json marshal p2p response data failed", "error", err, "messageIDs", messageIDs)
		return
	}
	if err := r.topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish response data to p2p network failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) reportSuccess(messageIDs []string, state proto.MessageState, comment string) {
	d := p2p.Data{
		Type: p2p.Response,
		Report: &proto.ReportRequest{
			MessageIDs: messageIDs,
			State:      state,
			Comment:    comment,
		},
	}
	j, err := json.Marshal(&d)
	if err != nil {
		slog.Error("json marshal p2p response data failed", "error", err, "messageIDs", messageIDs)
		return
	}
	if err := r.topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish response data to p2p network failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) getMessageIDs(ms []*proto.Message) []string {
	ids := []string{}
	for _, m := range ms {
		ids = append(ids, m.MessageID)
	}
	return ids
}
