package profile

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/usecase"
	"net/http"
)

type Handler struct {
	Usecase profile.Usecase
}

func NewProfileHandler(u profile.Usecase) *Handler {
	return &Handler{Usecase: u}
}

func (h *Handler) handleError(err error) (int, string) {
	switch {
	case errors.Is(err, usecase.ServerError):
		return http.StatusInternalServerError, "internal server error"
	case errors.Is(err, usecase.UserExists):
		return http.StatusUnauthorized, "user not exists"
	case errors.Is(err, usecase.UserNotFound):
		return http.StatusUnauthorized, "user not found"
	case errors.Is(err, usecase.SessionNotFound):
		return http.StatusUnauthorized, "session not found"
	case errors.Is(err, usecase.SessionNotCreated):
		return http.StatusUnauthorized, "session not created"
	case errors.Is(err, usecase.SessionNotSet):
		return http.StatusUnauthorized, "session not set"
	case errors.Is(err, usecase.UserNotCreated):
		return http.StatusUnauthorized, "user not created"
	case errors.Is(err, usecase.UserNotUpdated):
		return http.StatusUnauthorized, "user not updated"
	case errors.Is(err, usecase.UserNotDeleted):
		return http.StatusUnauthorized, "user not deleted"
	case errors.Is(err, usecase.UserNotGet):
		return http.StatusUnauthorized, "user not get"
	case errors.Is(err, usecase.ProfileNotCreated):
		return http.StatusUnauthorized, "profile not created"
	case errors.Is(err, usecase.ProfileNotUpdated):
		return http.StatusUnauthorized, "profile not updated"
	case errors.Is(err, usecase.ProfileNotDeleted):
		return http.StatusUnauthorized, "profile not deleted"
	case errors.Is(err, usecase.ProfileNotGet):
	}
	return 0, "nil"
}
