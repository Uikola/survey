package test

import (
	"github.com/stretchr/testify/require"
	"survey/internal/http/handlers/vote"
	"testing"
)

func TestValidateVote(t *testing.T) {
	in := &vote.Input{
		AnswerID: 1,
		SurveyID: 1,
	}
	err := vote.ValidateReq(in)
	if err != nil {
		t.Fatalf("invalid answerID and SurveyID")
	}

	require.NoError(t, err)
}

func TestValidateVoteError(t *testing.T) {
	cases := []struct {
		name   string
		in     *vote.Input
		expErr error
	}{
		{
			name:   "bad_answer_id",
			in:     &vote.Input{AnswerID: 0, SurveyID: 1},
			expErr: vote.ErrInvalidAID,
		},
		{
			name:   "bad_survey_id",
			in:     &vote.Input{AnswerID: 1, SurveyID: 0},
			expErr: vote.ErrInvalidSID,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := vote.ValidateReq(tCase.in)
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
