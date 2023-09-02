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

//go:generate mockgen -source=repository.go -destination=mocks/mock_repository.go
type AnswerRepo interface {
	CreateAnswer(ctx context.Context, log logrus.FieldLogger, answer *entities.Answer) (*entities.Answer, error)
	DeleteAnswer(ctx context.Context, log logrus.FieldLogger, answerID, surveyID uint64) error
	Vote(ctx context.Context, log logrus.FieldLogger, answerID, surveyID uint64) error
}

func NewAnswerRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateAnswer(ctx context.Context, log logrus.FieldLogger, answer *entities.Answer) (*entities.Answer, error) {
	query := `
	INSERT INTO answers(text, survey_id)
	VALUES ($1, $2)
	RETURNING id, text, survey_id, votes`

	rows, err := r.db.QueryContext(ctx, query, answer.Text, answer.SurveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	var id, surveyID, votes uint64
	var text string

	for rows.Next() {
		err = rows.Scan(&id, &text, &surveyID, &votes)
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
		ID:       id,
		Text:     text,
		SurveyID: surveyID,
		Votes:    votes,
	}, nil
}

func (r *Repository) DeleteAnswer(ctx context.Context, log logrus.FieldLogger, answerID, surveyID uint64) error {
	query := `
	DELETE FROM answers
	WHERE id = $1 AND survey_id = $2`

	_, err := r.db.ExecContext(ctx, query, answerID, surveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return err
	}
	return nil
}

func (r *Repository) Vote(ctx context.Context, log logrus.FieldLogger, answerID, surveyID uint64) error {
	query := `
	UPDATE answers
	SET votes = votes + 1
	WHERE id = $1 AND survey_id = $2`

	_, err := r.db.ExecContext(ctx, query, answerID, surveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return err
	}

	return nil
}
