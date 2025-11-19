package cookies

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const SessionID = "session_id"

func SetCookie(w http.ResponseWriter, sessionId uuid.UUID) {
	cookie := &http.Cookie{
		Name:     SessionID,
		Value:    sessionId.String(),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(60 * time.Minute),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie(SessionID)
	if err != nil {
		logger.Error(nil, "cookie not found: %v", err)
		return nil, errors.New("cookie not found")
	}
	return cookie, nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(SessionID)
	if err != nil {
		logger.Error(nil, "cookie not found: %v", err)
		return errors.New("cookie not found")
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	return nil
}
