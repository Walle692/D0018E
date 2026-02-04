package utils

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

func NewPG(ctx context.Context, connString string) (*Postgres, error) {
	var initErr error

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			initErr = err
			return
		}

		pgInstance = &Postgres{db}
	})

	if initErr != nil {
		fmt.Println("DEBUG: initerror")
		return nil, initErr
	}

	if pgInstance == nil {
		fmt.Println("DEBUG: Pg instance nil")
		return nil, fmt.Errorf("postgres instance not initialized")
	}

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}
