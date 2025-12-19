package search_bar

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"net/http"
)

func (h *Handler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryText := r.URL.Query().Get("q")
	if queryText == "" {
		json.WriteError(w, http.StatusBadRequest, "query is required")
		return
	}

	users, err := h.Usecase.SearchUsers(ctx, queryText)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}
	output := h.mapUserToOutput(users)

	json.Write(w, http.StatusOK, output)
}
