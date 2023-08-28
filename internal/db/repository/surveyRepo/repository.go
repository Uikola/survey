package surveyRepo

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"survey/internal/entities"
)

type Repository struct {
	db *sql.DB
}

type SurveyRepo interface {
	CreateSurvey(ctx context.Context, log logrus.FieldLogger, survey *entities.Survey) (*entities.Survey, error)
	DeleteSurvey(ctx context.Context, logger logrus.FieldLogger, surveyID uint64) error
}

func NewSurveyRepo(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateSurvey(ctx context.Context, log logrus.FieldLogger, survey *entities.Survey) (*entities.Survey, error) {
	query := `
	INSERT INTO surveys(title)
	VALUES ($1)
	RETURNING id, title`

	rows, err := r.db.QueryContext(ctx, query, survey.Title)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	var id uint64
	var title string

	for rows.Next() {
		err = rows.Scan(&id, &title)
		if err != nil {
			log.Errorf("error while scanning: %s", err.Error())
			return nil, err
		}
	}
	if rows.Err() != nil {
		log.Errorf("rows error: %s", rows.Err())
		return nil, rows.Err()
	}

	return &entities.Survey{
		ID:    id,
		Title: title,
	}, nil
}

func (r *Repository) DeleteSurvey(ctx context.Context, log logrus.FieldLogger, surveyID uint64) error {
	query := `
	DELETE FROM surveys
	WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, surveyID)
	if err != nil {
		log.Errorf("query execution error: %s", err.Error())
	}

	return nil
}
