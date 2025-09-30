package registration

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

type UserRegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request, sessions *session.InMemorySession,
	users *repository.InMemoryUser) {

	newUserData := new(UserRegisterInput)
	err := json.Read(r, newUserData)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := validateEmail(newUserData.Email); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validatePassword(newUserData.Password); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateName(newUserData.Name); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	User, err := users.CreateUser(newUserData.Email, newUserData.Password, newUserData.Name)
	if err != nil {
		json.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	session, err := sessions.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cookies.SetCookie(w, session.SessionId)

	_, err = sessions.SetSessionUserId(session.SessionId, User.Id)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Write(w, http.StatusCreated, User)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

var emailRequired = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func validateEmail(email string) error {
	if email == "" {
		return errors.New("Email, password and name are required")
	}
	if !emailRequired.MatchString(email) {
		return errors.New("Email is invalid")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("Email, password and name are required")
	}

	if utf8.RuneCountInString(password) < 4 {
		return errors.New("Password is too short")
	}

	if strings.Contains(password, " ") {
		return errors.New("Password is invalid")
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("Email, password and name are required")
	}

	if strings.Contains(name, " ") {
		return errors.New("Name is invalid")
	}

	if utf8.RuneCountInString(name) < 4 {
		return errors.New("Name is too short")
	}

	return nil
}
