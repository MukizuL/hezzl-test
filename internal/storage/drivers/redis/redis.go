package redis

import (
	"context"
	"github.com/MukizuL/hezzl-test/internal/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func newRedisConnection(lc fx.Lifecycle, cfg *config.Config) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisDsn,
		Password: "",
		DB:       0,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return conn.Ping(ctx).Err()
		},
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	return conn, nil
}

func Provide() fx.Option {
	return fx.Provide(newRedisConnection)
}
