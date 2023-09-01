package startSurvey

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate(t *testing.T) {
	in := &Input{
		Title: "Test Text",
	}

	err := validateReq(in)
	if err != nil {
		t.Fatalf("invalid text: %s", err.Error())
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
			name:   "bad_title",
			in:     &Input{Title: ""},
			expErr: ErrInvalidTitle,
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
