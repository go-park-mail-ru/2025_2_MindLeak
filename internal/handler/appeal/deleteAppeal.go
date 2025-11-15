package appeal

import (
	"net/http"
)

func (h *Handler) DeleteAppeal(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// idStr := r.URL.Query().Get("id")
	// if idStr == "" {
	// 	json.WriteError(w, http.StatusBadRequest, "missing appeal ID")
	// 	return
	// }
	// appealID, err := uuid.Parse(idStr)
	// if err != nil {
	// 	json.WriteError(w, http.StatusBadRequest, "invalid appeal ID")
	// 	return
	// }
	// deleted, err := h.Usecase.DeleteAppeal(ctx, appealID)
	// if err != nil {
	// 	code, msg := h.handleError(err)
	// 	json.WriteError(w, code, msg)
	// 	return
	// }
	// if !deleted {
	// 	json.WriteError(w, http.StatusNotFound, "appeal not found")
	// 	return
	// }
	// w.WriteHeader(http.StatusNoContent)
}
