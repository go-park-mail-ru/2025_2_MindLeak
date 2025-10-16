package models

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth/dto"

type User struct {
	Email    string
	Password string
	Name     string
}

func Converter(userInputDto dto.UserInputRegistration) User {
	return User{
		Email:    userInputDto.Email,
		Name:     userInputDto.Name,
		Password: userInputDto.Password,
	}
}
