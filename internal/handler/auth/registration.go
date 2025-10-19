package auth

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"net/http"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Эта проверка уйдет с переходом на гориллу
	if r.Method != http.MethodPost {
		json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userInputDto := &dto.UserInputRegistration{}
	err := json.Read(r, userInputDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	newUserEntity := &models.User{
		Email:    userInputDto.Email,
		Password: userInputDto.Password,
		Name:     userInputDto.Name,
	}

	output, sessionID, err := h.Usecase.Registration(ctx, *newUserEntity)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	userOutputDto := &dto.UserOutputRegistration{
		Email:  output.Email,
		Name:   output.Name,
		Avatar: output.Avatar,
	}

	cookies.SetCookie(w, sessionID)

	err = json.Write(w, http.StatusCreated, userOutputDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}
}
