package message

import (
	"context"
	"log/slog"
	"time"

	"github.com/machinefi/sprout/output/chain/eth"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/test/contract"
	"github.com/machinefi/sprout/vm"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Processor struct {
	vmHandler          *vm.Handler
	projectManager     *project.Manager
	chainEndpoint      string
	operatorPrivateKey string
	sequencerClient    proto.SequencerClient
}

func NewProcessor(vmHandler *vm.Handler, projectManager *project.Manager, chainEndpoint, sequencerEndpoint, operatorPrivateKey string) (*Processor, error) {
	conn, err := grpc.Dial(sequencerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial sequencer server")
	}

	h := &Processor{
		vmHandler:          vmHandler,
		chainEndpoint:      chainEndpoint,
		operatorPrivateKey: operatorPrivateKey,
		projectManager:     projectManager,
		sequencerClient:    proto.NewSequencerClient(conn),
	}
	return h, nil
}

// TODO support batch message fetch & fetch frequency define
func (r *Processor) Run() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	projectIDs := r.projectManager.GetAllProjectID()

	for range ticker.C {
		for _, projectID := range projectIDs {
			m, err := r.sequencerClient.Fetch(context.Background(), &proto.FetchRequest{ProjectID: projectID})
			if err != nil {
				slog.Error("fetch message from sequencer failed", "error", err)
			}
			if len(m.Messages) != 0 {
				r.process(m.Messages)
			}
		}
	}
}

func (r *Processor) process(ms []*proto.Message) {
	mids := r.getMessageIDs(ms)
	slog.Debug("message popped", "message_ids", mids)
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
	if _, err := r.sequencerClient.Report(context.Background(), &proto.ReportRequest{
		MessageIDs: messageIDs,
		State:      proto.MessageState_FAILED,
		Comment:    err.Error(),
	}); err != nil {
		slog.Error("report to sequencer failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) reportSuccess(messageIDs []string, state proto.MessageState, comment string) {
	if _, err := r.sequencerClient.Report(context.Background(), &proto.ReportRequest{
		MessageIDs: messageIDs,
		State:      state,
		Comment:    comment,
	}); err != nil {
		slog.Error("report to sequencer failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Processor) getMessageIDs(ms []*proto.Message) []string {
	ids := []string{}
	for _, m := range ms {
		ids = append(ids, m.MessageID)
	}
	return ids
}
