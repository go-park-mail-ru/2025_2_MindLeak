package profile

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/profile/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) ShowProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()
	targetIDStr := queryParams.Get("id")

	var targetID uuid.UUID

	if targetIDStr == "" {

		cookie, err := cookies.GetCookie(r)
		if err != nil {
			logger.Error(ctx, "GetCookie: %v", err)
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}

		sessionID, err := uuid.Parse(cookie.Value)
		if err != nil {
			logger.Error(ctx, "Parse: %v", err)
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}

		session, err := h.Usecase.GetSession(ctx, sessionID)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}

		targetID = session.UserId
		logger.Info(ctx, "Fetching own profile for user ID: %v", targetID)
	} else {

		var err error
		targetID, err = uuid.Parse(targetIDStr)
		if err != nil {
			logger.Error(ctx, "Invalid user ID format: %v", err)
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}

		logger.Info(ctx, "Fetching other user's profile with ID: %v", targetID)
	}

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
