package test

import (
	"bytes"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"survey/internal/entities"
	"survey/internal/http/handlers/startSurvey"
	"survey/internal/usecases/surveyUC"
	"survey/pkg/logger"
	mock_surveyRepo "survey/test/mocks"
	"testing"
)

func TestStartSurvey(t *testing.T) {
	log := logger.New()
	ctx := context.Background()
	survey := &entities.Survey{
		Title: "Test Survey",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_surveyRepo.NewMockSurveyRepo(ctrl)

	exp := &entities.Survey{
		ID:      1,
		Title:   "Test Survey",
		Answers: nil,
	}
	repo.EXPECT().CreateSurvey(ctx, log, survey).Return(exp, nil).Times(1)

	uCase := surveyUC.NewUseCase(repo)
	h := startSurvey.New(uCase, log)
	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/start-survey",
		bytes.NewBuffer([]byte(`{"title": "Test Survey"}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := `{"ID":1,"Title":"Test Survey","Answers":null}` + "\n"

	require.Equal(t, expected, string(data))
}

func TestStartSurveyBadJSON(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_surveyRepo.NewMockSurveyRepo(ctrl)

	uCase := surveyUC.NewUseCase(repo)

	h := startSurvey.New(uCase, log)

	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/start-survey",
		bytes.NewBuffer([]byte(`{"title": "Test Survey"`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "bad json(parsing):unexpected EOF\n"

	require.Equal(t, expected, string(data))
}

func TestStartSurveyBadReq(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_surveyRepo.NewMockSurveyRepo(ctrl)

	uCase := surveyUC.NewUseCase(repo)

	h := startSurvey.New(uCase, log)

	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/start-survey",
		bytes.NewBuffer([]byte(`{"title": ""}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "bad json(validating):invalid title\n"

	require.Equal(t, expected, string(data))
}

func TestStartSurveyUCError(t *testing.T) {
	log := logger.New()
	ctx := context.Background()
	survey := &entities.Survey{
		Title: "Test Survey",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_surveyRepo.NewMockSurveyRepo(ctrl)

	repoErr := errors.New("can't create the survey, db is down")
	repo.EXPECT().CreateSurvey(ctx, log, survey).Return(nil, repoErr).Times(1)

	uCase := surveyUC.NewUseCase(repo)
	h := startSurvey.New(uCase, log)
	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/start-survey",
		bytes.NewBuffer([]byte(`{"title": "Test Survey"}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "create survey err:can't create the survey, db is down\n"

	require.Equal(t, expected, string(data))
}
