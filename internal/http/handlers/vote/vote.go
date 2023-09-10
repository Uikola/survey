package vote

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"survey/internal/entities"
	"survey/internal/usecases/answerUC"
)

var ErrInvalidAID = errors.New("invalid answer id")
var ErrInvalidSID = errors.New("invalid survey id")

// New @Summary Vote
//
//	@Description	Votes for the answer
//	@Tags			answer
//	@Accept			json
//	@Produce		json
//	@Param			input	body		Input	true	"Vote input"
//	@Success		200		{object}	entities.Response
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Router			/vote [post]
func New(uCase *answerUC.UseCase, log logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		in := &Input{}
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			log.Errorf("can't parse the request: %s", err.Error())
			http.Error(w, "bad json: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = ValidateReq(in)
		if err != nil {
			log.Errorf("can't validate the data: %s", err.Error())
			http.Error(w, "bad json(validating): "+err.Error(), http.StatusBadRequest)
			return
		}

		err = uCase.Vote(ctx, log, in.AnswerID, in.SurveyID)
		if err != nil {
			log.Errorf("voting error: %s", err.Error())
			http.Error(w, "voting error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, entities.Response{Message: "vote counted successfully"})
	}
}
