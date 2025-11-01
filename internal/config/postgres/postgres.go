package postgres

import (
	"database/sql"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	_ "github.com/lib/pq"
	"os"
)

type PostgresConfig struct {
	dsn string
}

func NewPostgresConfig() *PostgresConfig {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatalf("Ошибка загрузки .env файла: %v", err)
	//}

	return &PostgresConfig{
		dsn: os.Getenv("DATABASE_URL"),
	}
}

func (conn *PostgresConfig) PGconnect() (*sql.DB, error) {

	db, err := sql.Open("postgres", conn.dsn)
	if err != nil {
		logger.Error(nil, "Error connecting to pg database: %v", err)
		return nil, err
	}
	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		logger.Error(nil, "Error pinging pg database: %v", err)
		return nil, err
	}

	logger.Info(nil, "✅ Connected to pg database")
	return db, nil
}
