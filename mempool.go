package bundler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/0xsequence/bundler/config"
	"github.com/0xsequence/bundler/endorser"
	"github.com/0xsequence/bundler/proto"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/go-chi/httplog/v2"
)

type TrackedOperation struct {
	proto.Operation

	ReservedSince *time.Time
	CreatedAt     time.Time
	ReadyAt       time.Time

	EndorserResult *endorser.EndorserResult
}

type Mempool struct {
	logger *httplog.Logger

	Provider *ethrpc.Provider
	MaxSize  uint

	flock           sync.Mutex
	FreshOperations *[]*proto.Operation

	olock      sync.Mutex
	Operations []*TrackedOperation

	digests map[common.Hash]struct{}
}

func NewMempool(cfg *config.MempoolConfig, logger *httplog.Logger, provider *ethrpc.Provider) (*Mempool, error) {
	mp := &Mempool{
		logger:   logger,
		Provider: provider,
		MaxSize:  cfg.Size,

		flock: sync.Mutex{},
		olock: sync.Mutex{},

		FreshOperations: &[]*proto.Operation{},
		Operations:      []*TrackedOperation{},

		digests: map[common.Hash]struct{}{},
	}

	return mp, nil
}

func (mp *Mempool) Size() int {
	return len(mp.Operations) + len(*mp.FreshOperations)
}

func (mp *Mempool) AddOperation(op *proto.Operation) error {
	if op == nil {
		return fmt.Errorf("mempool: operation is nil")
	}

	digest := op.Digest()
	if _, ok := mp.digests[digest]; ok {
		return fmt.Errorf("mempool: duplicate operation")
	}
	mp.digests[digest] = struct{}{}

	mp.flock.Lock()
	defer mp.flock.Unlock()

	if mp.Size() >= int(mp.MaxSize) {
		return fmt.Errorf("mempool: max size reached")
	}

	mp.logger.Info("mempool: adding operation to fresh", "op", op)

	nlist := append(*mp.FreshOperations, op)
	mp.FreshOperations = &nlist

	return nil
}

func (mp *Mempool) ReserveOps(ctx context.Context, selec func([]*TrackedOperation) []*TrackedOperation) {
	mp.olock.Lock()
	defer mp.olock.Unlock()

	// Filter out the operations that are already reserved
	// and the ones that are not ready
	avalOps := []*TrackedOperation{}
	for _, op := range mp.Operations {
		if op.ReservedSince != nil {
			continue
		}
		avalOps = append(avalOps, op)
	}

	// Select the operations to reserve
	resOps := selec(avalOps)
	for _, op := range resOps {
		n := time.Now()
		op.ReservedSince = &n
	}
}

func (mp *Mempool) DiscardOps(ctx context.Context, ops []*TrackedOperation) {
	mp.olock.Lock()
	defer mp.olock.Unlock()

	var kops []*TrackedOperation
	for _, op := range mp.Operations {
		discard := false

		for _, dop := range ops {
			if op.Digest() == dop.Digest() {
				discard = true
				break
			}
		}

		if discard {
			continue
		}

		kops = append(kops, op)
	}

	// Remove them from the digest map too
	delete(mp.digests, ops[0].Digest())

	mp.Operations = kops
}

func (mp *Mempool) StartProcessor(ctx context.Context) {
	// Run every 500 ms
	go func() {
		ticker := time.NewTicker(time.Duration(500 * time.Millisecond))
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := mp.HandleFreshOps(ctx)
				if err != nil {
					mp.logger.Error("mempool: error handling fresh operations", "err", err)
				}
			case <-ctx.Done():
				// Context cancelled, stop the ticker
				return
			}
		}
	}()
}

func (mp *Mempool) HandleFreshOps(ctx context.Context) error {
	// Take each operation from the fresh mempool, pass them through the
	// endorser. If they are not ready, then we drop them.

	// TODO: Parallelize this

	if len(*mp.FreshOperations) == 0 {
		return nil
	}

	// Create a local copy of the fresh operations
	// and clear the fresh operations list. This is going
	// to take a while and we don't want to block new operations.

	// NOTICE that the mempool could grow over the limit while we are
	// processing the fresh operations. This is fine for now.

	mp.flock.Lock()
	freshOps := mp.FreshOperations
	mp.FreshOperations = &[]*proto.Operation{}
	mp.flock.Unlock()

	for _, op := range *freshOps {
		res, err := endorser.IsOperationReady(ctx, mp.Provider, op)
		if err != nil {
			mp.logger.Warn("dropping operation", "op", op, "reason", "endorser error", "err", err)
			continue
		}

		if !res.Readiness {
			mp.logger.Debug("dropping operation", "op", op, "reason", "not ready")
			continue
		}

		// Check the constraints
		okc, err := res.CheckConstraints(ctx, mp.Provider)
		if err != nil {
			mp.logger.Warn("dropping operation", "op", op, "reason", "constraint error", "err", err)
			continue
		}

		if !okc {
			mp.logger.Debug("dropping operation", "op", op, "reason", "constraint not met")
			continue
		}

		// If the operation is ready
		// then we add it to the mempool

		mp.olock.Lock()
		mp.logger.Info("operation added to mempool", "op", op)
		mp.Operations = append(mp.Operations, &TrackedOperation{
			Operation: *op,

			CreatedAt: time.Now(),
			ReadyAt:   time.Now(),

			EndorserResult: res,
		})
		mp.olock.Unlock()
	}

	return nil
}
