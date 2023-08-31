package entities

type Answer struct {
	ID       uint64
	Text     string
	SurveyID uint64
	Votes    uint64
}
