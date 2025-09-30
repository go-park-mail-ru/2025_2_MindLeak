package handler

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func FeedHandler(w http.ResponseWriter, r *http.Request, sessions *repository.InMemorySession, articles *repository.InMemoryArticle) {
	if r.Method != http.MethodGet {
		json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	cookie, err := cookies.GetCookie(r)
	if err == nil {
		if sessionID, parseErr := uuid.Parse(cookie.Value); parseErr == nil {
			if _, sessErr := sessions.GetSessionById(sessionID); sessErr == nil {
				returnFeed(w, articles)
				return
			}
		}
	}

	session, err := sessions.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	cookies.SetCookie(w, session.SessionId)

	returnFeed(w, articles)
}

func returnFeed(w http.ResponseWriter, articles *repository.InMemoryArticle) {
	mockArticles, err := articles.GetAllArticles()
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.Write(w, http.StatusOK, mockArticles); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
	}
}
