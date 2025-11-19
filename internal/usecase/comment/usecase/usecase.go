package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/comment"
	repoComment "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/comment"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
)

var (
	CommentFetchingError   = errors.New("failed to fetch comments")
	CommentUpdatingError   = errors.New("failed to update comment")
	CommentDeletingError   = errors.New("failed to delete comment")
	CommentCreationError   = errors.New("failed to create comment")
	CommentForeignKeyError = errors.New("foreign key resolution fails")
	ServerError            = errors.New("internal server error")
)

type Usecase struct {
	commentRepo comment.CommentRepository
	sessionRepo session.SessionRepository
}

func NewCommentUsecase(commentRepo comment.CommentRepository, sessionRepo session.SessionRepository) *Usecase {
	return &Usecase{commentRepo: commentRepo, sessionRepo: sessionRepo}
}

func (u *Usecase) mapToDto(model models.Comment) dto.CommentDto {
	return dto.CommentDto{
		Id:           model.Id,
		Content:      model.Content,
		ArticleId:    model.ArticleId,
		UserId:       model.UserId,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		ReplyTo:      model.ReplyTo,
		AuthorName:   model.AuthorName,
		AuthorAvatar: model.AuthorAvatar,
		ArticleTitle: model.ArticleTitle,
	}
}

func (u *Usecase) mapToModel(commentDto dto.CommentDto) models.Comment {
	return models.Comment{
		Id:           commentDto.Id,
		Content:      commentDto.Content,
		ArticleId:    commentDto.ArticleId,
		UserId:       commentDto.UserId,
		CreatedAt:    commentDto.CreatedAt,
		UpdatedAt:    commentDto.UpdatedAt,
		ReplyTo:      commentDto.ReplyTo,
		AuthorName:   commentDto.AuthorName,
		AuthorAvatar: commentDto.AuthorAvatar,
		ArticleTitle: commentDto.ArticleTitle,
	}
}

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, repoComment.ErrGettingComment):
		return CommentFetchingError
	case errors.Is(err, repoComment.ErrUpdatingComment):
		return CommentUpdatingError
	case errors.Is(err, repoComment.ErrDeletingComment):
		return CommentDeletingError
	case errors.Is(err, repoComment.ErrResolvingArticleTitle):
		return CommentForeignKeyError
	case errors.Is(err, repoComment.ErrResolvingAuthor):
		return CommentForeignKeyError
	case errors.Is(err, repoComment.ErrCreatingComment):
		return CommentCreationError
	default:
		return ServerError
	}
}
