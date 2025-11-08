package config

import (
	"context"
	"os"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	dsn string
}

func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		dsn: os.Getenv("DATABASE_URL"),
	}
}

func (conn *PostgresConfig) PGconnect() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), conn.dsn)
	if err != nil {
		logger.Error(nil, "Error connecting to pg database: %v", err)
		return nil, err
	}

	if err = pool.Ping(context.Background()); err != nil {
		logger.Error(nil, "Error pinging pg database: %v", err)
		return nil, err
	}

	logger.Info(nil, "✅ Connected to pg database")
	return pool, nil
}
