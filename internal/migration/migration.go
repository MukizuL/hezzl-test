package migration

import (
	"context"
	"crypto/tls"
	"database/sql"
	"embed"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/MukizuL/hezzl-test/internal/config"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed "ch/*.sql"
var chMigrations embed.FS

//go:embed "pg/*.sql"
var pgMigrations embed.FS

type Migrator struct{}

func newMigrator(cfg *config.Config, logger *zap.Logger) (*Migrator, error) {
	logger.Info("Starting migrations!")

	logger.Info("Migrating ClickHouse...")
	err := migrateCh(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate ClickHouse: %w", err)
	}

	goose.ResetGlobalMigrations()

	logger.Info("Migrating Postgresql...")
	err = migratePg(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate Postgresql: %w", err)
	}

	return &Migrator{}, nil
}

func migrateCh(cfg *config.Config) error {
	db := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{cfg.ChAddr},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: cfg.ChUser,
			Password: cfg.ChPass,
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: time.Second * 30,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
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
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return err
	}

	goose.SetBaseFS(chMigrations)

	if err := goose.SetDialect("clickhouse"); err != nil {
		return err
	}

	// Should not be in release
	err := goose.Reset(db, "ch")
	if err != nil {
		return err
	}

	if err := goose.Up(db, "ch"); err != nil {
		return err
	}

	return nil
}

func migratePg(cfg *config.Config) error {
	db, err := sql.Open("pgx", cfg.PgDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(pgMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// Should not be in release
	err = goose.Reset(db, "pg")
	if err != nil {
		return err
	}

	if err := goose.Up(db, "pg"); err != nil {
		return err
	}

	return nil
}

func Provide() fx.Option {
	return fx.Provide(newMigrator)
}
