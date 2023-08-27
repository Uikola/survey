package startSurvey

func validateReq(in *Input) error {
	if in.Title == "" {
		return ErrInvalidTitle
	}
	return nil
}
