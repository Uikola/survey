package getResult_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	mock_survey "survey/internal/db/repository/surveyRepo/mocks"
	"survey/internal/entities"
	"survey/internal/http/handlers/getResult"
	"survey/internal/usecases/surveyUC"
	"survey/pkg/logger"
	"testing"
)

func TestGetResult(t *testing.T) {
	ctx := context.Background()
	log := logger.New()
	surveyID := uint64(1)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_survey.NewMockSurveyRepo(ctrl)

	exp := &entities.Survey{
		ID:    1,
		Title: "Test Title",
		Answers: []entities.Answer{
			{
				ID:       1,
				Text:     "test",
				SurveyID: 1,
				Votes:    5,
			},
			{
				ID:       2,
				Text:     "testTest",
				SurveyID: 1,
				Votes:    3,
			},
		},
	}

	repo.EXPECT().GetResult(ctx, log, surveyID).Return(exp, nil).Times(1)
	uCase := surveyUC.NewUseCase(repo)
	h := getResult.New(uCase, log)

	serverFunc := h.ServeHTTP
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/getResult",
		bytes.NewBuffer([]byte(`{"survey_id": 1}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := `{"ID":1,"Title":"Test Title","Answers":[{"ID":1,"Text":"test","SurveyID":1,"Votes":5},{"ID":2,"Text":"testTest","SurveyID":1,"Votes":3}]}` + "\n"
	require.Equal(t, expected, string(data))

}

func TestGetResultBadJSON(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_survey.NewMockSurveyRepo(ctrl)
	uCase := surveyUC.NewUseCase(repo)
	h := getResult.New(uCase, log)

	serverFunc := h.ServeHTTP
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/getResult",
		bytes.NewBuffer([]byte(`{"survey_id": 1`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "bad json(parsing): unexpected EOF\n"
	require.Equal(t, expected, string(data))

}

func TestGetResultBadReq(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_survey.NewMockSurveyRepo(ctrl)
	uCase := surveyUC.NewUseCase(repo)
	h := getResult.New(uCase, log)

	serverFunc := h.ServeHTTP
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/getResult",
		bytes.NewBuffer([]byte(`{"survey_id": 0}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "bad json(validating): invalid survey id\n"
	require.Equal(t, expected, string(data))

}

func TestGetResultUCError(t *testing.T) {
	ctx := context.Background()
	log := logger.New()
	surveyID := uint64(1)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_survey.NewMockSurveyRepo(ctrl)

	repoErr := errors.New("db is down")
	repo.EXPECT().GetResult(ctx, log, surveyID).Return(nil, repoErr).Times(1)
	uCase := surveyUC.NewUseCase(repo)
	h := getResult.New(uCase, log)

	serverFunc := h.ServeHTTP
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/getResult",
		bytes.NewBuffer([]byte(`{"survey_id": 1}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "can't get result: db is down\n"
	require.Equal(t, expected, string(data))

}
