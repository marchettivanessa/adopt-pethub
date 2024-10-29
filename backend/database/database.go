package database

import (
	"adopt-pethub/backend/config"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Postgres = "postgres"
)

type Database struct {
	Config     config.DatabaseConfig
	Connection *gorm.DB
}

// Função para criar a conexão com o banco de dados usando o GORM
func newDatabaseConn(dbc config.DatabaseConfig) (*Database, error) {
	connectionString := buildConnectionString(dbc)

	// Criando a conexão usando GORM e o driver PostgreSQL
	gormDB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Errorf("failed to create the database connection: %v", err)
		return nil, fmt.Errorf("failed to create the database connection: %w", err)
	}

	return &Database{
		Config:     dbc,
		Connection: gormDB,
	}, nil
}

// Função para construir a string de conexão
func buildConnectionString(dbc config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d search_path=public,%s sslmode=disable",
		dbc.Host,
		dbc.Port,
		dbc.Username,
		dbc.Password,
		dbc.Name,
		dbc.ConnectTimeout,
		dbc.Schema,
	)
}

// Migrate execute migrations, using GORM conection
func (db *Database) Migrate() error {
	sqlDB, err := db.Connection.DB() // Conects with GORM
	if err != nil {
		return fmt.Errorf("failed to get native DB connection from GORM: %w", err)
	}

	driver, err := migratePostgres.WithInstance(sqlDB, &migratePostgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	migrationFullPath := fmt.Sprintf("file://%s", db.Config.MigrationPath)
	migration, err := migrate.NewWithDatabaseInstance(migrationFullPath, Postgres, driver)
	if err != nil {
		log.Errorf("failed initializing database migration: %v", err)
		return fmt.Errorf("failed initializing database migration: %w", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed migrating the database: %w", err)
	}

	log.Info("Database migration successful")
	return nil
}

// NewDatabaseWithMigrations creates a db connection and run the migrations
func NewDatabaseWithMigrations(c config.DatabaseConfig) (*Database, error) {
	database, err := newDatabaseConn(c)
	if err != nil {
		return nil, fmt.Errorf("failed to create the database connection: %w", err)
	}
	err = database.Migrate()
	if err != nil {
		return nil, fmt.Errorf("failed to migrate the database: %w", err)
	}
	return database, nil
}

// ResetMigration resets migrations
func (db *Database) ResetMigration() error {
	updateStatement := "DROP TABLE IF EXISTS schema_migrations"
	sqlDB, err := db.Connection.DB()
	if err != nil {
		return fmt.Errorf("failed to get native DB connection: %w", err)
	}
	_, err = sqlDB.Exec(updateStatement)
	if err != nil {
		return fmt.Errorf("failed resetting migrations: %w", err)
	}
	return nil
}
