package comment

import (
	"errors"
	"net/http"

	dto2 "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/comment/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/usecase"
)

type Handler struct {
	Usecase comment.Usecase
}

func NewCommentHandler(u comment.Usecase) *Handler {
	return &Handler{u}
}

func (h *Handler) mapToIODto(dto dto.CommentDto) dto2.CommentIODto {
	return dto2.CommentIODto{
		Id:           dto.Id,
		ArticleId:    dto.ArticleId,
		UserId:       dto.UserId,
		Content:      dto.Content,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		ReplyTo:      dto.ReplyTo,
		AuthorName:   dto.AuthorName,
		AuthorAvatar: dto.AuthorAvatar,
		ArticleTitle: dto.ArticleTitle,
	}
}

func (h *Handler) mapToDto(ioDto dto2.CommentIODto) dto.CommentDto {
	return dto.CommentDto{
		Id:           ioDto.Id,
		ArticleId:    ioDto.ArticleId,
		UserId:       ioDto.UserId,
		Content:      ioDto.Content,
		CreatedAt:    ioDto.CreatedAt,
		UpdatedAt:    ioDto.UpdatedAt,
		ReplyTo:      ioDto.ReplyTo,
		AuthorName:   ioDto.AuthorName,
		AuthorAvatar: ioDto.AuthorAvatar,
		ArticleTitle: ioDto.ArticleTitle,
	}
}

func (h *Handler) mapSliceToOutputDto(slice []dto.CommentDto) dto2.CommentsOutputDto {
	var ioSlice []dto2.CommentIODto

	for _, dtoObj := range slice {
		mapped := h.mapToIODto(dtoObj)
		ioSlice = append(ioSlice, mapped)
	}

	return dto2.CommentsOutputDto{
		Comments: ioSlice,
	}
}

func (h *Handler) handleError(err error) (int, string) {
	switch {
	case errors.Is(err, usecase.CommentForeignKeyError):
		return http.StatusInternalServerError, "internal apiserver error"
	case errors.Is(err, usecase.ServerError):
		return http.StatusInternalServerError, "internal apiserver error"
	case errors.Is(err, usecase.CommentCreationError):
		return http.StatusInternalServerError, "internal apiserver error"
	case errors.Is(err, usecase.CommentDeletingError):
		return http.StatusInternalServerError, "internal apiserver error"
	case errors.Is(err, usecase.CommentUpdatingError):
		return http.StatusInternalServerError, "internal apiserver error"
	case errors.Is(err, usecase.CommentFetchingError):
		return http.StatusInternalServerError, "internal apiserver error"
	default:
		return http.StatusInternalServerError, "unexpected error"
	}
}
