package auth

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err.Error())
		logger.Error(ctx, err.Error(), nil)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		logger.Error(ctx, err.Error(), nil)
		return
	}

	output, err := h.Usecase.Me(ctx, sessionID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error(), nil)
		return
	}

	userOutputDto := &dto.UserOutputMe{
		Email:  output.Email,
		Name:   output.Name,
		Avatar: output.Avatar,
	}

	err = json.Write(w, http.StatusOK, userOutputDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error(), nil)
		return
	}

}
