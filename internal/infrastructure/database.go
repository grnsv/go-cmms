package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/grnsv/go-cmms/internal/config"
	postgres "github.com/grnsv/go-cmms/internal/infrastructure/postgres/sqlc"
)

// NewDatabase инициализирует подключение к БД и возвращает Querier
func NewDatabase(cfg config.DatabaseConfig) (*postgres.Queries, *sql.DB, error) {
	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeMS) * time.Millisecond)

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %w", err)
	}

	queries := postgres.New(db)
	return queries, db, nil
}
