package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
)

func (u *Usecase) UpdateComment(ctx context.Context, commentDto dto.CommentDto) (dto.CommentDto, error) {
	updated, err := u.commentRepo.UpdateComment(ctx, u.mapToModel(commentDto))

	if err != nil {
		logger.Error(ctx, err.Error())
		return dto.CommentDto{}, u.handleError(err)
	}

	return u.mapToDto(updated), nil
}
