package test_setup

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/walle692/D0018E/BackEnd/version2/global"
)

func SetUpDB(t *testing.T) (ctx context.Context, pool *pgxpool.Pool) {
	t.Helper()
	load(".env")

	dsn := os.Getenv("DATABASE_URL")
	ctx = context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	require.NoError(t, err)

	global.SetPoolForTests(pool)

	t.Cleanup(pool.Close)
	return ctx, pool
}

//Shamelessly stolen from github comment to find path of .env when called not in root

// Load loads the environment variables from the .env file.
func load(envFile string) {
	err := godotenv.Load(dir(envFile))
	if err != nil {
		panic(fmt.Errorf("Error loading .env file: %w", err))
	}
}

// dir returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Errorf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
