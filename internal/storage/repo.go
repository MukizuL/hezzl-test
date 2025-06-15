package storage

import (
	"context"
	"github.com/MukizuL/hezzl-test/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Repo interface {
	CreateGoods(ctx context.Context, projectId int, name string) (int, error)
	GetGood(ctx context.Context, id int) (*models.Goods, error)
	GetGoodsSortPriority(ctx context.Context) ([]models.Goods, error)
	GetGoodsSortId(ctx context.Context) ([]models.Goods, error)
	GetGoodsWithLimit(ctx context.Context, limit, offset int) ([]models.Goods, error)
	UpdateGood(ctx context.Context, id, projectId int, name, description string) error
	RemoveGoods(ctx context.Context, id, projectId int) error
	Get(ctx context.Context, limit, offset int) ([]models.Goods, error)
	Set(ctx context.Context) error
	Invalidate(ctx context.Context)
}

type Storage struct {
	pg *pgxpool.Pool

	redis  *redis.Client
	logger *zap.Logger
}

func newStorage(pg *pgxpool.Pool, redis *redis.Client, logger *zap.Logger) (Repo, error) {
	return &Storage{
		pg:     pg,
		redis:  redis,
		logger: logger,
	}, nil
}

func Provide() fx.Option {
	return fx.Provide(newStorage)
}
