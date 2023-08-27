package deleteSurvey

func validateReq(in *Input) error {
	if in.SurveyID < 0 {
		return ErrInvalidSurveyID
	}

	return nil
}
