package article

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		json.WriteError(w, http.StatusBadRequest, "missing article ID")
		return
	}

	articleID, err := uuid.Parse(idStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid article ID")
		return
	}

	article, err := h.Usecase.GetArticleById(ctx, articleID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	out := dto.ArticleOutput{
		ID:           article.ID,
		Title:        article.Title,
		AuthorID:     article.AuthorID,
		Content:      article.Content,
		MediaURL:     article.MediaURL,
		TopicID:      article.TopicID,
		Status:       article.Status,
		AuthorName:   article.AuthorName,
		AuthorAvatar: article.AuthorAvatar,
		Topic:        article.Topic,
	}

	json.Write(w, http.StatusOK, out)
}
