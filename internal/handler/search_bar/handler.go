package search_bar

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/search_bar/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/search_bar"
	"net/http"
	"strings"
)

type Handler struct {
	Usecase search_bar.Usecase
}

func NewSearchBarHandler(usecase search_bar.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

func (h *Handler) handleError(err error) (int, string) {

	switch {
	case errors.Is(err, user.ErrUserNotFound):
		return http.StatusNotFound, "user not found"

	case strings.Contains(err.Error(), "failed to get user"):
		return http.StatusInternalServerError, "failed to get user"

	case strings.Contains(err.Error(), "failed to create user"):
		return http.StatusBadRequest, "failed to create user"

	case errors.Is(err, article.ErrArticleNotFound):
		return http.StatusNotFound, "article not found"

	default:
		return http.StatusInternalServerError, "internal server error"
	}
}

func (h *Handler) mapUserToDto(user models.User) dto.UserDto {
	return dto.UserDto{
		Id:          user.Id,
		Name:        user.Name,
		Avatar:      user.Avatar,
		Subscribers: user.Subscribers,
	}
}

func (h *Handler) mapUserToOutput(users []models.User) dto.UsersDto {
	dtoSlice := make([]dto.UserDto, 0, len(users))

	for _, user := range users {
		dtoSlice = append(dtoSlice, h.mapUserToDto(user))
	}

	return dto.UsersDto{Users: dtoSlice}
}

func (h *Handler) mapArticleToDto(article models.Article) dto.ArticleDto {
	return dto.ArticleDto{
		ID:            article.ID,
		AuthorID:      article.AuthorID,
		Title:         article.Title,
		Content:       article.Content,
		MediaURL:      article.MediaURL,
		TopicID:       article.TopicID,
		Status:        article.Status,
		CommentsCount: article.CommentsCount,
		RepostsCount:  article.RepostsCount,
		ViewsCount:    article.ViewsCount,
		Topic:         article.Topic,
		AuthorName:    article.AuthorName,
		AuthorAvatar:  article.AuthorAvatar,
	}
}

func (h *Handler) mapArticleToOutput(articles []models.Article) dto.ArticlesDto {
	dtoSlice := make([]dto.ArticleDto, 0, len(articles))

	for _, article := range articles {
		dtoSlice = append(dtoSlice, h.mapArticleToDto(article))
	}

	return dto.ArticlesDto{Articles: dtoSlice}
}
