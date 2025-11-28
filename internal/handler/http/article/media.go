// internal/handler/article/media.go
package article

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/article/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// --- ID статьи ---
	vars := mux.Vars(r)
	articleID, err := uuid.Parse(vars["id"])
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid article id")
		return
	}

	// --- Файл ---
	file, header, err := r.FormFile("file")
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "file not provided")
		return
	}
	defer file.Close()

	// --- Юзкейс ---
	article, err := h.Usecase.UploadArticleMedia(ctx, articleID, file, header)
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
		Topic:        article.Topic,
		Status:       article.Status,
		AuthorName:   article.AuthorName,
		AuthorAvatar: article.AuthorAvatar,
	}

	json.Write(w, http.StatusOK, out)
}

func (h *Handler) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	articleID, err := uuid.Parse(vars["id"])
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid article id")
		return
	}

	article, err := h.Usecase.DeleteArticleMedia(ctx, articleID)
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
		Topic:        article.Topic,
		Status:       article.Status,
		AuthorName:   article.AuthorName,
		AuthorAvatar: article.AuthorAvatar,
	}

	json.Write(w, http.StatusOK, out)
}
