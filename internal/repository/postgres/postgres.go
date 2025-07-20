package postgres

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/maxviazov/signal-flow/internal/config"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*Repository, error) {
	// Валидация входных параметров
	if ctx == nil {
		return nil, fmt.Errorf("context cannot be nil")
	}
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	// Безопасное формирование DSN с экранированием
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		url.QueryEscape(cfg.Postgres.POSTGRES_USER),
		url.QueryEscape(cfg.Postgres.POSTGRES_PASSWORD),
		cfg.Postgres.POSTGRES_HOST,
		cfg.Postgres.POSTGRES_PORT,
		url.QueryEscape(cfg.Postgres.POSTGRES_DB),
	)

	// Настройка пула соединений
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// Настройка параметров пула
	config.MaxConns = 30
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &Repository{pool: pool}, nil
}

func (r *Repository) Close() {
	if r.pool != nil {
		r.pool.Close()
	}
}
