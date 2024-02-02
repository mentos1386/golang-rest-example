export GO_VERSION := "1.21"

_default:
  @just --list

# Run the app in development mode
run:
  docker compose up --build app swagger

delete:
  docker compose down

# Generate OpenAPI files
gen:
  docker compose run --build --rm app \
    go generate ./...


# Create a new migration file
migration-create name:
  @docker compose run --build --rm app \
    migrate create -ext sql -dir migrations -seq {{name}}

# Build production image
build:
  docker build \
    --build-arg GO_VERSION=$(GO_VERSION) \
    -t  golang-rest-example \
    .

# Run pgcli to connect to the database
db-cli:
  docker compose exec db psql -U postgres -d postgres
