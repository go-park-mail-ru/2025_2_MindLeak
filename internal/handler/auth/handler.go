package auth

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/usecase"
)

type Handler struct {
	Usecase auth.Usecase
}

func NewAuthHandler(u auth.Usecase) *Handler {
	return &Handler{Usecase: u}
}

func (h *Handler) handleError(err error) (int, string) {
	switch {
	case errors.Is(err, usecase.InvalidEmail),
		errors.Is(err, usecase.InvalidPassword),
		errors.Is(err, usecase.InvalidName):
		return http.StatusBadRequest, err.Error()

	case errors.Is(err, usecase.InvalidCredentials):
		return http.StatusUnauthorized, "invalid credentials"

	case errors.Is(err, usecase.SessionNotFound):
		return http.StatusUnauthorized, "session not found"

	case errors.Is(err, usecase.UserExists):
		return http.StatusConflict, "user already exists"

	case errors.Is(err, usecase.UserNotFound):
		return http.StatusNotFound, "user not found"

	case errors.Is(err, usecase.ServerError):
		return http.StatusInternalServerError, "internal server error"

	default:
		return http.StatusInternalServerError, "unexpected error"
	}
}
