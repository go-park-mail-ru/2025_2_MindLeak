package article

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func (h *Handler) GetOwnArticles(w http.ResponseWriter, r *http.Request) {
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

	articles, err := h.Usecase.GetArticlesByAuthorId(ctx, sessionID)
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
		})
	}

	json.Write(w, http.StatusOK, out)
}
