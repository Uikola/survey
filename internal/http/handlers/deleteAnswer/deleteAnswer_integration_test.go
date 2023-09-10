package deleteAnswer_test

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"survey/internal/http/handlers/deleteAnswer"
	"survey/internal/usecases/answerUC"
	"survey/pkg/logger"
	mock_answer "survey/test/mocks"
	"testing"
)

func TestDeleteAns(t *testing.T) {
	log := logger.New()
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_answer.NewMockAnswerRepo(ctrl)

	answerID := uint64(1)
	surveyID := uint64(1)

	repo.EXPECT().DeleteAnswer(ctx, log, answerID, surveyID).Return(nil).Times(1)

	uCase := answerUC.NewUseCase(repo)
	h := deleteAnswer.New(uCase, log)

	serverFunc := h.ServeHTTP

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/delete-ans?survey_id=1&ans_id=1",
		bytes.NewBuffer([]byte("")),
	)
	req.Header.Set("Content-Type", "Application/Json")
	serverFunc(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	require.NoError(t, err)

	expected := `"message":"answer deleted"` + "\n"

	require.Equal(t, expected, string(data))
}
