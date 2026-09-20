package core_postgres_pool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Close()
	OpTimeout() time.Duration
}

type ConnectionPool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func (p *ConnectionPool) OpTimeout() time.Duration {
	return p.opTimeout
}

func NewConnectionPool(config Config, ctx context.Context) (*ConnectionPool, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Port,
		config.Host,
		config.Database,
		config.Ssl,
	)

	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("Parse pgx config: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("Create pgx pool: %w", err)
	}

	return &ConnectionPool{
		Pool:      pool,
		opTimeout: config.Timeout,
	}, nil

}
