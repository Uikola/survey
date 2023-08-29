package questionUC

import (
	"context"
	"github.com/sirupsen/logrus"
	"survey/internal/db/repository/questionRepo"
	"survey/internal/entities"
)

type UseCase struct {
	repo questionRepo.QuestionRepo
}

func NewUseCase(repo questionRepo.QuestionRepo) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) AddQuestion(ctx context.Context, log logrus.FieldLogger, question *entities.Question) (*entities.Question, error) {
	return uc.repo.CreateQuestion(ctx, log, question)
}

func (uc *UseCase) DeleteQuestion(ctx context.Context, log logrus.FieldLogger, surveyID, questionID uint64) error {
	return uc.repo.DeleteQuestion(ctx, log, surveyID, questionID)
}
