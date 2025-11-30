package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"sso/internal/app/config"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func NewClient(ctx context.Context, maxAttempts int, cf config.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cf.User, cf.Password, cf.Hostname, cf.Port, cf.DBName,
	)

	var pool *pgxpool.Pool

	err := doWithTries(func() error {
		c, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		var err error
		pool, err = pgxpool.Connect(c, dsn)
		return err
	}, maxAttempts, 5*time.Second)

	if err != nil {
		return nil, fmt.Errorf("cannot connect to postgres: %w", err)
	}

	return pool, nil
}

func doWithTries(fn func() error, attempts int, delay time.Duration) error {
	var err error

	for attempts > 0 {
		if err = fn(); err == nil {
			return nil
		}

		attempts--
		if attempts > 0 {
			time.Sleep(delay)
		}
	}

	return err
}
