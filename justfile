export GO_VERSION := "1.21"

_default:
  @just --list

_run:
  #!/bin/env bash
  export DATABASE_URL=postgresql://postgres:postgres@localhost:1233/postgres?sslmode=disable
  export DATABASE_MIGRATIONS=file://$(pwd)/migrations
  export PORT=1234
  go build -o dist/server cmd/server.go
  ./dist/server

# Run the app in development mode
run:
  docker compose up db swagger --detach --wait
  @just _run

# Run the app in development mode and watch for changes
run-watch:
  docker compose up db swagger --detach --wait
  watchexec -r -e go just _run

# Generate OpenAPI files
gen:
  go generate ./...

# Run e2e tests
test:
  #!/bin/env bash
  export TESTCONTAINERS_RYUK_DISABLED=true
  export DATABASE_MIGRATIONS=file://$(pwd)/migrations
  go test -v -cover ./...

# Run e2e tests in watch mode
test-watch:
  watchexec -r -e go just test

# Create a new migration file
migration-create name:
  go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.1.0 \
    create -ext sql -dir migrations -seq {{name}}

# Run pgcli to connect to the database
db-cli:
  docker compose exec db psql -U postgres -d postgres

# Deploy the app with docker-compose
deploy:
  docker compose up --build

# Cleanup the deployment
cleanup:
  docker compose down
