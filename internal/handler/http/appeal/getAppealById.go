package appeal

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func (h *Handler) GetAppealByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		json.WriteError(w, http.StatusBadRequest, "missing appeal ID")
		return
	}

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}
	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	appealID, err := uuid.Parse(idStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid appeal ID")
		return
	}

	appeal, err := h.Usecase.GetAppealByID(ctx, sessionID, appealID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	out := dto.AppealOutputDto{
		Id:                 appeal.AppealID,
		EmailRegistered:    appeal.EmailRegistered,
		Status:             dto.Status(appeal.Status),
		ProblemDescription: appeal.ProblemDescription,
		ScreenshotURL:      appeal.ScreenshotURL,
		CategoryID:         appeal.CategoryID,
		Name:               appeal.Name,
		EmailForConnect:    appeal.EmailForConnect,
	}

	json.Write(w, http.StatusOK, out)
}
