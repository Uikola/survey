package startSurvey

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"survey/internal/entities"
	"survey/internal/usecases/surveyUC"
)

var ErrInvalidTitle = errors.New("invalid title")

func (in *Input) SurveyFromDTO() *entities.Survey {
	return &entities.Survey{
		Title: in.Title,
	}
}

// New @Summary Add an answer
//
//	@Description	Starts a new survey
//	@Tags			survey
//	@Accept			json
//	@Produce		json
//	@Param			input	body		Input	true	"startSurvey input"
//	@Success		200		{object}	entities.Survey
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Router			/start-survey [post]
func New(uCase *surveyUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		in := &Input{}
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			log.Errorf("can't parse the request: %s", err.Error())
			http.Error(w, "bad json(parsing):"+err.Error(), http.StatusBadRequest)
			return
		}

		err = ValidateReq(in)
		if err != nil {
			log.Errorf("can't validate the data: &s", err.Error())
			http.Error(w, "bad json(validating):"+err.Error(), http.StatusBadRequest)
			return
		}
		survey := in.SurveyFromDTO()
		response, err := uCase.StartSurvey(ctx, log, survey)
		if err != nil {
			log.Errorf("can't create the survey: %s", err.Error())
			http.Error(w, "create survey err:"+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, response)
	}
}
