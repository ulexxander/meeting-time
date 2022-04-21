package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/ulexxander/meeting-time/graphql"
	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/storage"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	log := logrus.New()
	if err := run(log); err != nil {
		log.WithError(err).Fatal("Fatal error")
	}
}

type flags struct {
	logLevel         string
	addr             string
	postgresHost     string
	postgresPort     int
	postgresUser     string
	postgresPassword string
	postgresDatabase string
	postgresSSLMode  string
}

func parseFlags() *flags {
	var flags flags
	flag.StringVar(&flags.logLevel, "log-level", logrus.TraceLevel.String(), "Log level, available: "+logLevels())
	flag.StringVar(&flags.addr, "addr", ":80", "HTTP server address")
	flag.StringVar(&flags.postgresHost, "postgres-host", "localhost", "PostgreSQL host")
	flag.IntVar(&flags.postgresPort, "postgres-port", 5432, "PostgreSQL port")
	flag.StringVar(&flags.postgresUser, "postgres-user", "meeting-time", "PostgreSQL user")
	flag.StringVar(&flags.postgresPassword, "postgres-password", "123", "PostgreSQL password")
	flag.StringVar(&flags.postgresDatabase, "postgres-database", "meeting-time", "PostgreSQL database name")
	flag.StringVar(&flags.postgresSSLMode, "postgres-ssl-mode", "disable", "PostgreSQL SSL mode")
	flag.Parse()
	return &flags
}

func logLevels() string {
	var levels []string
	for _, l := range logrus.AllLevels {
		levels = append(levels, l.String())
	}
	return strings.Join(levels, ", ")
}

func run(log *logrus.Logger) error {
	log.Info("Parsing flags")
	flags := parseFlags()

	logLevel, err := logrus.ParseLevel(flags.logLevel)
	if err != nil {
		return fmt.Errorf("parsing log level: %w", err)
	}
	log.SetLevel(logLevel)

	db, err := setupDB(flags, log)
	if err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}
	defer db.Close()

	organizationsStore := storage.OrganizationsStore{DB: db}

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{
			OrganizationsStore: organizationsStore,
		},
	})
	server := handler.NewDefaultServer(schema)

	http.Handle("/query", server)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	log.Warn("GraphQL playground is enabled")

	log.WithField("addr", flags.addr).Info("Starting HTTP server")
	if err := http.ListenAndServe(flags.addr, nil); err != nil {
		return fmt.Errorf("listening HTTP: %w", err)
	}

	return nil
}

func setupDB(flags *flags, log *logrus.Logger) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=%s",
		flags.postgresHost,
		flags.postgresPort,
		flags.postgresUser,
		flags.postgresPassword,
		flags.postgresDatabase,
		flags.postgresSSLMode,
	)

	log.WithFields(logrus.Fields{
		"host":     flags.postgresHost,
		"port":     flags.postgresPort,
		"user":     flags.postgresUser,
		"database": flags.postgresDatabase,
		"sslMode":  flags.postgresSSLMode,
	}).Info("Connecting to Postgres")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("connecting to Postgres: %w", err)
	}

	migrationsDriver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("instanciating Postgres migrations driver: %w", err)
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://storage/migrations", "postgres", migrationsDriver)
	if err != nil {
		return nil, fmt.Errorf("instanciating migrations: %w", err)
	}

	version, dirty, err := migrations.Version()
	if err != nil {
		return nil, fmt.Errorf("querying migrations version: %w", err)
	}

	log.WithFields(logrus.Fields{
		"version": version,
		"dirty":   dirty,
	}).Info("Running migrations")
	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("running migrations: %w", err)
	}

	return db, nil
}
