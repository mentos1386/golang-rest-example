package api

import (
	"database/sql"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/mentos1386/golang-rest-example/pkg/config"
	"github.com/mentos1386/golang-rest-example/pkg/openapi"
	"go.uber.org/zap"
)

type ApiService struct {
	logger *zap.Logger
	db     *sql.DB

	Config *config.Config

	openapi.UnimplementedHandler
}

func NewApiService() *ApiService {
	logger, _ := zap.NewDevelopment()

	config, err := config.NewConfig()
	if err != nil {
		logger.Fatal("Failed to load the config", zap.Error(err))
	}

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to the database", zap.Error(err))
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	err = db.Ping()
	if err != nil {
		logger.Fatal("Failed to ping the database", zap.Error(err))
	}
	logger.Info("Connected to the database")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal("Failed to create the driver", zap.Error(err))
	}
	m, err := migrate.NewWithDatabaseInstance(
		config.DatabaseMigrations,
		"postgres", driver)
	if err != nil {
		logger.Fatal("Failed to create the migration", zap.Error(err))
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Fatal("Failed to apply the migrations", zap.Error(err))
	}

	logger.Info("Migrations applied")

	return &ApiService{
		logger: logger,
		db:     db,
		Config: config,
	}
}
