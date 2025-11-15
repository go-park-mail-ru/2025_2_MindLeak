package appeal

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

func (h *Handler) CreateAppeal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := dto.AppealInputDto{}
	if err := json.Read(r, &input); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	appeal, err := h.Usecase.CreateAppeal(ctx, input)
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
