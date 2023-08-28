package addQuestion

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"survey/internal/entities"
	"survey/internal/usecases/questionUC"
)

var ErrInvalidSurveyID = errors.New("invalid survey id")
var ErrInvalidText = errors.New("invalid text")

func (in *Input) QuestionFromDTO() *entities.Question {
	return &entities.Question{
		Text:     in.Text,
		SurveyID: in.SurveyID,
	}
}

func New(uCase *questionUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		in := &Input{}
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			log.Errorf("can't parse the request: %s", err.Error())
			http.Error(w, "bad json(parsing)"+err.Error(), http.StatusBadRequest)
			return
		}

		err = validateReq(in)
		if err != nil {
			log.Errorf("can't validate the data: %s", err.Error())
			http.Error(w, "bad json(validating)"+err.Error(), http.StatusInternalServerError)
			return
		}

		question := in.QuestionFromDTO()
		response, err := uCase.AddQuestion(ctx, log, question)
		if err != nil {
			log.Errorf("can't add question: %s", err.Error())
			http.Error(w, "can't add question"+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, response)
	}
}
