package getResult

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"survey/internal/usecases/surveyUC"
)

var ErrInvalidSurveyID = errors.New("invalid survey id")

// New @Summary Get a result
//
//	@Description	Gets a result of the survey
//	@Tags			survey
//	@Accept			json
//	@Produce		json
//	@Param			survey_id	body		Input	true	"Survey ID"
//	@Success		200			{object}	entities.Survey
//	@Failure		400			{object}	string
//	@Failure		500			{object}	string
//	@Router			/get-result [post]
func New(uCase *surveyUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		in := &Input{}
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			log.Errorf("can't parse the request: %s", err.Error())
			http.Error(w, "bad json(parsing): "+err.Error(), http.StatusBadRequest)
			return
		}

		err = ValidateReq(in)
		if err != nil {
			log.Errorf("can't validate the data: %s", err.Error())
			http.Error(w, "bad json(validating): "+err.Error(), http.StatusBadRequest)
			return
		}

		response, err := uCase.GetResult(ctx, log, in.SurveyID)
		if err != nil {
			log.Errorf("can't get result: %s", err.Error())
			http.Error(w, "can't get result: "+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, response)
	}
}
