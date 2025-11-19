package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
)

func (u *Usecase) AddComment(ctx context.Context, comment dto.CommentDto) (dto.CommentDto, error) {
	added, err := u.commentRepo.CreateComment(ctx, u.mapToModel(comment))
	if err != nil {
		logger.Error(ctx, err.Error())
		return dto.CommentDto{}, err
	}

	return u.mapToDto(added), nil
}
