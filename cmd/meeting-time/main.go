package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ulexxander/meeting-time/graphql"
	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	flagAddr             = flag.String("addr", ":80", "HTTP server address")
	flagPostgresHost     = flag.String("postgres-host", "localhost", "PostgreSQL host")
	flagPostgresPort     = flag.Int("postgres-port", 5432, "PostgreSQL port")
	flagPostgresUser     = flag.String("postgres-user", "meeting-time", "PostgreSQL user")
	flagPostgresPassword = flag.String("postgres-password", "123", "PostgreSQL password")
	flagPostgresDatabase = flag.String("postgres-database", "meeting-time", "PostgreSQL database name")
)

func main() {
	log := log.New(os.Stdout, "", log.LstdFlags)
	if err := run(log); err != nil {
		log.Fatalln("Fatal error:", err)
	}
}

func run(log *log.Logger) error {
	log.Print("Parsing flags")
	flag.Parse()

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s",
		*flagPostgresHost,
		*flagPostgresPort,
		*flagPostgresUser,
		*flagPostgresPassword,
		*flagPostgresDatabase,
	)
	log.Println("Connecting to PostgreSQL with DSN", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("opening db: %w", err)
	}
	store := &storage.Store{DB: db}
	organizationsStore := storage.OrganizationsStore{Store: store}

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{
			OrganizationsStore: organizationsStore,
		},
	})
	server := handler.NewDefaultServer(schema)

	http.Handle("/query", server)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	log.Println("GraphQL playground is enabled")

	log.Println("Starting HTTP server on", *flagAddr)
	if err := http.ListenAndServe(*flagAddr, nil); err != nil {
		return fmt.Errorf("listening HTTP: %w", err)
	}

	return nil
}
