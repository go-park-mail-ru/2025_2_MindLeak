package sessions

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
)

func SetCookie(w http.ResponseWriter, sessionId uuid.UUID) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId.String(),
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, errors.New("cookie not found")
	}
	return cookie, nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return errors.New("cookie not found")
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	return nil
}
