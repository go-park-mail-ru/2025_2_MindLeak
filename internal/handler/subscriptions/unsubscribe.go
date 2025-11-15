package subscriptions

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		logger.Error(ctx, "GetCookie: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		logger.Error(ctx, "Parse: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		json.WriteError(w, http.StatusBadRequest, "missing user id")
		return
	}

	targetID, err := uuid.Parse(idStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid uuid")
		return
	}

	flag, err := h.Usecase.Unsubscribe(ctx, targetID, sessionID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

	if flag {
		msg := map[string]string{
			"message": "user has been unsubscribed",
		}
		json.Write(w, http.StatusOK, msg)
		return
	}
}
