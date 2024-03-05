package bundler_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/0xsequence/bundler"
	"github.com/0xsequence/bundler/config"
	"github.com/0xsequence/bundler/contracts/gen/solabis/abiendorser"
	"github.com/0xsequence/bundler/endorser"
	"github.com/0xsequence/bundler/mempool"
	"github.com/0xsequence/bundler/mocks"
	"github.com/0xsequence/bundler/proto"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/go-chi/httplog/v2"
	"github.com/stretchr/testify/mock"
)

func TestIddlePull(t *testing.T) {
	mockMempool := &mocks.MockMempool{}

	done := make(chan bool, 2)

	mockMempool.On("ReserveOps", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		done <- true
	}).Return(
		[]*mempool.TrackedOperation{},
	).Maybe()

	pruner := bundler.NewPruner(config.PrunerConfig{
		RunWaitMillis: 1,
	}, mockMempool, nil, nil)
	ctx, cancel := context.WithCancel(context.Background())
	go pruner.Run(ctx)

	<-done
	<-done

	mockMempool.AssertCalled(t, "ReserveOps", mock.Anything, mock.Anything)
	mockMempool.AssertNumberOfCalls(t, "ReserveOps", 2)

	cancel()
}

func TestPullAndDiscartStateErr(t *testing.T) {
	mockMempool := &mocks.MockMempool{}
	mockEndorser := &mocks.MockEndorser{}
	logger := httplog.NewLogger("")

	op1 := &mempool.TrackedOperation{}

	done := make(chan bool)

	mockMempool.On("ReserveOps", mock.Anything, mock.Anything).Return(
		[]*mempool.TrackedOperation{op1},
	).Maybe()

	mockMempool.On("DiscardOps", mock.Anything, mock.Anything).Run(func(mock.Arguments) {
		done <- true
	}).Return().Once()

	mockEndorser.On("DependencyState", mock.Anything, op1.EndorserResult).Return(
		nil, fmt.Errorf("error"),
	).Once()

	pruner := bundler.NewPruner(config.PrunerConfig{
		RunWaitMillis: 1,
	}, mockMempool, mockEndorser, logger)

	ctx, cancel := context.WithCancel(context.Background())
	go pruner.Run(ctx)

	<-done
	cancel()
}

func TestPullAndDiscartHasChangedErr(t *testing.T) {
	mockMempool := &mocks.MockMempool{}
	mockEndorser := &mocks.MockEndorser{}
	logger := httplog.NewLogger("")

	er1 := &endorser.EndorserResult{
		Dependencies: []abiendorser.EndorserDependency{},
	}
	er2 := &endorser.EndorserResultState{
		AddrDependencies: make(map[common.Address]*endorser.AddrDependencyState),
	}
	op1 := &mempool.TrackedOperation{
		EndorserResult:      er1,
		EndorserResultState: er2,
	}

	done := make(chan bool)

	mockMempool.On("ReserveOps", mock.Anything, mock.Anything).Return(
		[]*mempool.TrackedOperation{op1},
	).Maybe()

	mockMempool.On("DiscardOps", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(1).([]string)
		if arg[0] == op1.Hash() {
			done <- true
		}
	}).Return().Once()

	mockEndorser.On("DependencyState", mock.Anything, op1.EndorserResult).Return(
		nil, fmt.Errorf("error"),
	).Once()

	pruner := bundler.NewPruner(config.PrunerConfig{
		RunWaitMillis: 1,
	}, mockMempool, mockEndorser, logger)

	ctx, cancel := context.WithCancel(context.Background())
	go pruner.Run(ctx)

	<-done
	cancel()
}

func TestPullAndReleaseNotChanged(t *testing.T) {
	mockMempool := &mocks.MockMempool{}
	mockEndorser := &mocks.MockEndorser{}
	logger := httplog.NewLogger("")

	er1 := &endorser.EndorserResult{
		Dependencies: []abiendorser.EndorserDependency{},
	}
	er2 := &endorser.EndorserResultState{
		AddrDependencies: make(map[common.Address]*endorser.AddrDependencyState),
	}
	op1 := &mempool.TrackedOperation{
		EndorserResult:      er1,
		EndorserResultState: er2,
	}

	done := make(chan bool)

	mockMempool.On("ReserveOps", mock.Anything, mock.Anything).Return(
		[]*mempool.TrackedOperation{op1},
	).Maybe()

	mockMempool.On(
		"ReleaseOps",
		mock.Anything,
		mock.Anything,
		proto.ReadyAtChange_Now,
	).Run(func(args mock.Arguments) {
		arg := args.Get(1).([]string)
		if arg[0] == op1.Hash() {
			done <- true
		}
	}).Return().Once()

	mockEndorser.On("DependencyState", mock.Anything, op1.EndorserResult).Return(
		er2, nil,
	).Once()

	pruner := bundler.NewPruner(config.PrunerConfig{
		RunWaitMillis: 1,
	}, mockMempool, mockEndorser, logger)

	ctx, cancel := context.WithCancel(context.Background())
	go pruner.Run(ctx)

	<-done
	cancel()
}

func TestDiscardNotReady(t *testing.T) {
	mockMempool := &mocks.MockMempool{}
	mockEndorser := &mocks.MockEndorser{}
	logger := httplog.NewLogger("")

	da := common.HexToAddress("0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E")
	er1 := &endorser.EndorserResult{
		Dependencies: []abiendorser.EndorserDependency{
			{
				Addr:    da,
				Balance: true,
			},
		},
	}
	er2m := make(map[common.Address]*endorser.AddrDependencyState)
	er2m[da] = &endorser.AddrDependencyState{
		Balance: new(big.Int).SetUint64(100),
	}

	er2 := &endorser.EndorserResultState{
		AddrDependencies: er2m,
	}

	er3m := make(map[common.Address]*endorser.AddrDependencyState)
	er3m[da] = &endorser.AddrDependencyState{
		Balance: new(big.Int).SetUint64(200),
	}

	er3 := &endorser.EndorserResultState{
		AddrDependencies: er3m,
	}

	op1 := &mempool.TrackedOperation{
		EndorserResult:      er1,
		EndorserResultState: er2,
	}

	done := make(chan bool)

	mockMempool.On("ReserveOps", mock.Anything, mock.Anything).Return(
		[]*mempool.TrackedOperation{op1},
	).Maybe()

	mockEndorser.On("DependencyState", mock.Anything, op1.EndorserResult).Return(
		er3, nil,
	).Once()

	mockEndorser.On("IsOperationReady", mock.Anything, &op1.Operation).Return(
		&endorser.EndorserResult{
			Readiness:        false,
			Dependencies:     []abiendorser.EndorserDependency{},
			GlobalDependency: abiendorser.EndorserGlobalDependency{},
		}, nil,
	).Once()

	mockMempool.On("DiscardOps", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(1).([]string)
		if arg[0] == op1.Hash() {
			done <- true
		}
	}).Return().Once()

	pruner := bundler.NewPruner(config.PrunerConfig{
		RunWaitMillis: 1,
	}, mockMempool, mockEndorser, logger)

	ctx, cancel := context.WithCancel(context.Background())
	go pruner.Run(ctx)

	<-done
	cancel()
}

func TestKeepReady(t *testing.T) {
	mockMempool := &mocks.MockMempool{}
	mockEndorser := &mocks.MockEndorser{}
	logger := httplog.NewLogger("")

	da := common.HexToAddress("0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E")
	er1 := &endorser.EndorserResult{
		Dependencies: []abiendorser.EndorserDependency{
			{
				Addr:    da,
				Balance: true,
			},
		},
	}
	er2m := make(map[common.Address]*endorser.AddrDependencyState)
	er2m[da] = &endorser.AddrDependencyState{
		Balance: new(big.Int).SetUint64(100),
	}

	er2 := &endorser.EndorserResultState{
		AddrDependencies: er2m,
	}

	er3m := make(map[common.Address]*endorser.AddrDependencyState)
	er3m[da] = &endorser.AddrDependencyState{
		Balance: new(big.Int).SetUint64(200),
	}

	er3 := &endorser.EndorserResultState{
		AddrDependencies: er3m,
	}

	op1 := &mempool.TrackedOperation{
		EndorserResult:      er1,
		EndorserResultState: er2,
	}

	done := make(chan bool)

	mockMempool.On("ReserveOps", mock.Anything, mock.Anything).Return(
		[]*mempool.TrackedOperation{op1},
	).Maybe()

	mockEndorser.On("DependencyState", mock.Anything, op1.EndorserResult).Return(
		er3, nil,
	).Once()

	mockEndorser.On("IsOperationReady", mock.Anything, &op1.Operation).Return(
		&endorser.EndorserResult{
			Readiness:        true,
			Dependencies:     []abiendorser.EndorserDependency{},
			GlobalDependency: abiendorser.EndorserGlobalDependency{},
		}, nil,
	).Once()

	mockMempool.On(
		"ReleaseOps",
		mock.Anything,
		mock.Anything,
		proto.ReadyAtChange_Now,
	).Run(func(args mock.Arguments) {
		arg := args.Get(1).([]string)
		if arg[0] == op1.Hash() {
			done <- true
		}
	}).Return().Once()

	pruner := bundler.NewPruner(config.PrunerConfig{
		RunWaitMillis: 1,
	}, mockMempool, mockEndorser, logger)

	ctx, cancel := context.WithCancel(context.Background())
	go pruner.Run(ctx)

	<-done
	cancel()
}

func TestSkipRecentOps(t *testing.T) {
	mockMempool := &mocks.MockMempool{}

	done := make(chan bool)
	pruner := bundler.NewPruner(config.PrunerConfig{
		RunWaitMillis: 1,
	}, mockMempool, nil, nil)

	mockMempool.On("ReserveOps", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		clb := args.Get(1).(func([]*mempool.TrackedOperation) []*mempool.TrackedOperation)
		res := clb([]*mempool.TrackedOperation{
			{
				ReadyAt: time.Now().Add(-(pruner.GracePeriod / 2)),
			},
		})
		if len(res) == 0 {
			done <- true
		}
	}).Return(
		[]*mempool.TrackedOperation{},
	).Maybe()

	ctx, cancel := context.WithCancel(context.Background())
	go pruner.Run(ctx)

	<-done
	cancel()
}
