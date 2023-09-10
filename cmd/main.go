package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"survey/internal/config"
	"survey/internal/db"
	"survey/internal/http/handlers"
	"survey/pkg/logger"

	"net/http"
)

//	@title			Survey
//	@version		1.0
//	@description	API Server for Survey Application

//	@host		localhost:8080
//	@BasePath	/api

func main() {
	log := logger.New()
	if err := mainNoExit(log); err != nil {
		log.Fatalf("fatal err: %s", err.Error())
	}
}

func mainNoExit(log logrus.FieldLogger) error {
	log.Infoln("initializing config")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error while loading config")
	}

	log.Infoln("initializing db")
	dataBase := db.InitDB(cfg, log)
	defer dataBase.Close()

	log.Infoln("initializing router")
	router := chi.NewRouter()
	handlers.Router(dataBase, router, log)

	log.Infoln("starting app")
	return http.ListenAndServe(
		cfg.Port,
		router,
	)
}
