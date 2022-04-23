package testutil

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/db"
)

func Context(t *testing.T) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}

var (
	flagPostgresHost     = flag.String("postgres-host", "localhost", "PostgreSQL host")
	flagPostgresPort     = flag.Int("postgres-port", 5432, "PostgreSQL port")
	flagPostgresUser     = flag.String("postgres-user", "meeting-time", "PostgreSQL user")
	flagPostgresPassword = flag.String("postgres-password", "123", "PostgreSQL password")
	flagPostgresDatabase = flag.String("postgres-database", "meeting-time", "PostgreSQL database name")
	flagPostgresSSLMode  = flag.String("postgres-ssl-mode", "disable", "PostgreSQL SSL mode")
)

func Queries(t *testing.T) *db.Queries {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=%s",
		*flagPostgresHost,
		*flagPostgresPort,
		*flagPostgresUser,
		*flagPostgresPassword,
		*flagPostgresDatabase,
		*flagPostgresSSLMode,
	)

	postgresDB, err := sql.Open("postgres", dsn)
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := postgresDB.Close(); err != nil {
			t.Log("error closing db:", err)
		}
	})

	err = postgresDB.Ping()
	require.NoError(t, err)

	migrationsDriver, err := postgres.WithInstance(postgresDB, &postgres.Config{})
	require.NoError(t, err)

	migrations, err := migrate.NewWithDatabaseInstance("file://../db/migrations", "postgres", migrationsDriver)
	require.NoError(t, err)

	err = migrations.Down()
	if err != migrate.ErrNoChange {
		require.NoError(t, err)
	}

	err = migrations.Up()
	if err != migrate.ErrNoChange {
		require.NoError(t, err)
	}

	return db.New(postgresDB)
}
