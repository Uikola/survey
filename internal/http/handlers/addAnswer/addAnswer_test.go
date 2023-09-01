package addAnswer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate(t *testing.T) {
	in := &Input{
		Text:     "Test Text",
		SurveyID: 1,
	}

	err := validateReq(in)
	if err != nil {
		t.Fatalf("invalid text and surveyID: %s", err.Error())
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
			name:   "bad_text",
			in:     &Input{Text: "", SurveyID: 1},
			expErr: ErrInvalidText,
		},
		{
			name:   "bad_survey_id",
			in:     &Input{Text: "Test Text", SurveyID: 0},
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
