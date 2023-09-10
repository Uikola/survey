package test

import (
	"github.com/stretchr/testify/require"
	"survey/internal/http/handlers/getResult"
	"testing"
)

func TestValidate(t *testing.T) {
	in := &getResult.Input{
		SurveyID: 1,
	}

	err := getResult.ValidateReq(in)
	if err != nil {
		t.Fatalf("invalid surveyID")
	}

	require.NoError(t, err)
}

func TestValidateError(t *testing.T) {
	cases := []struct {
		name   string
		in     *getResult.Input
		expErr error
	}{
		{
			name:   "bad_answer_id",
			in:     &getResult.Input{SurveyID: 0},
			expErr: getResult.ErrInvalidSurveyID,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := getResult.ValidateReq(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
