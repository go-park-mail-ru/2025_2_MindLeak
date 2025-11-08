package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) GetCommentsByArticle(ctx context.Context, articleId uuid.UUID) ([]dto.CommentDto, error) {
	comments, err := u.commentRepo.GetCommentsByArticle(ctx, articleId)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, u.handleError(err)
	}
	var commentDtos []dto.CommentDto

	for _, comment := range comments {
		commentDtos = append(commentDtos, u.mapToDto(comment))
	}
	return commentDtos, nil
}
