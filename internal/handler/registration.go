package handler

import (
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

type UserRegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request, sessions *repository.InMemorySession,
	users *repository.InMemoryUser) {

	newUserData := new(UserRegisterInput)
	err := json.Read(r, newUserData)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if newUserData.Email == "" || newUserData.Password == "" || newUserData.Name == "" {
		json.WriteError(w, http.StatusBadRequest, "Email or Password or Name is required")
		return
	}

	User, err := users.CreateUser(newUserData.Email, newUserData.Password, newUserData.Name) //Add new user in storage
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
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

	err = json.Write(w, http.StatusCreated, User) //Writes json with User and Status as an answer
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("REGISTER:", newUserData.Email, newUserData.Password, newUserData.Name)
}
