package pg

import (
	"context"
	"fmt"
	"github.com/MukizuL/hezzl-test/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func newPostgreSQLConnection(lc fx.Lifecycle, cfg *config.Config) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.PgDsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create db pool: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return dbpool.Ping(ctx)
		},
		OnStop: func(ctx context.Context) error {
			dbpool.Close()
			return nil
		},
	})

	return dbpool, nil
}

func Provide() fx.Option {
	return fx.Provide(newPostgreSQLConnection)
}
