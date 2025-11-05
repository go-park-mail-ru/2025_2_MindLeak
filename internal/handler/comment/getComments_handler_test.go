package comment

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	commentmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/mock"
)

func TestHandler_GetCommentsHandler_ByArticle_Success(t *testing.T) {
	usecase := new(commentmock.CommentUsecaseMock)

	articleID := uuid.New()
	expected := []dto.CommentDto{
		{Id: uuid.New(), Content: "First comment"},
		{Id: uuid.New(), Content: "Second comment"},
	}

	usecase.On("GetCommentsByArticle", mock.Anything, articleID).
		Return(expected, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodGet, "/comments?articleId="+articleID.String(), nil)
	rr := httptest.NewRecorder()

	h.GetCommentsHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)

	var resp struct {
		Comments []dto.CommentDto `json:"comments"`
	}
	err := json.NewDecoder(rr.Body).Decode(&resp)
	require.NoError(t, err)
	require.Len(t, resp.Comments, 2)
	require.Equal(t, "First comment", resp.Comments[0].Content)
	require.Equal(t, "Second comment", resp.Comments[1].Content)

	usecase.AssertExpectations(t)
}
