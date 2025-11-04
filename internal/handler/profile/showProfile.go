package profile

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/profile/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) ShowProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	vars := mux.Vars(r)
	var targetID uuid.UUID
	if idStr, ok := vars["id"]; ok && idStr != "" {
		targetID, err = uuid.Parse(idStr)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}
	} else {
		session, err := h.Usecase.GetSession(ctx, sessionID)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}
		targetID = session.UserId
	}

	prof, err := h.Usecase.ShowProfile(ctx, targetID)
	if err != nil {
		logger.Error(ctx, "ShowProfile: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	profileOutDto := dto.ProfileOutputDto{
		Phone:       prof.Phone,
		Country:     prof.Country,
		Language:    prof.Language,
		Sex:         prof.Sex,
		DateOfBirth: prof.DateOfBirth,
		Age:         prof.Age,
		Description: prof.Description,
		CoverURL:    prof.CoverURL,
		Name:        prof.Name,
		AvatarURL:   prof.Avatar,
		Email:       prof.Email,
	}

	err = json.Write(w, http.StatusOK, profileOutDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

}
