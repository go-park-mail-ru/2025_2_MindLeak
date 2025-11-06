package article

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// internal/handler/article/update.go
func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	articleID, err := uuid.Parse(vars["id"])
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid article id")
		return
	}

	input := dto.UpdateArticleInput{}
	if err := json.Read(r, &input); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	article, err := h.Usecase.UpdateArticle(ctx, sessionID, articleID, input.Title, input.Content, input.Status, input.TopicID, nil)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	out := dto.ArticleOutput{
		ID:           article.ID,
		Title:        article.Title,
		Content:      article.Content,
		MediaURL:     article.MediaURL,
		TopicID:      article.TopicID,
		Status:       article.Status,
		AuthorName:   article.AuthorName,
		AuthorAvatar: article.AuthorAvatar,
	}

	json.Write(w, http.StatusOK, out)
}
