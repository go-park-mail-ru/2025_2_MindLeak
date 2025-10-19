package auth

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cookies.DeleteCookie(w, r)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	sessionId, err := uuid.Parse(cookie.Value)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	flag, err := h.Usecase.Logout(ctx, sessionId)
	if err != nil {
		///
	}
	if flag {
		json.Write(w, http.StatusOK, map[string]string{"messaage": "logged out"})
		return
	}
}
