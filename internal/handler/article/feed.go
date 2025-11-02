package article

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	usecaseDTO "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/dto"
	"github.com/gorilla/schema"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

var decoder = schema.NewDecoder()

func (h *Handler) Feed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger.Info(ctx, "[article.Feed] handler start", nil)

	decoder.IgnoreUnknownKeys(true)

	feedInputDTO := &dto.FeedInputDTO{}
	if err := decoder.Decode(feedInputDTO, r.URL.Query()); err != nil {
		logger.Error(ctx, "[article.Feed] failed to decode query params: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	logger.Info(ctx, "[article.Feed] decoded params: offset=%d", feedInputDTO.Offset)

	feedEntity := &models.Feed{
		Offset: feedInputDTO.Offset,
	}

	output, err := h.Usecase.Feed(ctx, *feedEntity)
	if err != nil {
		logger.Error(ctx, "[article.Feed] usecase.Feed error: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	logger.Info(ctx, "[article.Feed] usecase.Feed success, preparing response", nil)

	feedOutputDto := toOutputDTO(output)

	if err = json.Write(w, http.StatusOK, feedOutputDto); err != nil {
		logger.Error(ctx, "[article.Feed] failed to write response: %v", err)
		return
	}

	logger.Info(ctx, "[article.Feed] handler finished successfully", nil)
}

func toOutputDTO(usecaseDto usecaseDTO.ReceivedFeedDTO) dto.FeedOutputDTO {
	result := make([]dto.ArticleOutputDTO, len(usecaseDto.Articles))
	for i, a := range usecaseDto.Articles {
		result[i] = dto.ArticleOutputDTO{
			Id:           a.ID,
			AuthorId:     a.AuthorID,
			Title:        a.Title,
			Content:      a.Content,
			CreatedAt:    a.PublishedAt,
			Image:        a.ImageURL,
			AuthorName:   a.AuthorName,
			AuthorAvatar: a.AuthorAvatar,
		}
	}
	return dto.FeedOutputDTO{Articles: result}
}
