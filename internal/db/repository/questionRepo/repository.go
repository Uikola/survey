package questionRepo

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"survey/internal/entities"
)

type Repository struct {
	db *sql.DB
}

type QuestionRepo interface {
	CreateQuestion(ctx context.Context, log logrus.FieldLogger, question *entities.Question) (*entities.Question, error)
	DeleteQuestion(ctx context.Context, log logrus.FieldLogger, surveyID, questionID uint64) error
}

func NewQuestionRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateQuestion(ctx context.Context, log logrus.FieldLogger, question *entities.Question) (*entities.Question, error) {
	query := `
	INSERT INTO questions(text, survey_id)
	VALUES ($1, $2)
	RETURNING id, text, survey_id`

	rows, err := r.db.QueryContext(ctx, query, question.Text, question.SurveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	var id, surveyID uint64
	var text string

	for rows.Next() {
		err = rows.Scan(&id, &text, &surveyID)
		if err != nil {
			log.Errorf("error while scanning: %s", err.Error())
			return nil, err
		}
	}
	if rows.Err() != nil {
		log.Errorf("rows error: %s", rows.Err())
		return nil, rows.Err()
	}

	return &entities.Question{
		ID:       id,
		Text:     text,
		SurveyID: surveyID,
	}, nil
}

func (r *Repository) DeleteQuestion(ctx context.Context, log logrus.FieldLogger, surveyID, questionID uint64) error {
	query := `
	DELETE FROM questions
	WHERE id = $1 AND survey_id = $2`

	_, err := r.db.ExecContext(ctx, query, questionID, surveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return err
	}

	return nil
}
