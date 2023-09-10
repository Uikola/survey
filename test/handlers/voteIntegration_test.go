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
	"survey/internal/http/handlers/vote"
	"survey/internal/usecases/answerUC"
	"survey/pkg/logger"
	mock_answer "survey/test/mocks"
	"testing"
)

func TestVote(t *testing.T) {
	log := logger.New()
	ctx := context.Background()
	answerID := uint64(1)
	surveyID := uint64(1)

	ctrl := gomock.NewController(t)

	repo := mock_answer.NewMockAnswerRepo(ctrl)

	repo.EXPECT().Vote(ctx, log, answerID, surveyID).Return(nil).Times(1)

	uCase := answerUC.NewUseCase(repo)
	h := vote.New(uCase, log)
	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/vote",
		bytes.NewBuffer([]byte(`{"answer_id":1,"survey_id":1}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := `{"Message":"vote counted successfully"}` + "\n"

	require.Equal(t, expected, string(data))

}

func TestVoteBadJSON(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)

	repo := mock_answer.NewMockAnswerRepo(ctrl)
	uCase := answerUC.NewUseCase(repo)
	h := vote.New(uCase, log)
	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/vote",
		bytes.NewBuffer([]byte(`{"answer_id":1,"survey_id":1`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "bad json: unexpected EOF\n"

	require.Equal(t, expected, string(data))
}

func TestVoteBadReq(t *testing.T) {
	log := logger.New()

	ctrl := gomock.NewController(t)

	repo := mock_answer.NewMockAnswerRepo(ctrl)
	uCase := answerUC.NewUseCase(repo)
	h := vote.New(uCase, log)
	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/vote",
		bytes.NewBuffer([]byte(`{"answer_id":1,"survey_id":0}`)),
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

func TestVoteUCError(t *testing.T) {
	log := logger.New()
	ctx := context.Background()
	answerID := uint64(1)
	surveyID := uint64(1)

	ctrl := gomock.NewController(t)

	repo := mock_answer.NewMockAnswerRepo(ctrl)

	repoErr := errors.New("db is down")
	repo.EXPECT().Vote(ctx, log, answerID, surveyID).Return(repoErr).Times(1)

	uCase := answerUC.NewUseCase(repo)
	h := vote.New(uCase, log)
	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/vote",
		bytes.NewBuffer([]byte(`{"answer_id":1,"survey_id":1}`)),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := "voting error: db is down\n"

	require.Equal(t, expected, string(data))

}
