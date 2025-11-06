package usecase

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) UpdateArticle(
	ctx context.Context,
	sessionID uuid.UUID,
	articleID uuid.UUID,
	title, content, status *string,
	topicID *uuid.UUID,
	file *multipart.FileHeader,
) (models.Article, error) {

	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, "session not found: %v", err)
		return models.Article{}, u.handleError(err)
	}

	authorID := session.UserId

	old, err := u.articleRepo.GetArticleById(ctx, articleID)
	if err != nil {
		logger.Error(ctx, "article not found: %v", err)
		return models.Article{}, u.handleError(err)
	}

	if old.AuthorID != authorID {
		return models.Article{}, u.handleError(fmt.Errorf("not article owner"))
	}

	changed := false
	oldMediaURL := old.MediaURL

	if title != nil && *title != old.Title {
		old.Title = *title
		changed = true
	}
	if content != nil && *content != old.Content {
		old.Content = *content
		changed = true
	}
	if status != nil && *status != old.Status {
		old.Status = *status
		changed = true
	}
	if topicID != nil && *topicID != old.TopicID {
		old.TopicID = *topicID
		changed = true
	}

	if file != nil {
		f, err := file.Open()
		if err != nil {
			return models.Article{}, u.handleError(err)
		}
		defer f.Close()

		newURL, err := u.minioClient.UploadMedia(ctx, old.ID, f, file)
		if err != nil {
			return models.Article{}, u.handleError(err)
		}

		old.MediaURL = newURL
		changed = true
	}

	if changed {
		updated, err := u.articleRepo.UpdateArticle(ctx, old)
		if err != nil {
			logger.Error(ctx, "update failed: %v", err)
			if file != nil {
				u.minioClient.DeleteMedia(ctx, old.ID)
			}
			return models.Article{}, u.handleError(err)
		}
		old = updated
	}

	if oldMediaURL != "" && old.MediaURL != oldMediaURL {
		u.minioClient.DeleteMedia(ctx, old.ID)
	}

	return old, nil
}
