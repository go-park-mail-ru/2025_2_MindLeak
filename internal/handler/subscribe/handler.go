package subscribe

type Handler struct {
	Usecase subscribers.Usecase
}

func (h *Handler) NewSubsHandler(usecase subscribers.Usecase) *Handler {
	return &Handler{Usecase: usecase}
}
