package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TimeScaleDB struct {
	pool *pgxpool.Pool
}

func NewTimeScaleDB(connStr string) (*TimeScaleDB, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return &TimeScaleDB{pool: pool}, nil
}

func (receiver *TimeScaleDB) FindXLogByTxid() *XLog {
	return nil
}

func (receiver *TimeScaleDB) FindAll() []XLog {
	return nil
}

func (receiver *TimeScaleDB) PoolClose() {
	receiver.pool.Close()
}
