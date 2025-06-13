package storage

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/MukizuL/hezzl-test/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Repo interface {
	CreateGoods(ctx context.Context, projectId int, name string) (int, error)
	GetGood(ctx context.Context, id int) (*models.Goods, error)
	UpdateGood(ctx context.Context, id, projectId int, name, description string) error
	RemoveGoods(ctx context.Context, id, projectId int) error
}

type Storage struct {
	pg     *pgxpool.Pool
	ch     driver.Conn
	logger *zap.Logger
}

func newStorage(pg *pgxpool.Pool, ch driver.Conn, logger *zap.Logger) (Repo, error) {
	return &Storage{
		pg:     pg,
		ch:     ch,
		logger: logger,
	}, nil
}

func Provide() fx.Option {
	return fx.Provide(newStorage)
}
