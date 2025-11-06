package article

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
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

	ok, err := h.Usecase.DeleteArticle(ctx, sessionID, articleID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}
	if !ok {
		json.WriteError(w, http.StatusNotFound, "article not found")
		return
	}

	json.Write(w, http.StatusOK, map[string]string{"message": "article deleted"})
}
