package appeal

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (h *Handler) CreateAppeal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.AppealInputDto{}
	if err := json.Read(r, &input); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid json")
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

	logger.Warn(ctx, "CreateAppeal", input)

	appeal, err := h.Usecase.CreateAppeal(ctx, sessionID, input)
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
	json.Write(w, http.StatusCreated, out)

}
