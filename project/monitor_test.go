package project

import (
	"testing"
	"time"

	. "github.com/bytedance/mockey"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMonitor(t *testing.T) {
	PatchConvey("NewMonitorDialChainFailed", t, func() {
		Mock(ethclient.Dial).Return(nil, errors.New(t.Name())).Build()

		_, err := NewMonitor("", []string{}, []string{}, 1, 100, 3*time.Second)
		So(err.Error(), ShouldContainSubstring, t.Name())
	})
	PatchConvey("NewManagerSuccess", t, func() {
		Mock(ethclient.Dial).Return(nil, nil).Build()

		_, err := NewMonitor("", []string{"0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26"}, []string{"ProjectUpserted(uint64,string,bytes32)"}, 1, 100, 3*time.Second)
		So(err, ShouldBeEmpty)
	})
	PatchConvey("NewDefaultMonitorDialChainFailed", t, func() {
		Mock(ethclient.Dial).Return(nil, errors.New(t.Name())).Build()

		_, err := NewDefaultMonitor("", []string{}, []string{})
		So(err.Error(), ShouldContainSubstring, t.Name())
	})
	PatchConvey("NewDefaultMonitorGetBlockNumberFailed", t, func() {
		Mock(ethclient.Dial).Return(ethclient.NewClient(nil), nil).Build()
		Mock((*ethclient.Client).Close).Return().Build()
		Mock((*ethclient.Client).BlockNumber).Return(uint64(0), errors.New(t.Name())).Build()

		_, err := NewDefaultMonitor("", []string{}, []string{})
		So(err.Error(), ShouldContainSubstring, t.Name())
	})
	PatchConvey("NewDefaultMonitorSuccess", t, func() {
		Mock(ethclient.Dial).Return(ethclient.NewClient(nil), nil).Build()
		Mock((*ethclient.Client).Close).Return().Build()
		Mock((*ethclient.Client).BlockNumber).Return(uint64(100), nil).Build()
		Mock(NewMonitor).Return(&Monitor{}, nil).Build()

		_, err := NewDefaultMonitor("", []string{}, []string{})
		So(err, ShouldBeEmpty)
	})
}

func TestMonitorMethod(t *testing.T) {
	Convey("Err", t, func() {
		m := &Monitor{
			err: make(chan error, 1),
		}
		res := m.Err()
		err := errors.New(t.Name())
		m.err <- err
		So(<-res, ShouldEqual, err)
	})
	Convey("Unsubscribe", t, func() {
		m := &Monitor{
			stop: make(chan struct{}, 1),
		}
		m.Unsubscribe()
		So(<-m.stop, ShouldNotBeEmpty)
	})
	Convey("Events", t, func() {
		topic := "test"
		m := &Monitor{
			events: map[common.Hash]chan *types.Log{crypto.Keccak256Hash([]byte(topic)): make(chan *types.Log)},
		}
		_, ok := m.Events(topic)
		So(ok, ShouldBeTrue)
	})
	Convey("MustEvents", t, func() {
		topic := "test"
		m := &Monitor{
			events: map[common.Hash]chan *types.Log{crypto.Keccak256Hash([]byte(topic)): make(chan *types.Log)},
		}
		ch := m.MustEvents(topic)
		So(ch, ShouldNotBeNil)
	})
	Convey("DoRunStopped", t, func() {
		m := &Monitor{
			stop: make(chan struct{}, 1),
		}
		m.Unsubscribe()
		finished := m.doRun()
		So(finished, ShouldBeTrue)
	})
	PatchConvey("DoRunBlockNumberFailed", t, func() {
		Mock((*ethclient.Client).BlockNumber).Return(uint64(100), errors.New(t.Name())).Build()
		Mock(time.Sleep).Return().Build()

		m := &Monitor{
			stop: make(chan struct{}, 1),
		}
		finished := m.doRun()
		So(finished, ShouldBeFalse)
	})
	PatchConvey("DoRunBlockNumberBehind", t, func() {
		Mock((*ethclient.Client).BlockNumber).Return(uint64(100), nil).Build()
		Mock(time.Sleep).Return().Build()

		m := &Monitor{
			latest: 1000,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		So(finished, ShouldBeFalse)
	})
	PatchConvey("DoRunFilterLogsFailed", t, func() {
		Mock((*ethclient.Client).BlockNumber).Return(uint64(100), nil).Build()
		Mock((*ethclient.Client).FilterLogs).Return(nil, errors.New(t.Name())).Build()
		Mock(time.Sleep).Return().Build()

		m := &Monitor{
			latest: 1,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		So(finished, ShouldBeFalse)
	})
	PatchConvey("DoRunFilterLogsEmpty", t, func() {
		Mock((*ethclient.Client).BlockNumber).Return(uint64(100), nil).Build()
		Mock((*ethclient.Client).FilterLogs).Return(nil, nil).Build()
		Mock(time.Sleep).Return().Build()

		m := &Monitor{
			latest: 1,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		So(finished, ShouldBeFalse)
	})
	PatchConvey("DoRunFilterLogsEmpty", t, func() {
		Mock((*ethclient.Client).BlockNumber).Return(uint64(100), nil).Build()
		Mock((*ethclient.Client).FilterLogs).Return([]types.Log{{Topics: []common.Hash{crypto.Keccak256Hash([]byte("0"))}}}, nil).Build()
		Mock(time.Sleep).Return().Build()

		m := &Monitor{
			latest: 1,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		So(finished, ShouldBeFalse)
	})
	PatchConvey("Run", t, func() {
		moker := Mock((*Monitor).doRun).Return(true).Build()

		m := &Monitor{}
		m.run()
		So(moker.Times(), ShouldEqual, 1)
	})
}
