package handlers

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"survey/internal/db/repository/surveyRepo"
	"survey/internal/http/handlers/deleteSurvey"
	"survey/internal/http/handlers/startSurvey"
	"survey/internal/usecases/surveyUC"
)

func Router(db *sql.DB, router chi.Router, log logrus.FieldLogger) {
	surveyRepository := surveyRepo.NewSurveyRepo(db)
	surveyUseCase := surveyUC.NewUseCase(surveyRepository)

	router.Post("/start-survey", startSurvey.New(surveyUseCase, log))
	router.Post("/delete-survey", deleteSurvey.New(surveyUseCase, log))
}
