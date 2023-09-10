package deleteAnswer

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"survey/internal/entities"
	"survey/internal/usecases/answerUC"
)

// New @Summary Delete an answer
//
//	@Description	Deletes an answer
//	@Tags			answer
//	@Accept			json
//	@Produce		json
//	@Param			survey_id	path		int	true	"Survey ID"
//	@Param			ans_id	path		int	true	"Answer ID"
//	@Success		200			{object}	entities.Response
//	@Failure		400			{object}	string
//	@Failure		500			{object}	string
//	@Router			/delete-ans/{survey_id}/{ans_id} [delete]
func New(uCase *answerUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		surveyID, err := strconv.Atoi(chi.URLParam(r, "survey_id"))
		if err != nil || surveyID < 1 {
			log.Errorf("Invalid survey ID")
			http.Error(w, "Invalid survey ID", http.StatusBadRequest)
			return
		}
		answerID, err := strconv.Atoi(chi.URLParam(r, "ans_id"))
		if err != nil || answerID < 1 {
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

		render.JSON(w, r, entities.Response{Message: "Answer deleted successfully"})
	}
}
