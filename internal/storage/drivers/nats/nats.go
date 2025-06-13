package nats

import (
	"context"
	"github.com/MukizuL/hezzl-test/internal/config"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
)

func newNatsConnection(lc fx.Lifecycle, cfg *config.Config) (*nats.Conn, error) {
	nc, err := nats.Connect(cfg.NatsAddr)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := nc.LastError(); err != nil {
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			nc.Close()
			return nil
		},
	})

	return nc, nil
}

func Provide() fx.Option {
	return fx.Provide(newNatsConnection)
}
