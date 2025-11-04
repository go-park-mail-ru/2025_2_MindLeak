package categories

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/categories"

type Handler struct {
	Usecase categories.Usecase
}

func NewCategoriesHandler(u categories.Usecase) *Handler {
	return &Handler{Usecase: u}
}

//func (h *Handler) handleError(err error) (int, string) {
//
//}
