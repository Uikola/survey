package answerUC

import (
	"context"
	"github.com/sirupsen/logrus"
	"survey/internal/db/repository/answerRepo"
	"survey/internal/entities"
)

type UseCase struct {
	repo answerRepo.AnswerRepo
}

func NewUseCase(answerRepo answerRepo.AnswerRepo) *UseCase {
	return &UseCase{repo: answerRepo}
}

func (uc *UseCase) AddAnswer(ctx context.Context, log logrus.FieldLogger, answer *entities.Answer) (*entities.Answer, error) {
	return uc.repo.CreateAnswer(ctx, log, answer)
}

func (uc *UseCase) DeleteAnswer(ctx context.Context, log logrus.FieldLogger, answerID, questionID, surveyID uint64) error {
	return uc.repo.DeleteAnswer(ctx, log, answerID, questionID, surveyID)
}
