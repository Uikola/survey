package deleteAnswer

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"survey/internal/usecases/answerUC"
)

var ErrInvalidAID = errors.New("invalid answer id")
var ErrInvalidQID = errors.New("invalid question id")
var ErrInvalidSID = errors.New("invalid survey id")

func New(uCase *answerUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
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
			http.Error(w, "bad json(validating)"+err.Error(), http.StatusBadRequest)
			return
		}

		err = uCase.DeleteAnswer(ctx, log, in.AnswerID, in.QuestionID, in.SurveyID)
		if err != nil {
			log.Errorf("can't delete answer: %s", err.Error())
			http.Error(w, "can't delete answer"+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"message": "answer deleted",
		})
	}
}
