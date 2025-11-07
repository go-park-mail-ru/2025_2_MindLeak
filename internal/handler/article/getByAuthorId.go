package article

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func (h *Handler) GetArticlesByAuthorId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authorIDStr := r.URL.Query().Get("author_id")
	if authorIDStr == "" {
		json.WriteError(w, http.StatusBadRequest, "author_id is required")
		return
	}

	authorID, err := uuid.Parse(authorIDStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid author_id")
		return
	}

	articles, err := h.Usecase.GetArticlesByAuthorID(ctx, authorID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	var out []dto.ArticleOutput
	for _, a := range articles {
		out = append(out, dto.ArticleOutput{
			ID:           a.ID,
			Title:        a.Title,
			Content:      a.Content,
			MediaURL:     a.MediaURL,
			TopicID:      a.TopicID,
			Status:       a.Status,
			AuthorName:   a.AuthorName,
			AuthorAvatar: a.AuthorAvatar,
			Topic:        a.Topic,
		})
	}

	json.Write(w, http.StatusOK, out)
}
