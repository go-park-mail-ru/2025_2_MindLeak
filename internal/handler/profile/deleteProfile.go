package profile

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
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
	Session, err := h.Usecase.GetSession(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, "GetSession: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
	}

	flag, err := h.Usecase.DeleteProfile(ctx, Session.UserId)
	if err != nil {
		logger.Error(ctx, "DeleteProfile: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
	}
	if flag {
		json.Write(w, http.StatusOK, map[string]string{"message": "logged out"})
		return
	}
}
