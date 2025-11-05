package categories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	catmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/categories/mock"
)

func TestHandler_CategoriesHandler_Success(t *testing.T) {
	usecase := new(catmock.CategoriesUsecaseMock)

	expectedArticles := []*models.Article{
		{Title: "First post"},
		{Title: "Second post"},
	}

	usecase.On("GetFeedByTopic", mock.Anything, "music", 0).
		Return(expectedArticles, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodGet, "/categories?topic=music&offset=0", nil)
	rr := httptest.NewRecorder()

	h.CategoriesHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)

	var got []*models.Article
	err := json.NewDecoder(rr.Body).Decode(&got)
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, "First post", got[0].Title)
	require.Equal(t, "Second post", got[1].Title)

	usecase.AssertExpectations(t)
}
