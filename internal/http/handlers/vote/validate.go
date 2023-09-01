package vote

func validateReq(in *Input) error {
	if in.AnswerID < 1 {
		return ErrInvalidAID
	}
	if in.SurveyID < 1 {
		return ErrInvalidSID
	}
	return nil
}
