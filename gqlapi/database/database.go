package database

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq" // ...
	"go.uber.org/zap"

	"gqlapi/config"
)

// Database ...
type Database struct {
	db             *dbx.DB
	config         *config.DatabaseConfig
	logger         *zap.Logger
	userRepository *UserRepository
}

// NewDatabase ...
func NewDatabase(conf *config.Config, log *zap.Logger) *Database {
	return &Database{
		config: conf.Database,
		logger: log,
	}
}

// Open ...
func (d *Database) Open() error {
	db, err := dbx.Open(d.config.Driver, d.config.Dsn)
	if err != nil {
		return err
	}

	if err := db.DB().Ping(); err != nil {
		return err
	}

	d.db = db
	d.logger.Info("Database connection opened successfully")
	return nil
}

// Close ...
func (d *Database) Close() {
	d.logger.Info("Closing database connection...")
	d.db.Close()
}

// User ...
func (d *Database) User() *UserRepository {
	if d.userRepository == nil {
		d.userRepository = &UserRepository{
			db: d.db,
		}
	}
	return d.userRepository
}
