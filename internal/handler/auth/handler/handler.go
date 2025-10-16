package handler

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth"

type Handler struct {
	Usecase auth.AuthUsecase
}

func NewHandler(u auth.AuthUsecase) *Handler {
	return &Handler{Usecase: u}
}
