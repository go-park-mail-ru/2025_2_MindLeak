package usecase

import (
	"context"
	"mime/multipart"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) CreateArticle(
	ctx context.Context,
	sessionID uuid.UUID,
	title, content string,
	topicID int,
	file *multipart.FileHeader,
) (models.Article, error) {

	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Article{}, u.handleError(err)
	}

	authorID := session.UserId

	var mediaURL string
	if file != nil {
		f, err := file.Open()
		if err != nil {
			return models.Article{}, u.handleError(err)
		}
		defer f.Close()

		mediaURL, err = u.minioClient.UploadMedia(ctx, authorID, f, file)
		if err != nil {
			return models.Article{}, u.handleError(err)
		}
	}

	created, err := u.articleRepo.CreateArticle(ctx, authorID, title, content, topicID)
	if err != nil {
		logger.Error(ctx, err.Error())
		if mediaURL != "" {
			u.minioClient.DeleteMedia(ctx, authorID)
		}
		return models.Article{}, u.handleError(err)
	}

	if mediaURL != "" {
		created.MediaURL = mediaURL
		_, err = u.articleRepo.UpdateArticle(ctx, created)
		if err != nil {
			u.minioClient.DeleteMedia(ctx, authorID)
			return models.Article{}, u.handleError(err)
		}
	}

	return created, nil
}
