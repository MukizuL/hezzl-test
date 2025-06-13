package ch

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/MukizuL/hezzl-test/internal/config"
	"go.uber.org/fx"
	"time"
)

func newClickHouseConnection(lc fx.Lifecycle, cfg *config.Config) (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.ChAddr},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: cfg.ChUser,
			Password: cfg.ChPass,
		},
		//DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
		//	dialCount++
		//	var d net.Dialer
		//	return d.DialContext(ctx, "tcp", addr)
		//},
		Debug: false,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "hezzl-test", Version: "0.1"},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open clickhouse connection: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return conn.Ping(ctx)
		},
		OnStop: func(ctx context.Context) error {
			conn.Close()
			return nil
		},
	})

	return conn, nil
}

func Provide() fx.Option {
	return fx.Provide(newClickHouseConnection)
}
