package handler

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

// This is a test endpoint just to give user a guest cookie
func FeedHandler(w http.ResponseWriter, sessions *repository.InMemorySession) {
	session, err := sessions.CreateSession()
	if err != nil {
		///
	}
	cookies.SetCookie(w, session.SessionId)
	json.Write(w, http.StatusOK, map[string]string{"session_id": session.SessionId.String()})
}
