package subscriptions

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/subscriptions/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/subscriptions"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/subscriptions/usecase"
	"net/http"
)

type Handler struct {
	Usecase subscriptions.Usecase
}

func NewSubsHandler(usecase subscriptions.Usecase) *Handler {
	return &Handler{Usecase: usecase}
}

func (h *Handler) mapToSubscriptionDto(u models.User) dto.SubscriptionDto {
	return dto.SubscriptionDto{
		Id:          u.Id,
		Avatar:      u.Avatar,
		Name:        u.Name,
		Subscribers: u.Subscribers,
	}
}

func (h *Handler) mapSubscriptionsToOutputDto(users []models.User) dto.SubscriptionsOutputDto {
	res := make([]dto.SubscriptionDto, 0, len(users))

	for _, u := range users {
		res = append(res, h.mapToSubscriptionDto(u))
	}

	return dto.SubscriptionsOutputDto{
		Subsccriptions: res,
	}
}

func (h *Handler) mapToSubscriberDto(u models.User) dto.SubscriberDto {
	return dto.SubscriberDto{
		Id:          u.Id,
		Avatar:      u.Avatar,
		Name:        u.Name,
		Subscribers: u.Subscribers,
	}
}

func (h *Handler) mapSubscribersToOutputDto(users []models.User) dto.SubscribersOutputDto {
	res := make([]dto.SubscriberDto, 0, len(users))

	for _, u := range users {
		res = append(res, h.mapToSubscriberDto(u))
	}

	return dto.SubscribersOutputDto{
		Subscribers: res,
	}
}

func (h *Handler) handleError(err error) (int, string) {
	switch {

	case errors.Is(err, usecase.ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized"

	case errors.Is(err, usecase.ErrUserNotFound):
		return http.StatusNotFound, "user not found"

	case errors.Is(err, usecase.ErrSelfSubscribe):
		return http.StatusBadRequest, "cannot subscribe to yourself"

	case errors.Is(err, usecase.ErrAlreadySubscribed):
		return http.StatusConflict, "already subscribed"

	case errors.Is(err, usecase.ErrNotSubscribed):
		return http.StatusConflict, "not subscribed"

	case errors.Is(err, usecase.ErrInternal):
		return http.StatusInternalServerError, "internal server error"

	default:
		return http.StatusInternalServerError, "unexpected error"
	}
}
