package controller

import (
	"github.com/MukizuL/hezzl-test/internal/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Controller struct {
	services *services.Services
	logger   *zap.Logger
}

func newController(services *services.Services, logger *zap.Logger) *Controller {
	return &Controller{
		services: services,
		logger:   logger,
	}
}

func Provide() fx.Option {
	return fx.Provide(newController)
}
