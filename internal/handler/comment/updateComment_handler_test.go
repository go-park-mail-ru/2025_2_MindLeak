package comment

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	commentmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/mock"
)

func TestHandler_UpdateCommentHandler_Success(t *testing.T) {
	usecase := new(commentmock.CommentUsecaseMock)

	input := dto.CommentDto{Id: uuid.New(), Content: "Old text"}
	expected := dto.CommentDto{Id: input.Id, Content: "Updated text"}

	usecase.On("UpdateComment", mock.Anything, mock.MatchedBy(func(c dto.CommentDto) bool {
		return c.Id == input.Id
	})).Return(expected, nil)

	h := &Handler{Usecase: usecase}

	body := `{"id":"` + input.Id.String() + `","content":"Updated text"}`
	req := httptest.NewRequest(http.MethodPut, "/comment", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h.UpdateCommentHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "Updated text")

	usecase.AssertExpectations(t)
}
