package handlers

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"survey/internal/db/repository/answerRepo"
	"survey/internal/db/repository/surveyRepo"
	"survey/internal/http/handlers/addAnswer"
	"survey/internal/http/handlers/deleteAnswer"
	"survey/internal/http/handlers/deleteSurvey"
	"survey/internal/http/handlers/startSurvey"
	"survey/internal/http/handlers/vote"
	"survey/internal/usecases/answerUC"
	"survey/internal/usecases/surveyUC"
)

func Router(db *sql.DB, router chi.Router, log logrus.FieldLogger) {
	surveyRepository := surveyRepo.NewSurveyRepo(db)
	answerRepository := answerRepo.NewAnswerRepo(db)
	surveyUseCase := surveyUC.NewUseCase(surveyRepository)
	answerUseCase := answerUC.NewUseCase(answerRepository)

	router.Route("/api", func(r chi.Router) {
		r.Post("/start-survey", startSurvey.New(surveyUseCase, log))
		r.Delete("/delete-survey/{survey_id}", deleteSurvey.New(surveyUseCase, log))
		r.Post("/add-ans", addAnswer.New(answerUseCase, log))
		r.Delete("/delete-ans/{survey_id}/{ans_id}", deleteAnswer.New(answerUseCase, log))
		r.Post("/vote", vote.New(answerUseCase, log))
	})

}
