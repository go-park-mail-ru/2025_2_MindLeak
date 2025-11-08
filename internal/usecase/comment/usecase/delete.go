package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) DeleteComment(ctx context.Context, commentId uuid.UUID) (bool, error) {
	flag, err := u.commentRepo.DeleteComment(ctx, commentId)
	if err != nil {
		logger.Error(ctx, err.Error())
		return false, u.handleError(err)
	}
	return flag, nil
}
