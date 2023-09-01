package getResult

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate(t *testing.T) {
	in := &Input{
		SurveyID: 1,
	}

	err := validateReq(in)
	if err != nil {
		t.Fatalf("invalid surveyID")
	}

	require.NoError(t, err)
}

func TestValidateError(t *testing.T) {
	cases := []struct {
		name   string
		in     *Input
		expErr error
	}{
		{
			name:   "bad_answer_id",
			in:     &Input{SurveyID: 0},
			expErr: ErrInvalidSurveyID,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := validateReq(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
