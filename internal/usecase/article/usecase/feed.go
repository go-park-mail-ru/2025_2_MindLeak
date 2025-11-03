package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/dto"
)

func (u *Usecase) Feed(ctx context.Context, feed models.Feed) (dto.ReceivedFeedDTO, error) {
	output, err := u.articleRepo.GetFeedArticles(ctx, feed)
	if err != nil {
		return dto.ReceivedFeedDTO{}, fmt.Errorf("get articles: %w", err)
	}

	return toDTO(output), nil
}

func toDTO(articles []*models.Article) dto.ReceivedFeedDTO {
	result := make([]models.Article, len(articles))
	for i, a := range articles {
		result[i] = models.Article{
			ID:           a.ID,
			AuthorID:     a.AuthorID,
			Title:        a.Title,
			Content:      a.Content,
			CreatedAt:    a.CreatedAt,
			ImageURL:     a.ImageURL,
			AuthorName:   a.AuthorName,
			AuthorAvatar: a.AuthorAvatar,
		}
	}
	return dto.ReceivedFeedDTO{Articles: result}
}
