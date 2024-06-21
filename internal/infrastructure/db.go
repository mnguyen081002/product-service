package infrastructure

import (
	"context"
	"fmt"
	config "productservice/config"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DatabaseTransaction interface {
	WithTransaction(txFunc func(tx *Database) error) (err error)
}

type Database struct {
	Sqlx    *sqlx.DB
	Context context.Context // use for mongo transaction
	logger  *zap.Logger
}

func NewDatabase(config *config.Config, logger *zap.Logger) Database {
	var err error

	logger.Info("Connecting to database...")
	sqlx, err := getDatabaseInstance(config)
	if err != nil {
		for i := 0; i < 5; i++ {
			sqlx, err = getDatabaseInstance(config)
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		logger.Fatal("Database connection error", zap.Error(err))
	} else {
		logger.Info("Database connected")
	}

	db := Database{Sqlx: sqlx, logger: logger}
	return db
}

func getDatabaseInstance(config *config.Config) (*sqlx.DB, error) {
	switch config.Database.Driver {
	case "postgres":
		db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			config.Database.Postgres.Host, config.Database.Postgres.Port, config.Database.Postgres.Username, config.Database.Postgres.Password, config.Database.Postgres.Name, config.Database.Postgres.SSLMode, config.Database.Postgres.TimeZone))

		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}

		return db, nil
	}

	return nil, fmt.Errorf("unsupported database driver: %s", config.Database.Driver)
}
