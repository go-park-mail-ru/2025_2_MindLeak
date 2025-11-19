package appeal

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := h.Usecase.GetCategories(ctx)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	outputDto := dto.CategoryOutputDto{
		Categories: make([]dto.CategoryDto, 0, len(categories)),
	}

	for _, c := range categories {
		outputDto.Categories = append(outputDto.Categories, dto.CategoryDto{
			CategoryID: c.CategoryID,
			Name:       c.Name,
		})
	}

	json.Write(w, http.StatusOK, outputDto)
}
