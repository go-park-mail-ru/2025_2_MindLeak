package topBlogs

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	topblogsmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/topBlogs/mock"
)

func TestHandler_ShowTopBlogs_Success(t *testing.T) {
	usecase := new(topblogsmock.TopBlogsUsecaseMock)

	expectedBlogs := []models.User{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	usecase.On("ShowTopBlogs", mock.Anything).
		Return(expectedBlogs, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodGet, "/topBlogs", nil)
	rr := httptest.NewRecorder()

	h.ShowTopBlogs(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)

	var resp struct {
		Blogs []models.User `json:"blogs"`
	}
	err := json.NewDecoder(rr.Body).Decode(&resp)
	require.NoError(t, err)
	require.Len(t, resp.Blogs, 2)
	require.Equal(t, "Alice", resp.Blogs[0].Name)
	require.Equal(t, "Bob", resp.Blogs[1].Name)

	usecase.AssertExpectations(t)
}
