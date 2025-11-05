package comment

import (
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	commentmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/mock"
)

func TestHandler_CreateCommentHandler_Success(t *testing.T) {
	usecase := new(commentmock.CommentUsecaseMock)

	input := dto.CommentDto{Content: "Hello world"}
	expected := dto.CommentDto{
		Content: "Hello world",
		Id:      uuid.New(),
	}
	usecase.On("AddComment", mock.Anything, mock.MatchedBy(func(c dto.CommentDto) bool {
		return c.Content == input.Content
	})).Return(expected, nil)

	h := &Handler{Usecase: usecase}

	body := `{"content":"Hello world"}`
	req := httptest.NewRequest(http.MethodPost, "/comments", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h.CreateCommentHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "Hello world")

	usecase.AssertExpectations(t)
}
