package appeal

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (h *Handler) UploadScreenshot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		logger.Error(ctx, err.Error())
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		logger.Error(ctx, err.Error())
		json.WriteError(w, http.StatusUnauthorized, "invalid session")
		return
	}

	appealID, err := uuid.Parse(r.URL.Query().Get("appealId"))
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid appeal id")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		logger.Error(ctx, err.Error())
		json.WriteError(w, http.StatusBadRequest, "file not provided")
		return
	}
	defer file.Close()

	appeal, err := h.Usecase.UploadScreenshot(ctx, sessionID, appealID, file, header)
	if err != nil {
		logger.Error(ctx, err.Error())
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, appeal)
}
