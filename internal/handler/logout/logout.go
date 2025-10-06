package logout

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"

	"github.com/google/uuid"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, sessions session.SessionRepository) {
	if r.Method != http.MethodPost {
		json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

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

	flag, err := sessions.DeleteSessionById(sessionId)
	if flag {
		json.Write(w, http.StatusOK, map[string]string{
			"message": "logged out",
		})
		return
	}

	json.WriteError(w, http.StatusBadRequest, err.Error())
}
