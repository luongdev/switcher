package internal

import (
	"github.com/luongdev/switcher/workflow/types"
	"github.com/uber-go/tally"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"io"
	"os"
	"time"
)

type WorkerImpl struct {
	registry types.Registry
	client   types.Client
	worker   worker.Worker
	closer   io.Closer
	logger   *zap.Logger
}

func (w *WorkerImpl) Start() {
}

func NewWorker(domain, taskList string, client types.Client, registry types.Registry) (types.Worker, error) {
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}

	scope, closer := tally.NewRootScope(tally.ScopeOptions{
		Prefix:   client.GetName(),
		Tags:     map[string]string{"env": os.Getenv("ENV")},
		Reporter: tally.NullStatsReporter,
	}, 5*time.Second)

	workerOptions := worker.Options{Logger: logger, MetricsScope: scope}
	w := worker.New(client, domain, taskList, workerOptions)

	return &WorkerImpl{client: client, registry: registry, worker: w, closer: closer, logger: logger}, nil
}

func (w *WorkerImpl) Stop() {
	w.worker.Stop()
	err := w.closer.Close()
	if err != nil {
		w.logger.Error("failed to close metrics scope", zap.Error(err))
		return
	}
}

var _ types.Worker = (*WorkerImpl)(nil)
