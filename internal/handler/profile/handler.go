package profile

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile"

type Handler struct {
	Usecase profile.Usecase
}

func NewShowProfileHandler(u profile.Usecase) *Handler {
	return &Handler{Usecase: u}
}

func (h *Handler) handleError(err error) (int, string) {
	///
	return 0, "nil"
}
