package categories

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/categories"
	"net/http"
)

type Handler struct {
	Usecase categories.Usecase
}

func NewCategoriesHandler(u categories.Usecase) *Handler {
	return &Handler{Usecase: u}
}

func (h *Handler) handleError(err error) (int, string) {
	switch {
	case errors.Is(err, usecase.ErrArticleNotFound):
		return http.StatusNotFound, "Article Not Found"
	case errors.Is(err, usecase.ErrArticleExists):
		return http.StatusConflict, "Article Exists"
	case errors.Is(err, usecase.ErrServerError):
		return http.StatusInternalServerError, err.Error()
	default:
		return http.StatusInternalServerError, err.Error()
	}
}
