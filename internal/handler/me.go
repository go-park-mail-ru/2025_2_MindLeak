package handler

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func MeHandler(w http.ResponseWriter, r *http.Request, sessions *session.InMemorySession, users *repository.InMemoryUser) {
	cookie, err := cookies.GetCookie(r)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session, err := sessions.GetSessionById(sessionID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := users.GetUserById(session.UserId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = json.Write(w, http.StatusOK, user)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
}
