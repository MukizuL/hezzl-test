package services

import (
	"github.com/MukizuL/hezzl-test/internal/storage"
	"github.com/MukizuL/hezzl-test/internal/zlog"
	"go.uber.org/fx"
)

type Services struct {
	storage storage.Repo
	logger  *zlog.Logger
}

func newServices(storage storage.Repo, logger *zlog.Logger) *Services {
	return &Services{
		storage: storage,
		logger:  logger,
	}
}

func Provide() fx.Option {
	return fx.Provide(newServices)
}
