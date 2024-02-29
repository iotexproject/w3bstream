package project

import (
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/machinefi/sprout/project/contracts"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewManager(t *testing.T) {
	PatchConvey("NewManagerDialChainFailed", t, func() {
		Mock(ethclient.Dial).Return(nil, errors.New(t.Name())).Build()

		_, err := NewManager("", "", "")
		So(err.Error(), ShouldContainSubstring, t.Name())
	})
	PatchConvey("NewManagerNewContractsFailed", t, func() {
		Mock(ethclient.Dial).Return(nil, nil).Build()
		Mock(contracts.NewContracts).Return(nil, errors.New(t.Name())).Build()

		_, err := NewManager("", "", "")
		So(err.Error(), ShouldContainSubstring, t.Name())
	})
	PatchConvey("NewManagerNewDefaultMonitorFailed", t, func() {
		Mock(ethclient.Dial).Return(ethclient.NewClient(&rpc.Client{}), nil).Build()
		Mock(contracts.NewContracts).Return(nil, nil).Build()
		Mock((*Manager).fillProjectPool).Return().Build()
		Mock(NewDefaultMonitor).Return(nil, errors.New(t.Name())).Build()

		_, err := NewManager("", "", "")
		So(err.Error(), ShouldContainSubstring, t.Name())
	})
	PatchConvey("NewManagerSuccess", t, func() {
		Mock(ethclient.Dial).Return(ethclient.NewClient(&rpc.Client{}), nil).Build()
		Mock(contracts.NewContracts).Return(nil, nil).Build()
		Mock((*Manager).fillProjectPool).Return().Build()
		Mock(NewDefaultMonitor).Return(&Monitor{}, nil).Build()
		Mock((*Monitor).run).Return().Build()
		Mock((*Monitor).MustEvents).Return(make(chan *types.Log)).Build()
		Mock((*Manager).watchProjectRegistrar).Return().Build()

		_, err := NewManager("", "", "")
		So(err, ShouldBeEmpty)
	})
}

type testSubscription struct {
	errChain chan error
}

func (s testSubscription) Err() <-chan error {
	return s.errChain
}

func (s testSubscription) Unsubscribe() {}

func TestManagerMethod(t *testing.T) {
	Convey("GetNotExist", t, func() {
		m := &Manager{}
		_, err := m.Get(1, "0.1")
		So(err.Error(), ShouldContainSubstring, "project config not exist")
	})
	Convey("GetSuccess", t, func() {
		m := &Manager{
			pool: map[key]*Config{getKey(1, "0.1"): {}},
		}
		_, err := m.Get(1, "0.1")
		So(err, ShouldBeEmpty)
	})
	Convey("SetSuccess", t, func() {
		m := &Manager{
			pool:       map[key]*Config{},
			projectIDs: map[uint64]bool{},
		}
		m.Set(1, "0.1", &Config{})
	})
	Convey("GetAllSuccess", t, func() {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
		}
		ids := m.GetAllProjectID()
		So(len(ids), ShouldEqual, 1)
		So(ids[0], ShouldEqual, uint64(1))
	})
	Convey("GetNotifySuccess", t, func() {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
			notify:     make(chan uint64, 1),
		}
		notify := m.GetNotify()
		m.notify <- uint64(1)
		d := <-notify
		So(d, ShouldEqual, uint64(1))
	})
	PatchConvey("DoProjectRegistrarWatchSuccess", t, func() {
		Mock((*contracts.Contracts).ParseProjectUpserted).Return(&contracts.ContractsProjectUpserted{ProjectId: 1}, nil).Build()
		Mock((*ProjectMeta).GetConfigs).Return([]*Config{{}}, nil).Build()

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			notify:     make(chan uint64, 1),
			instance:   &contracts.Contracts{},
		}

		errChain := make(chan error)
		logChain := make(chan *types.Log, 1)
		logChain <- &types.Log{}

		m.doProjectRegistrarWatch(logChain, testSubscription{errChain})
		notify := m.GetNotify()
		m.notify <- uint64(1)
		d := <-notify
		So(d, ShouldEqual, uint64(1))
	})
	PatchConvey("DoProjectPoolFillReadChainFailed", t, func() {
		Mock((*contracts.ContractsCaller).Projects).Return(nil, errors.New(t.Name())).Build()

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		So(finished, ShouldBeFalse)
		So(len(m.GetAllProjectID()), ShouldEqual, 0)
	})
	PatchConvey("DoProjectPoolFillFinished", t, func() {
		Mock((*contracts.ContractsCaller).Projects).Return(struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{}, nil).Build()

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		So(finished, ShouldBeTrue)
		So(len(m.GetAllProjectID()), ShouldEqual, 0)
	})
	PatchConvey("DoProjectPoolFillGetConfigFailed", t, func() {
		Mock((*contracts.ContractsCaller).Projects).Return(struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{
			Uri:  "test",
			Hash: [32]byte{1},
		}, nil).Build()
		Mock((*ProjectMeta).GetConfigs).Return(nil, errors.New(t.Name())).Build()

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		So(finished, ShouldBeFalse)
		So(len(m.GetAllProjectID()), ShouldEqual, 0)
	})
	PatchConvey("DoProjectPoolFillSuccess", t, func() {
		Mock((*contracts.ContractsCaller).Projects).Return(struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{
			Uri:  "test",
			Hash: [32]byte{1},
		}, nil).Build()
		Mock((*ProjectMeta).GetConfigs).Return([]*Config{{}}, nil).Build()

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		So(finished, ShouldBeFalse)
		So(len(m.GetAllProjectID()), ShouldEqual, 1)
	})
	PatchConvey("FillProjectPoolSuccess", t, func() {
		Mock((*Manager).doProjectPoolFill).Return(true).Build()

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		m.fillProjectPool()
		So(len(m.GetAllProjectID()), ShouldEqual, 0)
	})
}
