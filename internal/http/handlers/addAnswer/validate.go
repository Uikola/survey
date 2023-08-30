package addAnswer

func validateReq(in *Input) error {
	if in.Text == "" {
		return ErrInvalidText
	}
	if in.QuestionID < 0 {
		return ErrInvalidQuestionID
	}
	if in.SurveyID < 0 {
		return ErrInvalidSurveyID
	}

	return nil
}
