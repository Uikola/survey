package db

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"survey/internal/config"
)

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

	return db

}
