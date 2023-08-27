package entities

type Answer struct {
	ID       uint64
	Text     string
	Question uint64
	Votes    uint64
}
