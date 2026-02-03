package global

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

// function to inititalize the global pool
func InitPG(ctx context.Context, connString string) error {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			pgErr = err
			return
		}

		pgInstance = &Postgres{db}
	})

	return pgErr
}

// function to get the db connection from pool
func Get() *Postgres {
	if pgInstance == nil {
		panic("postgres not inited")
	}
	return pgInstance
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}

func (pg *Postgres) Pool() *pgxpool.Pool {
	return pg.db
}
