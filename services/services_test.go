package services_test

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var (
	flagPostgresHost     = flag.String("postgres-host", "localhost", "PostgreSQL host")
	flagPostgresPort     = flag.Int("postgres-port", 5432, "PostgreSQL port")
	flagPostgresUser     = flag.String("postgres-user", "meeting-time", "PostgreSQL user")
	flagPostgresPassword = flag.String("postgres-password", "123", "PostgreSQL password")
	flagPostgresDatabase = flag.String("postgres-database", "meeting-time", "PostgreSQL database name")
	flagPostgresSSLMode  = flag.String("postgres-ssl-mode", "disable", "PostgreSQL SSL mode")
)

func setupDB(t *testing.T) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=%s",
		*flagPostgresHost,
		*flagPostgresPort,
		*flagPostgresUser,
		*flagPostgresPassword,
		*flagPostgresDatabase,
		*flagPostgresSSLMode,
	)

	const attempts = 5
	var db *sqlx.DB
	for i := 0; i < attempts; i++ {
		var err error
		db, err = sqlx.Connect("postgres", dsn)
		if err != nil {
			t.Logf("Error connecting to Postgres, attempts left: %d", attempts-(i+1))
			time.Sleep(time.Second)
			continue
		}
		break
	}
	if db == nil {
		require.Fail(t, "Could not connect to Postgres")
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Log("error closing db:", err)
		}
	})

	migrationsDriver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	require.NoError(t, err)

	migrations, err := migrate.NewWithDatabaseInstance("file://../storage/migrations", "postgres", migrationsDriver)
	require.NoError(t, err)

	err = migrations.Down()
	if err != migrate.ErrNoChange {
		require.NoError(t, err)
	}

	err = migrations.Up()
	if err != migrate.ErrNoChange {
		require.NoError(t, err)
	}

	return db
}
