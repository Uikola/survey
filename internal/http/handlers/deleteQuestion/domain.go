package deleteQuestion

type Input struct {
	SurveyID   uint64 `json:"survey_id"`
	QuestionID uint64 `json:"question_id"`
}
