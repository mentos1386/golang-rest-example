package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/mentos1386/golang-rest-example/pkg/api"
	"github.com/mentos1386/golang-rest-example/pkg/openapi"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupPostgres(ctx context.Context, t *testing.T) *postgres.PostgresContainer {
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	return postgresContainer
}

func setup(t *testing.T) (func(), *openapi.Server) {
	ctx := context.Background()
	postgresContainer := setupPostgres(ctx, t)

	connectionString, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to get connection string: %s", err)
	}
	os.Setenv("DATABASE_URL", connectionString)

	// Create a new API service
	service := api.NewApiService()
	// Create a new server
	server, err := openapi.NewServer(service)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container
	return func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}, server
}

func toJson(t *testing.T, request interface{ MarshalJSON() ([]byte, error) }) *bytes.Buffer {
	json, err := request.MarshalJSON()
	assert.Nil(t, err)
	return bytes.NewBuffer(json)
}

func TestHealth(t *testing.T) {
	teardown, srv := setup(t)
	defer teardown()

	// Create a new request
	req, err := http.NewRequest("GET", "/healthz", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"message\":\"OK\"}", rr.Body.String())
}

func TestCreatingGroup(t *testing.T) {
	teardown, srv := setup(t)
	defer teardown()

	// TODO: Maybe we can create a "newRequest" method to wrap this 6 lines?
	req, err := http.NewRequest("POST", "/groups", toJson(t, &openapi.GroupUpdate{Name: "test"}))
	assert.Nil(t, err)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	srv.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
