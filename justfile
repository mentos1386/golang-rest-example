export GO_VERSION := "1.21"

_default:
  @just --list

# Run the app in development mode
run:
  docker compose up --build app

# Generate OpenAPI files
gen:
  docker compose run --rm app \
    go generate ./...

# Build production image
build:
  docker build \
    --build-arg GO_VERSION=$(GO_VERSION) \
    -t  golang-rest-example \
    .
