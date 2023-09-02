package addAnswer

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"survey/internal/entities"
	"survey/internal/usecases/answerUC"
)

var ErrInvalidText = errors.New("invalid text")
var ErrInvalidSurveyID = errors.New("invalid survey id")

func (in *Input) AnswerFromDTO() *entities.Answer {
	return &entities.Answer{
		Text:     in.Text,
		SurveyID: in.SurveyID,
	}
}

func New(uCase *answerUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		in := &Input{}
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			log.Errorf("can't parse the request: %s", err.Error())
			http.Error(w, "bad json(parsing): "+err.Error(), http.StatusBadRequest)
			return
		}

		err = validateReq(in)
		if err != nil {
			log.Errorf("can't validate the data: %s", err.Error())
			http.Error(w, "bad json(validating): "+err.Error(), http.StatusBadRequest)
			return
		}

		answer := in.AnswerFromDTO()
		response, err := uCase.AddAnswer(ctx, log, answer)
		if err != nil {
			log.Errorf("can't add a new answer: %s", err.Error())
			http.Error(w, "can't add a new answer: "+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, response)
	}
}
