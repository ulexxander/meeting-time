# Meeting Time

Interactive meeting timekeeper.

## Tech stack

Backend:

- Go programming language
- PostgreSQL relational database
- SQLX for accessing the database
- Golang Migrate for database migrations
- GraphQL API

## Development

### Requirements

- Go 1.17
- Docker
- Docker Compose
- [gqlgen](https://gqlgen.com)
- [migrate](https://github.com/golang-migrate/migrate)

### Setup

```sh
# Spin up database.
docker-compose up -d

# Run tests.
go test ./services

# Run service locally on port 4000.
go run ./cmd/meeting-time/ -addr=:4000
```

### Adding new features

```sh
# Adding database migration.
migrate create -ext sql -dir storage/migrations -seq migration_name
```

## Links

- [Scrum meeting](https://www.productplan.com/glossary/scrum-meeting)
- [What is a standup meeting](https://www.wework.com/ideas/professional-development/management-leadership/what-is-a-standup-meeting)
