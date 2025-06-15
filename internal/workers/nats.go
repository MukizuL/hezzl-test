package workers

import (
	"context"
	"encoding/json"
	"github.com/MukizuL/hezzl-test/internal/zlog"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"sync"
)

type NATSConsumer struct {
	nc          *nats.Conn
	writer      *ClickHouseWriter
	logger      *zap.Logger
	sub         *nats.Subscription
	shutdownCtx context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
}

func newNATSConsumer(lc fx.Lifecycle, nc *nats.Conn, writer *ClickHouseWriter, logger *zap.Logger) (*NATSConsumer, error) {
	ctx, cancel := context.WithCancel(context.Background())

	consumer := &NATSConsumer{
		nc:          nc,
		writer:      writer,
		logger:      logger,
		shutdownCtx: ctx,
		cancel:      cancel,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return consumer.Start("logs")
		},
		OnStop: func(ctx context.Context) error {
			return consumer.Stop()
		},
	})

	return consumer, nil
}

func (c *NATSConsumer) Start(subject string) error {
	sub, err := c.nc.Subscribe(subject, func(msg *nats.Msg) {
		c.wg.Add(1)
		defer c.wg.Done()

		var logData zlog.LogData
		if err := json.Unmarshal(msg.Data, &logData); err != nil {
			c.logger.Error("Failed to unmarshal message", zap.Error(err))
			return
		}

		if err := c.writer.Write(logData); err != nil {
			c.logger.Error("Failed to write to ClickHouse worker", zap.Error(err))
		}
	})
	if err != nil {
		return err
	}

	c.sub = sub
	return nil
}

func (c *NATSConsumer) Stop() error {
	c.cancel()
	if err := c.sub.Unsubscribe(); err != nil {
		return err
	}
	c.wg.Wait()
	return c.writer.Close()
}
