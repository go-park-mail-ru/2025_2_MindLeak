package subscriptions

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		json.Write(w, http.StatusBadRequest, "missing user id")
		return
	}

	userID, err := uuid.Parse(idStr)
	if err != nil {
		json.Write(w, http.StatusBadRequest, "invalid uuid")
		return
	}

	subscriptions, err := h.Usecase.GetSubscriptions(ctx, userID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

	output := h.mapSubscriptionsToOutputDto(subscriptions)

	json.Write(w, http.StatusOK, output)
}
