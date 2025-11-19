package topBlogs

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"
)

func (h *Handler) ShowTopBlogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogs, err := h.Usecase.ShowTopBlogs(ctx)
	if err != nil {
		code, msg := h.handleError(err)
		logger.Error(ctx, err.Error())
		json.WriteError(w, code, msg)
		return
	}
	outputDto := h.mapToOutputSlice(blogs)

	json.Write(w, http.StatusOK, outputDto)
}
