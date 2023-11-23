package message

import (
	"context"
	"log/slog"
	"time"

	"github.com/machinefi/sprout/output/chain/eth"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/tasks"
	"github.com/machinefi/sprout/test/contract"
	"github.com/machinefi/sprout/util/mq"
	"github.com/machinefi/sprout/util/mq/gochan"
	"github.com/machinefi/sprout/vm"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler struct {
	mq                 mq.MQ
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

	q := gochan.New()
	h := &Handler{
		mq:                 q,
		vmProcessor:        vmProcessor,
		chainEndpoint:      chainEndpoint,
		operatorPrivateKey: operatorPrivateKey,
		projectManager:     projectManager,
		sequencerClient:    proto.NewSequencerClient(conn),
		projectID:          projectID,
	}
	go q.Watch(h.asyncHandle)
	go h.fetchMessage()
	return h, nil
}

// TODO support batch message fetch & fetch frequency define
func (r *Handler) fetchMessage() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		m, err := r.sequencerClient.Fetch(context.Background(), &proto.FetchRequest{ProjectID: r.projectID})
		if err != nil {
			slog.Error("fetch message from sequencer failed", "error", err)
		}
		if len(m.Messages) != 0 {
			r.mq.Enqueue(m.Messages[0])
		}
	}
}

func (r *Handler) Handle(m *proto.Message) error {
	slog.Debug("push message into sequencer")
	tasks.New(m)
	return r.mq.Enqueue(m)
}

func (r *Handler) asyncHandle(m *proto.Message) {
	slog.Debug("message popped", "message_id", m.MessageID)

	project, err := r.projectManager.Get(m.ProjectID)
	if err != nil {
		slog.Error("get project failed:", "error", err)
		tasks.OnFailed(m.MessageID, err)
		return
	}

	tasks.OnSubmitProving(m.MessageID)
	res, err := r.vmProcessor.Process(m, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed:", "error", err)
		tasks.OnFailed(m.MessageID, err)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	tasks.OnProved(m.MessageID, string(res))

	if r.operatorPrivateKey == "" {
		info := "missing operator private key, will not write to chain"
		slog.Debug(info)
		tasks.OnSucceeded(m.MessageID, info)
		return
	}

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		tasks.OnFailed(m.MessageID, err)
		return
	}

	slog.Debug("writing proof to chain")

	tasks.OnSubmitToBlockchain(m.MessageID)
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x6e30b42554DDA34bAFca9cB00Ec4B464f452a671", data)
	if err != nil {
		slog.Error(err.Error())
		tasks.OnFailed(m.MessageID, err)
		return
	}
	tasks.OnSucceeded(m.MessageID, txHash)
	slog.Debug("transaction hash", "tx_hash", txHash)
}
