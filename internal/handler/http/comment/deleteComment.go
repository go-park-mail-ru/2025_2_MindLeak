package comment

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commentIdString := r.URL.Query().Get("id")

	if commentIdString == "" {
		json.WriteError(w, http.StatusBadRequest, "mandatory query parameter not supplied")
		logger.Error(ctx, "mandatory query parameter not supplied")
		return
	}

	commentId, err := uuid.Parse(commentIdString)

	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

	flag, err := h.Usecase.DeleteComment(ctx, commentId)

	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

	if err = json.Write(w, http.StatusOK, flag); err != nil {
		logger.Error(ctx, err.Error())
		return
	}
}
