package workers

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/MukizuL/hezzl-test/internal/zlog"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

const BatchSize = 10
const FlushInterval = 5 * time.Second

type ClickHouseWriter struct {
	ctx         context.Context
	ch          driver.Conn
	batchSize   int
	batchBuffer []zlog.LogData
	flushTicker *time.Ticker
	logger      *zap.Logger
}

func newClickHouseWriter(lc fx.Lifecycle, ch driver.Conn, logger *zap.Logger) (*ClickHouseWriter, error) {
	w := &ClickHouseWriter{
		ctx:         context.Background(),
		ch:          ch,
		batchSize:   BatchSize,
		batchBuffer: make([]zlog.LogData, 0, BatchSize),
		flushTicker: time.NewTicker(FlushInterval),
		logger:      logger,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go w.periodicFlush()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return w.Close()
		},
	})

	return w, nil
}

func (w *ClickHouseWriter) Write(data zlog.LogData) error {
	w.batchBuffer = append(w.batchBuffer, data)

	if len(w.batchBuffer) >= w.batchSize {
		return w.Flush()
	}
	return nil
}

func (w *ClickHouseWriter) Flush() error {
	if len(w.batchBuffer) == 0 {
		return nil
	}

	batch, err := w.ch.PrepareBatch(w.ctx, `INSERT INTO logs`)
	if err != nil {
		return err
	}
	defer batch.Close()

	for _, log := range w.batchBuffer {
		err = batch.Append(
			log.ID,
			log.ProjectID,
			log.Name,
			log.Description,
			log.Priority,
			log.Removed,
			log.Timestamp,
		)
		if err != nil {
			return err
		}
	}

	w.batchBuffer = w.batchBuffer[:0]
	return batch.Send()
}

func (w *ClickHouseWriter) periodicFlush() {
	for range w.flushTicker.C {
		if err := w.Flush(); err != nil {
			// Log error or handle it appropriately
		}
	}
}

func (w *ClickHouseWriter) Close() error {
	w.flushTicker.Stop()
	w.ctx.Done()
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
