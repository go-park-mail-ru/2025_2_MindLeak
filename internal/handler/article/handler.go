package article

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
)

type Handler struct {
	Usecase article.Usecase
}

func NewHandler(u article.Usecase) *Handler {
	return &Handler{Usecase: u}
}

func (h *Handler) handleError(err error) (int, string) {
	switch {
	case errors.Is(err, usecase.ErrServerError):
		return http.StatusInternalServerError, "internal server error"
	case errors.Is(err, usecase.ErrArticleExists):
		return http.StatusInternalServerError, "internal server error"
	case errors.Is(err, usecase.ErrArticleNotFound):
		return http.StatusInternalServerError, "internal server error"
	default:
		return http.StatusInternalServerError, "unexpected error"
	}
}
