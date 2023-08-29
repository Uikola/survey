package deleteQuestion

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"survey/internal/usecases/questionUC"
)

var ErrInvalidSurveyID = errors.New("invalid survey id")
var ErrInvalidQuestionID = errors.New("invalid question id")

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
			http.Error(w, "bad json(validating)"+err.Error(), http.StatusBadRequest)
			return
		}

		err = uCase.DeleteQuestion(ctx, log, in.SurveyID, in.QuestionID)
		if err != nil {
			log.Errorf("can't delete question: %s", err.Error())
			http.Error(w, "can't delete question"+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"message": "question deleted",
		})
	}
}
