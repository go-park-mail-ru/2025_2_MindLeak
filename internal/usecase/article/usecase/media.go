package usecase

import (
	"context"
	"mime/multipart"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

// UploadArticleMedia — загружает медиа к статье (с откатом и удалением старой)
func (u *Usecase) UploadArticleMedia(
	ctx context.Context,
	articleID uuid.UUID,
	file multipart.File,
	header *multipart.FileHeader,
) (models.Article, error) {

	// 1. Получаем текущую статью
	article, err := u.articleRepo.GetArticleById(ctx, articleID)
	if err != nil {
		logger.Error(ctx, "get article failed: %v", err)
		return models.Article{}, u.handleError(err)
	}

	// 2. Сохраняем старую медиа
	oldMediaURL := article.MediaURL

	// 3. Загружаем новую
	newURL, err := u.minioClient.UploadMedia(ctx, articleID, file, header)
	if err != nil {
		logger.Error(ctx, "upload article media failed: %v", err)
		return models.Article{}, u.handleError(err)
	}

	// 4. Обновляем URL
	article.MediaURL = newURL

	// 5. Сохраняем в БД
	updated, err := u.articleRepo.UpdateArticle(ctx, article)
	if err != nil {
		logger.Error(ctx, "update article failed: %v", err)
		// Откат: удаляем новую медиа
		u.minioClient.DeleteMedia(ctx, articleID)
		return models.Article{}, u.handleError(err)
	}

	// 6. Удаляем старую медиа (если была и не дефолтная)
	if oldMediaURL != "" && oldMediaURL != u.minioClient.GetDefaultMedia() {
		u.minioClient.DeleteMedia(ctx, articleID)
	}

	return updated, nil
}

// DeleteArticleMedia — удаляет медиа и ставит дефолтную картинку
func (u *Usecase) DeleteArticleMedia(ctx context.Context, articleID uuid.UUID) (models.Article, error) {
	article, err := u.articleRepo.GetArticleById(ctx, articleID)
	if err != nil {
		logger.Error(ctx, "get article failed: %v", err)
		return models.Article{}, u.handleError(err)
	}

	// Удаляем из MinIO (если не дефолт)
	if article.MediaURL != "" && article.MediaURL != u.minioClient.GetDefaultMedia() {
		u.minioClient.DeleteMedia(ctx, articleID)
	}

	// Ставим дефолт
	article.MediaURL = u.minioClient.GetDefaultMedia()

	updated, err := u.articleRepo.UpdateArticle(ctx, article)
	if err != nil {
		logger.Error(ctx, "update article failed: %v", err)
		return models.Article{}, u.handleError(err)
	}

	return updated, nil
}
