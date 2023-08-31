package deleteSurvey

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"survey/internal/usecases/surveyUC"
)

func New(uCase *surveyUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		surveyID, err := strconv.Atoi(chi.URLParam(r, "survey_id"))
		if err != nil {
			log.Errorf("Invalid survey id")
			http.Error(w, "Invalid survey id", http.StatusBadRequest)
			return
		}

		err = uCase.DeleteSurvey(ctx, log, uint64(surveyID))
		if err != nil {
			log.Errorf("can't delete the survey")
			http.Error(w, "can't delete the survey"+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"message": "survey deleted",
		})
	}
}
