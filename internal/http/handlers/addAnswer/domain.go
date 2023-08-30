package addAnswer

type Input struct {
	Text       string `json:"text"`
	SurveyID   uint64 `json:"survey_id"`
	QuestionID uint64 `json:"question_id"`
}
