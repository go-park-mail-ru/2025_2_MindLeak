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

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		logger.Error(ctx, "GetCookie: %v", err)
		///
		return
	}
	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		logger.Error(ctx, "Parse: %v", err)
		///
		return
	}

	if myProfile, err := h.Usecase.ShowProfile(ctx, sessionID); err != nil {
		logger.Error(ctx, "ShowProfile: %v", err)
		///
		return
	}

	profileOutDto := dto.ProfileOutputDto{}

	err = json.Write(w, http.StatusOK, profileOutDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error(), nil)
		return
	}

}
