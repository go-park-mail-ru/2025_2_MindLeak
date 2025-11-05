package comment

import (
	dto2 "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/comment/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"
)

func (h *Handler) UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	inputDto := dto2.CommentIODto{}

	err := json.Read(r, &inputDto)

	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, "update comment: %v", err)
		return
	}

	updated, err := h.Usecase.UpdateComment(ctx, h.mapToDto(inputDto))

	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, "update comment: %v", err)
		return
	}

	if err = json.Write(w, http.StatusOK, updated); err != nil {
		logger.Error(ctx, "update comment: %v", err)
		return
	}
}
