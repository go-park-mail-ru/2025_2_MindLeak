package usecase

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/entities"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

func (p *Processor) Registration(w http.ResponseWriter, user entities.User) (*dto.RegisteredUserDto, error) {

	if err := validateEmail(user.Email); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validatePassword(user.Password); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateName(user.Name); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	//Тут видимо нужно дернуть методы репозитория

	//Возвращаем созданного в базе юзера
	//Возвращаться будет доменная модель уровня репы, которая также пойдет в базу. У него уже будет айдишник и тд
	newUser, err := p.userRepo.CreateUser(user.Email, user.Password, user.Name)
	if err != nil {
		json.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	//Сейм история только про сессии
	session, err := p.sessionRepo.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	cookies.SetCookie(w, session.SessionId)

	_, err = p.sessionRepo.SetSessionUserId(session.SessionId, newUser.Id)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	outDto := &dto.RegisteredUserDto{
		Email:  newUser.Email,
		Name:   newUser.Name,
		Avatar: newUser.Avatar,
	}

	return outDto, nil

	//err = json.Write(w, http.StatusCreated, user)
	//if err != nil {
	//	json.WriteError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

}

var emailRequired = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func validateEmail(email string) error {
	if email == "" {
		return errors.New("email, password and name are required")
	}

	if !emailRequired.MatchString(email) {
		return errors.New("email is invalid")
	}

	if utf8.RuneCountInString(email) > 320 {
		return errors.New("email is too long")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("email, password and name are required")
	}

	if utf8.RuneCountInString(password) < 4 {
		return errors.New("password is too short")
	}

	if strings.Contains(password, " ") {
		return errors.New("password is invalid")
	}

	if utf8.RuneCountInString(password) > 64 {
		return errors.New("password is too long")
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("email, password and name are required")
	}

	if strings.Contains(name, " ") {
		return errors.New("name is invalid")
	}

	if utf8.RuneCountInString(name) < 4 {
		return errors.New("name is too short")
	}

	if utf8.RuneCountInString(name) > 32 {
		return errors.New("name is too long")
	}

	return nil
}
