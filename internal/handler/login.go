package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

type UserLoginInput struct {
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request, sessions *repository.InMemorySession,
	users *repository.InMemoryUser) {
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

	Email := newUserData.Email
	Password := newUserData.Password
	fmt.Println(newUserData.Email, newUserData.Password)

	User, err := users.GetUserByEmail(Email)
	if err != nil {
		json.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	log.Println("FOUND:", User.Email, User.Password)

	if User.Password != Password {
		json.WriteError(w, http.StatusUnauthorized, "invalid password")
		return
	}

	session, err := sessions.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cookies.SetCookie(w, session.SessionId)

	_, err = sessions.SetSessionUserId(session.SessionId, User.Id) //Pair UserId and SessionId
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Write(w, http.StatusOK, User) //Writes json with User and Status as an answer
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

}
