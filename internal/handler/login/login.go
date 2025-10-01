package login

import (
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

type UserLoginInput struct {
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request, sessions *session.InMemorySession,
	users *user.InMemoryUser) {
	if r.Method != http.MethodPost {
		json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	newUserData := new(UserLoginInput)
	err := json.Read(r, newUserData)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if newUserData.Email == "" || newUserData.Password == "" {
		json.WriteError(w, http.StatusBadRequest, "Email or Password is required")
		return
	}

	email := newUserData.Email
	password := newUserData.Password

	user, err := users.GetUserByEmail(email)
	if err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	log.Println("FOUND:", user.Email, user.Password)

	if user.Password != password {
		json.WriteError(w, http.StatusUnauthorized, "invalid password")
		return
	}

	session, err := sessions.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cookies.SetCookie(w, session.SessionId)

	_, err = sessions.SetSessionUserId(session.SessionId, user.Id)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Write(w, http.StatusOK, user)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

}
