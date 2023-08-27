package surveyUC

import (
	"context"
	"github.com/sirupsen/logrus"
	"survey/internal/db/repository/surveyRepo"
	"survey/internal/entities"
)

type UseCase struct {
	repo surveyRepo.SurveyRepo
}

func NewUseCase(surveyRepo surveyRepo.SurveyRepo) *UseCase {
	return &UseCase{repo: surveyRepo}
}

func (uc *UseCase) CreateSurvey(ctx context.Context, log logrus.FieldLogger, survey *entities.Survey) (*entities.Survey, error) {
	return uc.repo.CreateSurvey(ctx, log, survey)
}

func (uc *UseCase) DeleteSurvey(ctx context.Context, log logrus.FieldLogger, surveyID uint64) error {
	return uc.repo.DeleteSurvey(ctx, log, surveyID)
}
