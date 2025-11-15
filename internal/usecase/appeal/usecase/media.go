package usecase

import (
	"context"
	"mime/multipart"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

func (u *Usecase) UploadAttachment(ctx context.Context, file multipart.File, header *multipart.FileHeader) (models.Appeal, error) {

}
