package startSurvey

func ValidateReq(in *Input) error {
	if in.Title == "" {
		return ErrInvalidTitle
	}
	return nil
}
