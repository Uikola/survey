package db

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
	"survey/internal/config"
)

//go:embed migrations/*.sql
var schemaFs embed.FS

func InitDB(cfg config.Config, log logrus.FieldLogger) *sql.DB {
	db, err := sql.Open(cfg.DriverName, cfg.ConnString)
	if err != nil {
		log.Errorf("can't connect to the db: %s", err.Error())
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("can't ping db: %s", err.Error())
		return nil
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Errorf("Construct Migrate driver: %w", err)
		return nil
	}
	d, err := iofs.New(schemaFs, "migrations") // Get migrations from sql folder
	if err != nil {
		log.Error(err)
	}
	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)

	err = m.Up()
	if err != nil {
		log.Error(err)
	}

	return db

}
