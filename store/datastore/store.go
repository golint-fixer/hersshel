package datastore

import (
	"database/sql"
	"time"

	"github.com/Sirupsen/logrus"
	_ "github.com/lib/pq" // Postgres Driver

	"github.com/hersshel/hersshel/config"
	"github.com/hersshel/hersshel/model"
	"github.com/hersshel/hersshel/store"
	"github.com/hersshel/hersshel/store/datastore/ddl"

	"github.com/rubenv/sql-migrate"
	"gopkg.in/gorp.v1"
)

// datastore is an implementation of a store.Store built on top
// of the sql/database driver with a relational database backend.
type datastore struct {
	*gorp.DbMap
}

// New creates a database connection for the given datasource configuration
// and returns a new Store.
func New(cfg *config.PostgreSQL) store.Store {
	if cfg == nil {
		logrus.Fatal("database connection failed: not configured")
	}
	return From(
		open(cfg),
	)
}

// From returns a Store using an existing database connection.
func From(db *gorp.DbMap) store.Store {
	return &datastore{db}
}

// open opens a new database connection with the specified
// driver and connection string and returns a gorp.DbMap.
func open(cfg *config.PostgreSQL) *gorp.DbMap {
	db, err := cfg.DB()
	if err != nil {
		logrus.Fatalf("database connection failed: %v", err)
	}

	gorp := setupGorp(db)

	if err := pingDatabase(db); err != nil {
		logrus.Fatalf("database ping attempts failed: %v", err)
	}

	if err := setupDatabase(db); err != nil {
		logrus.Fatalf("database migration failed: %v", err)
	}
	return gorp
}

// pingDatabase is an helper function to ping the database with backoff
// to ensure a connection can be established before we proceed with the
// database setup and migration.
func pingDatabase(db *sql.DB) error {
	var err error
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return nil
		}
		logrus.Infof("database ping failed. retry in 1s")
		time.Sleep(time.Second)
	}
	return err
}

// setupGorp will execute SQL queries to create the app tables.
func setupGorp(db *sql.DB) *gorp.DbMap {
	dbmap := &gorp.DbMap{Db: db, Dialect: &gorp.PostgresDialect{}}
	dbmap.AddTableWithNameAndSchema(model.Category{}, "hersshel", "category").SetKeys(true, "id")
	dbmap.AddTableWithNameAndSchema(model.Feed{}, "hersshel", "feed").SetKeys(true, "id")
	dbmap.AddTableWithNameAndSchema(model.Item{}, "hersshel", "item").SetKeys(true, "id")
	return dbmap
}

// setupDatabase performs a database migration.
func setupDatabase(db *sql.DB) error {
	var migrations = &migrate.AssetMigrationSource{
		AssetDir: ddl.AssetDir,
		Asset:    ddl.Asset,
		Dir:      "postgres",
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	return err
}
