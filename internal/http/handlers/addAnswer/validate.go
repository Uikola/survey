package addAnswer

func validateReq(in *Input) error {
	if in.Text == "" {
		return ErrInvalidText
	}
	if in.SurveyID < 0 {
		return ErrInvalidSurveyID
	}

	return nil
}
