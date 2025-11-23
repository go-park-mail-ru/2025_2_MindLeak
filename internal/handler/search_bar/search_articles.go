package search_bar

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"net/http"
)

func (h *Handler) SearchArticles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryText := r.URL.Query().Get("q")
	if queryText == "" {
		json.WriteError(w, http.StatusBadRequest, "query is required")
		return
	}

	articles, err := h.Usecase.SearchArticles(ctx, queryText)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	output := h.mapArticleToOutput(articles)

	json.Write(w, http.StatusOK, output)

}
