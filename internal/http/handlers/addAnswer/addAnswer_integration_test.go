package addAnswer_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	mock_answer "survey/internal/db/repository/answerRepo/mocks"
	"survey/internal/entities"
	"survey/internal/http/handlers/addAnswer"
	"survey/internal/usecases/answerUC"
	"survey/pkg/logger"
	"testing"
)

func TestAddAnswer(t *testing.T) {
	log := logger.New()
	ctx := context.Background()
	ans := &entities.Answer{
		Text:     "Test Ans",
		SurveyID: 1,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_answer.NewMockAnswerRepo(ctrl)

	exp := &entities.Answer{
		ID:       1,
		Text:     "Test Ans",
		SurveyID: 1,
		Votes:    0,
	}
	repo.EXPECT().CreateAnswer(ctx, log, ans).Return(exp, nil).Times(1)

	uCase := answerUC.NewUseCase(repo)
	h := addAnswer.New(uCase, log)

	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/add-ans",
		bytes.NewBuffer([]byte(`{"text": "Test Ans", "survey_id": 1}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := `{"ID":1,"Text":"Test Ans","SurveyID":1,"Votes":0}` + "\n"

	require.Equal(t, expected, string(data))
}

func TestAddAnswerBadJSON(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_answer.NewMockAnswerRepo(ctrl)
	uCase := answerUC.NewUseCase(repo)
	h := addAnswer.New(uCase, log)

	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/add-ans",
		bytes.NewBuffer([]byte(`{"text": "Test Ans", "survey_id": 1`)),
	)
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)
	require.Equal(t, "bad json(parsing): unexpected EOF\n", string(data))
}

func TestAddAnswerBadReq(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_answer.NewMockAnswerRepo(ctrl)
	uCase := answerUC.NewUseCase(repo)
	h := addAnswer.New(uCase, log)

	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/add-ans",
		bytes.NewBuffer([]byte(`{"text": "", "survey_id": 1}`)),
	)
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)
	require.Equal(t, "bad json(validating): invalid text\n", string(data))
}

func TestAddAnswerUCaseError(t *testing.T) {
	log := logger.New()
	ctx := context.Background()
	ans := &entities.Answer{
		Text:     "Test Ans",
		SurveyID: 1,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_answer.NewMockAnswerRepo(ctrl)

	repoErr := errors.New("can't create the answer, db is down")
	repo.EXPECT().CreateAnswer(ctx, log, ans).Return(nil, repoErr).Times(1)

	uCase := answerUC.NewUseCase(repo)
	h := addAnswer.New(uCase, log)

	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/add-ans",
		bytes.NewBuffer([]byte(`{"text": "Test Ans", "survey_id": 1}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "can't add a new answer: can't create the answer, db is down\n"

	require.Equal(t, expected, string(data))
}
