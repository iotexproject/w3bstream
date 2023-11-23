package message

import (
	"context"
	"fmt"
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

type Handler struct {
	vmProcessor        *vm.Processor
	projectManager     *project.Manager
	chainEndpoint      string
	operatorPrivateKey string
	sequencerClient    proto.SequencerClient
	projectID          uint64
}

func NewHandler(vmProcessor *vm.Processor, projectManager *project.Manager, chainEndpoint, sequencerEndpoint, operatorPrivateKey string, projectID uint64) (*Handler, error) {
	conn, err := grpc.Dial(sequencerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial sequencer server")
	}

	h := &Handler{
		vmProcessor:        vmProcessor,
		chainEndpoint:      chainEndpoint,
		operatorPrivateKey: operatorPrivateKey,
		projectManager:     projectManager,
		sequencerClient:    proto.NewSequencerClient(conn),
		projectID:          projectID,
	}
	return h, nil
}

// TODO support batch message fetch & fetch frequency define
func (r *Handler) Run() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		m, err := r.sequencerClient.Fetch(context.Background(), &proto.FetchRequest{ProjectID: r.projectID})
		if err != nil {
			slog.Error("fetch message from sequencer failed", "error", err)
		}
		if len(m.Messages) != 0 {
			r.handle(m.Messages[0])
		}
	}
}

func (r *Handler) handle(m *proto.Message) {
	slog.Debug("message popped", "message_id", m.MessageID)

	project, err := r.projectManager.Get(m.ProjectID)
	if err != nil {
		slog.Error("get project failed", "error", err)
		r.reportFail([]string{m.MessageID}, err)
		return
	}

	r.reportSuccess([]string{m.MessageID}, proto.MessageState_PROVING, "")
	res, err := r.vmProcessor.Process(m, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed", "error", err)
		r.reportFail([]string{m.MessageID}, err)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	r.reportSuccess([]string{m.MessageID}, proto.MessageState_PROVED, string(res))

	if r.operatorPrivateKey == "" {
		info := "missing operator private key, will not write to chain"
		slog.Debug(info)
		r.reportSuccess([]string{m.MessageID}, proto.MessageState_OUTPUTTED, info)
		return
	}

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail([]string{m.MessageID}, err)
		return
	}

	slog.Debug("writing proof to chain")

	r.reportSuccess([]string{m.MessageID}, proto.MessageState_OUTPUTTING, "writing proof to chain")
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x6e30b42554DDA34bAFca9cB00Ec4B464f452a671", data)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail([]string{m.MessageID}, err)
		return
	}
	r.reportSuccess([]string{m.MessageID}, proto.MessageState_OUTPUTTED, fmt.Sprintf("transaction hash is %s", txHash))
	slog.Debug("transaction hash", "tx_hash", txHash)
}

func (r *Handler) reportFail(messageIDs []string, err error) {
	if _, err := r.sequencerClient.Report(context.Background(), &proto.ReportRequest{
		MessageIDs: messageIDs,
		State:      proto.MessageState_FAILED,
		Comment:    err.Error(),
	}); err != nil {
		slog.Error("report to sequencer failed", "error", err, "messageIDs", messageIDs)
	}
}

func (r *Handler) reportSuccess(messageIDs []string, state proto.MessageState, comment string) {
	if _, err := r.sequencerClient.Report(context.Background(), &proto.ReportRequest{
		MessageIDs: messageIDs,
		State:      state,
		Comment:    comment,
	}); err != nil {
		slog.Error("report to sequencer failed", "error", err, "messageIDs", messageIDs)
	}
}
