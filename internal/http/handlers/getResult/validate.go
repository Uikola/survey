package getResult

func ValidateReq(in *Input) error {
	if in.SurveyID < 1 {
		return ErrInvalidSurveyID
	}
	return nil
}
