package vote

func validateReq(in *Input) error {
	if in.AnswerID < 0 {
		return ErrInvalidAID
	}
	if in.SurveyID < 0 {
		return ErrInvalidSID
	}
	return nil
}
