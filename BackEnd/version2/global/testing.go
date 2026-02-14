package global

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SetPoolForTests lets tests inject a pool without going through InitPG/pgOnce.
func SetPoolForTests(pool *pgxpool.Pool) {
	pgInstance = &Postgres{db: pool}
}

// ResetForTests clears the singleton so tests can re-init cleanly if needed.
func ResetForTests() {
	pgInstance = nil
	pgOnce = sync.Once{}
}
