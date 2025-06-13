package main

import (
	"github.com/MukizuL/hezzl-test/internal/config"
	"github.com/MukizuL/hezzl-test/internal/controller"
	"github.com/MukizuL/hezzl-test/internal/migration"
	"github.com/MukizuL/hezzl-test/internal/router"
	"github.com/MukizuL/hezzl-test/internal/server"
	"github.com/MukizuL/hezzl-test/internal/services"
	"github.com/MukizuL/hezzl-test/internal/storage"
	"github.com/MukizuL/hezzl-test/internal/storage/drivers/ch"
	"github.com/MukizuL/hezzl-test/internal/storage/drivers/nats"
	"github.com/MukizuL/hezzl-test/internal/storage/drivers/pg"
	"github.com/MukizuL/hezzl-test/internal/zlog"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		createApp(),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func createApp() fx.Option {
	return fx.Options(
		config.Provide(),
		fx.Provide(zap.NewDevelopment),

		migration.Provide(),

		controller.Provide(),
		router.Provide(),
		server.Provide(),

		pg.Provide(),
		ch.Provide(),
		storage.Provide(),

		nats.Provide(),
		zlog.Provide(),

		services.Provide(),
	)
}
