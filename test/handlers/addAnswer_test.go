package test

import (
	"github.com/stretchr/testify/require"
	"survey/internal/http/handlers/addAnswer"
	"testing"
)

func TestValidateAns(t *testing.T) {
	in := &addAnswer.Input{
		Text:     "Test Text",
		SurveyID: 1,
	}

	err := addAnswer.ValidateReq(in)
	if err != nil {
		t.Fatalf("invalid text and surveyID: %s", err.Error())
	}

	require.NoError(t, err)
}

func TestValidateAnsError(t *testing.T) {
	cases := []struct {
		name   string
		in     *addAnswer.Input
		expErr error
	}{
		{
			name:   "bad_text",
			in:     &addAnswer.Input{Text: "", SurveyID: 1},
			expErr: addAnswer.ErrInvalidText,
		},
		{
			name:   "bad_survey_id",
			in:     &addAnswer.Input{Text: "Test Text", SurveyID: 0},
			expErr: addAnswer.ErrInvalidSurveyID,
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := addAnswer.ValidateReq(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
