package project

import (
	"testing"
	"time"

	. "github.com/bytedance/mockey"
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
