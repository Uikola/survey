package deleteSurvey

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"survey/internal/entities"
	"survey/internal/usecases/surveyUC"
)

// New @Summary Delete a survey
//
//	@Description	Deletes a survey
//	@Tags			survey
//	@Accept			json
//	@Produce		json
//	@Param			survey_id	path		int	true	"Survey ID"
//	@Success		200			{object}	entities.Response
//	@Failure		400			{object}	string
//	@Failure		500			{object}	string
//	@Router			/delete-survey/{survey_id} [delete]
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

		render.JSON(w, r, entities.Response{Message: "Survey deleted successfully"})
	}
}
