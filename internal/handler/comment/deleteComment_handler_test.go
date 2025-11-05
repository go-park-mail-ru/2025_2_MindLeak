package comment

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	commentmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/mock"
)

func TestHandler_DeleteCommentHandler_Success(t *testing.T) {
	usecase := new(commentmock.CommentUsecaseMock)

	commentID := uuid.New()
	usecase.On("DeleteComment", mock.Anything, commentID).
		Return(true, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodDelete, "/comment?id="+commentID.String(), nil)
	rr := httptest.NewRecorder()

	h.DeleteCommentHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "true")

	usecase.AssertExpectations(t)
}
