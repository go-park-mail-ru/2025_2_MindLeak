package topBlogs

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/topBlogs"

type Handler struct {
	Usecase topBlogs.Usecase
}

func NewTopBlogsHandler(usecase topBlogs.Usecase) *Handler {
	return &Handler{Usecase: usecase}
}

//func (h *Handler) handleError(err error) (int, string) {}
