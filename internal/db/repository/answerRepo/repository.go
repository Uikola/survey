package answerRepo

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"survey/internal/entities"
)

type Repository struct {
	db *sql.DB
}

type AnswerRepo interface {
	CreateAnswer(ctx context.Context, log logrus.FieldLogger, answer *entities.Answer) (*entities.Answer, error)
	DeleteAnswer(ctx context.Context, log logrus.FieldLogger, answerID, questionID, surveyID uint64) error
}

func NewAnswerRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateAnswer(ctx context.Context, log logrus.FieldLogger, answer *entities.Answer) (*entities.Answer, error) {
	query := `
	INSERT INTO answers(text, question_id, survey_id)
	VALUES ($1, $2, $3)
	RETURNING id, text, survey_id, question_id, votes`

	rows, err := r.db.QueryContext(ctx, query, answer.Text, answer.QuestionID, answer.SurveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	var id, questionID, surveyID, votes uint64
	var text string

	for rows.Next() {
		err = rows.Scan(&id, &text, &surveyID, &questionID, &votes)
		if err != nil {
			log.Errorf("error while scanning: %s", err.Error())
			return nil, err
		}
	}
	if rows.Err() != nil {
		log.Errorf("rows error: %s", rows.Err())
		return nil, rows.Err()
	}

	return &entities.Answer{
		ID:         id,
		Text:       text,
		SurveyID:   surveyID,
		QuestionID: questionID,
		Votes:      votes,
	}, nil
}

func (r *Repository) DeleteAnswer(ctx context.Context, log logrus.FieldLogger, answerID, questionID, surveyID uint64) error {
	query := `
	DELETE FROM answers
	WHERE id = $1 AND question_id = $2 AND survey_id = $3`

	_, err := r.db.ExecContext(ctx, query, answerID, questionID, surveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return err
	}
	return nil
}
