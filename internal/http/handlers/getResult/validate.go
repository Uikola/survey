package getResult

func validateReq(in *Input) error {
	if in.SurveyID < 1 {
		return ErrInvalidSurveyID
	}
	return nil
}
