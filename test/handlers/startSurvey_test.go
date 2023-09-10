package test

import (
	"github.com/stretchr/testify/require"
	"survey/internal/http/handlers/startSurvey"
	"testing"
)

func TestValidateSurvey(t *testing.T) {
	in := &startSurvey.Input{
		Title: "Test Text",
	}

	err := startSurvey.ValidateReq(in)
	if err != nil {
		t.Fatalf("invalid text: %s", err.Error())
	}

	require.NoError(t, err)
}

func TestValidateSurveyError(t *testing.T) {
	cases := []struct {
		name   string
		in     *startSurvey.Input
		expErr error
	}{
		{
			name:   "bad_title",
			in:     &startSurvey.Input{Title: ""},
			expErr: startSurvey.ErrInvalidTitle,
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := startSurvey.ValidateReq(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
