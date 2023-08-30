package handlers

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"survey/internal/db/repository/answerRepo"
	"survey/internal/db/repository/questionRepo"
	"survey/internal/db/repository/surveyRepo"
	"survey/internal/http/handlers/addAnswer"
	"survey/internal/http/handlers/addQuestion"
	"survey/internal/http/handlers/deleteAnswer"
	"survey/internal/http/handlers/deleteQuestion"
	"survey/internal/http/handlers/deleteSurvey"
	"survey/internal/http/handlers/startSurvey"
	"survey/internal/usecases/answerUC"
	"survey/internal/usecases/questionUC"
	"survey/internal/usecases/surveyUC"
)

func Router(db *sql.DB, router chi.Router, log logrus.FieldLogger) {
	surveyRepository := surveyRepo.NewSurveyRepo(db)
	questionRepository := questionRepo.NewQuestionRepo(db)
	answerRepository := answerRepo.NewAnswerRepo(db)
	surveyUseCase := surveyUC.NewUseCase(surveyRepository)
	questionUseCase := questionUC.NewUseCase(questionRepository)
	answerUseCase := answerUC.NewUseCase(answerRepository)

	router.Post("/start-survey", startSurvey.New(surveyUseCase, log))
	router.Post("/delete-survey", deleteSurvey.New(surveyUseCase, log))
	router.Post("/add-question", addQuestion.New(questionUseCase, log))
	router.Post("/delete-question", deleteQuestion.New(questionUseCase, log))
	router.Post("/add-ans", addAnswer.New(answerUseCase, log))
	router.Post("/delete-ans", deleteAnswer.New(answerUseCase, log))
}
