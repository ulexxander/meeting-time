package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/ulexxander/meeting-time/graphql"
	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/storage"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	log := log.New(os.Stdout, "", log.LstdFlags)
	if err := run(log); err != nil {
		log.Fatalln("Fatal error:", err)
	}
}

type flags struct {
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

func run(log *log.Logger) error {
	log.Print("Parsing flags")
	flags := parseFlags()

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
	log.Print("GraphQL playground is enabled")

	log.Println("Starting HTTP server on", flags.addr)
	if err := http.ListenAndServe(flags.addr, nil); err != nil {
		return fmt.Errorf("listening HTTP: %w", err)
	}

	return nil
}

func setupDB(flags *flags, log *log.Logger) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=%s",
		flags.postgresHost,
		flags.postgresPort,
		flags.postgresUser,
		flags.postgresPassword,
		flags.postgresDatabase,
		flags.postgresSSLMode,
	)

	log.Println("Connecting to Postgres with DSN", dsn)
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

	log.Println("Running migrations")
	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("running migrations: %w", err)
	}

	return db, nil
}
