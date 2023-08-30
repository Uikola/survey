package deleteAnswer

func validateReq(in *Input) error {
	if in.AnswerID < 0 {
		return ErrInvalidAID
	}
	if in.QuestionID < 0 {
		return ErrInvalidQID
	}
	if in.SurveyID < 0 {
		return ErrInvalidSID
	}
	return nil
}
