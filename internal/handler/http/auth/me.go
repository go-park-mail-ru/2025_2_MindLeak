package auth

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger.Info(ctx, " handler start")

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		logger.Error(ctx, " failed to get cookie: %v", err)
		json.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	logger.Info(ctx, " cookie received: %s", cookie.Value)

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		logger.Error(ctx, " invalid session UUID: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	logger.Info(ctx, "parsed sessionID: %s", sessionID.String())

	output, err := h.Usecase.Me(ctx, sessionID)
	if err != nil {
		code, msg := h.handleError(err)
		logger.Error(ctx, " error: %v", err)
		json.WriteError(w, code, msg)
		return
	}

	logger.Info(ctx, " success, preparing response")

	userOutputDto := &dto.UserOutputMe{
		Id:     output.Id,
		Email:  output.Email,
		Name:   output.Name,
		Avatar: output.Avatar,
	}

	if err = json.Write(w, http.StatusOK, userOutputDto); err != nil {
		logger.Error(ctx, "[auth.Me] failed to write response: %v", err)
		return
	}

	logger.Info(ctx, "handler finished successfully")
}
