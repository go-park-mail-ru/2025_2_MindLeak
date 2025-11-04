package topBlogs

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/topBlogs/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/topBlogs"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/topBlogs/usecase"
	"net/http"
)

type Handler struct {
	Usecase topBlogs.Usecase
}

func NewTopBlogsHandler(usecase topBlogs.Usecase) *Handler {
	return &Handler{Usecase: usecase}
}

func (h *Handler) mapToDto(blog models.User) dto.TopBlogDto {
	return dto.TopBlogDto{
		Name:        blog.Name,
		Avatar:      blog.Avatar,
		Subscribers: blog.Subscribers,
	}
}

func (h *Handler) mapToOutputSlice(blogs []models.User) dto.TopBlogsDto {
	var dtoSlice []dto.TopBlogDto

	for _, blog := range blogs {
		dtoSlice = append(dtoSlice, h.mapToDto(blog))
	}

	return dto.TopBlogsDto{
		Blogs: dtoSlice,
	}
}

func (h *Handler) handleError(err error) (int, string) {
	switch {
	case errors.Is(err, usecase.NoTopBlogs):
		return http.StatusBadRequest, "no top blogs"
	case errors.Is(err, usecase.QueryFailed):
		return http.StatusBadRequest, "query failed"
	case errors.Is(err, usecase.ScanFailed):
		return http.StatusBadRequest, "scan failed"
	case errors.Is(err, usecase.ServerError):
		return http.StatusInternalServerError, "server error"
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
