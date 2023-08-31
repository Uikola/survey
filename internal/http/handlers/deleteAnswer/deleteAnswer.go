package deleteAnswer

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"survey/internal/usecases/answerUC"
)

func New(uCase *answerUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		surveyID, err := strconv.Atoi(chi.URLParam(r, "survey_id"))
		if err != nil {
			log.Errorf("Invalid survey ID")
			http.Error(w, "Invalid survey ID", http.StatusBadRequest)
			return
		}
		answerID, err := strconv.Atoi(chi.URLParam(r, "ans_id"))
		if err != nil {
			log.Errorf("Invalid answer ID")
			http.Error(w, "Invalid answer ID", http.StatusBadRequest)
			return
		}

		err = uCase.DeleteAnswer(ctx, log, uint64(answerID), uint64(surveyID))
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
