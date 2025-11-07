package article

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article/dto"
	"github.com/google/uuid"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}
	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	input := dto.CreateArticleInput{}
	if err := json.Read(r, &input); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	article, err := h.Usecase.CreateArticle(ctx, sessionID, input.Title, input.Content, input.TopicID, nil)
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
		TopicID:      article.Topic.TopicId,
		Status:       article.Status,
		AuthorName:   article.AuthorName,
		AuthorAvatar: article.AuthorAvatar,
	}

	json.Write(w, http.StatusCreated, out)
}
