package categories

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"
	"strconv"
)

func (h *Handler) CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	topic := r.URL.Query().Get("topic")
	offset := r.URL.Query().Get("offset")

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		logger.Error(ctx, err.Error())
		json.WriteError(w, http.StatusBadRequest, "Invalid offset")
		return
	}

	posts, err := h.Usecase.GetFeedByTopic(ctx, topic, offsetInt)
	if err != nil {
		logger.Error(ctx, err.Error())
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Write(w, http.StatusOK, posts)
	if err != nil {
		logger.Error(ctx, err.Error())
		return
	}
}
