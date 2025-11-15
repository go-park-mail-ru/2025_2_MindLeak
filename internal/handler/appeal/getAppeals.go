package appeal

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func (h *Handler) GetAppeals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authorIDStr := r.URL.Query().Get("author_id")
	if authorIDStr == "" {
		json.WriteError(w, http.StatusBadRequest, "missing author ID")
		return
	}
	authorID, err := uuid.Parse(authorIDStr)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid author ID")
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

	appeals, err := h.Usecase.GetAppeals(ctx, sessionID, authorID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	var out []dto.AppealOutputDto
	for _, appeal := range appeals {
		out = append(out, dto.AppealOutputDto{
			Id:                 appeal.AppealID,
			EmailRegistered:    appeal.EmailRegistered,
			Status:             dto.Status(appeal.Status),
			ProblemDescription: appeal.ProblemDescription,
			ScreenshotURL:      appeal.ScreenshotURL,
			CategoryID:         appeal.CategoryID,
			Name:               appeal.Name,
			EmailForConnect:    appeal.EmailForConnect,
		})
	}
	json.Write(w, http.StatusOK, out)
}
