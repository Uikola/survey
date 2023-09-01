package vote

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate(t *testing.T) {
	in := &Input{
		AnswerID: 1,
		SurveyID: 1,
	}
	err := validateReq(in)
	if err != nil {
		t.Fatalf("invalid answerID and SurveyID")
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
			in:     &Input{AnswerID: 0, SurveyID: 1},
			expErr: ErrInvalidAID,
		},
		{
			name:   "bad_survey_id",
			in:     &Input{AnswerID: 1, SurveyID: 0},
			expErr: ErrInvalidSID,
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
