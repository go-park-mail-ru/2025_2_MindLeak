package auth

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userInputDto := &dto.UserInputLogin{}
	err := json.Read(r, userInputDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

	oldUserEntity := models.User{
		Email:    userInputDto.Email,
		Password: userInputDto.Password,
	}

	user, sessionID, err := h.Usecase.Login(ctx, oldUserEntity)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

	userOutputDto := &dto.UserOutputLogin{
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
	}

	cookies.SetCookie(w, sessionID)

	err = json.Write(w, http.StatusOK, userOutputDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

}
