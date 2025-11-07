package profile

import (
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/profile/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) ShowOtherProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()
	targetIDStr := queryParams.Get("id")

	if targetIDStr == "" {
		logger.Error(ctx, "No user ID provided in query")
		code, msg := h.handleError(fmt.Errorf("no user ID"))
		json.WriteError(w, code, msg)
		return
	}

	targetID, err := uuid.Parse(targetIDStr)
	if err != nil {
		logger.Error(ctx, "Invalid user ID format: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	logger.Info(ctx, "Fetching profile for user ID: %v", targetID)
	prof, err := h.Usecase.ShowProfile(ctx, targetID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	profileOutDto := dto.ProfileOutputDto{
		Phone:         prof.Phone,
		Country:       prof.Country,
		Language:      prof.Language,
		Sex:           prof.Sex,
		DateOfBirth:   prof.DateOfBirth,
		Age:           prof.Age,
		Description:   prof.Description,
		CoverURL:      prof.CoverURL,
		Name:          prof.Name,
		AvatarURL:     prof.Avatar,
		Email:         prof.Email,
		CreatedAt:     prof.CreatedAt.Format("2006-01-02"),
		Subscribers:   prof.Subscribers,
		Subscriptions: prof.Subscriptions,
		Password:      prof.Password,
	}

	err = json.Write(w, http.StatusOK, profileOutDto)
	if err != nil {
		logger.Error(ctx, err.Error())
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}
}
