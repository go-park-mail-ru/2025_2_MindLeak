package auth

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"net/http"
)

type AuthUsecase interface {
	Registration(input *UserRegisterInput) (*UserRegisterOutput, error)
	Login(input *UserLoginInput) (*UserLoginOutput, error)
	Logout() error
	Me() (*UserResponse, error)
}

type Handler struct {
	Usecase AuthUsecase
}

func NewHandler(u AuthUsecase) *Handler {
	return &Handler{Usecase: u}
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	newUserDto := new(UserRegisterInput)
	err := json.Read(r, newUserDto)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	registerDto, err := h.Usecase.Registration(newUserDto)
	if err != nil {
		json.WriteError(w, http.StatusConflict, err.Error()) //Тут бы подумать о том, как прокидывать из юзкейсов ошибки
	}

	json.Write(w, http.StatusCreated, registerDto)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Me(w http.ResponseWriter, r *http.Request)     {}
