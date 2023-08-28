package addQuestion

func validateReq(in *Input) error {
	if in.SurveyID < 0 {
		return ErrInvalidSurveyID
	}
	if in.Text == "" {
		return ErrInvalidText
	}
	return nil
}
