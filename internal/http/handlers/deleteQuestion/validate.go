package deleteQuestion

func validateReq(in *Input) error {
	if in.SurveyID < 0 {
		return ErrInvalidSurveyID
	}
	if in.QuestionID < 0 {
		return ErrInvalidQuestionID
	}

	return nil
}
