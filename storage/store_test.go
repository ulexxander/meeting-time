package storage_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	flagPostgresHost     = flag.String("postgres-host", "localhost", "PostgreSQL host")
	flagPostgresPort     = flag.Int("postgres-port", 5432, "PostgreSQL port")
	flagPostgresUser     = flag.String("postgres-user", "meeting-time", "PostgreSQL user")
	flagPostgresPassword = flag.String("postgres-password", "123", "PostgreSQL password")
	flagPostgresDatabase = flag.String("postgres-database", "meeting-time", "PostgreSQL database name")
)

func setupDB(t *testing.T) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s",
		*flagPostgresHost,
		*flagPostgresPort,
		*flagPostgresUser,
		*flagPostgresPassword,
		*flagPostgresDatabase,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)
	err = db.AutoMigrate(storage.Organization{})
	require.NoError(t, err)
	return db
}
