package handler

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
	"net/http"
)

type UserRegisterInput struct {
	Email    string
	Password string
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request, sessions *repository.InMemorySession,
	users *repository.InMemoryUser) {

	newUserData := new(UserRegisterInput)
	err := json.Read(r, newUserData)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	User, err := users.CreateUser(newUserData.Email, newUserData.Password) //Add new user in storage
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cookie, err := cookies.GetCookie(r) //Search guest cookie
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	sessionId, err := uuid.Parse(cookie.Value)             //Search sessionId
	_, err = sessions.SetSessionUserId(sessionId, User.Id) //Pair UserId and SessionId
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = json.Write(w, http.StatusCreated, User) //Writes json with User and Status as an answer
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
}
