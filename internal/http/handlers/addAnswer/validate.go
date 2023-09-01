package addAnswer

func validateReq(in *Input) error {
	if in.Text == "" {
		return ErrInvalidText
	}
	if in.SurveyID < 1 {
		return ErrInvalidSurveyID
	}

	return nil
}
