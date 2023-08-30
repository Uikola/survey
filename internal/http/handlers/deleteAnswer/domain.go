package deleteAnswer

type Input struct {
	AnswerID   uint64 `json:"answer_id"`
	SurveyID   uint64 `json:"survey_id"`
	QuestionID uint64 `json:"question_id"`
}
