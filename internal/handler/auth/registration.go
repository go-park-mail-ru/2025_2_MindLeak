package auth

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		json.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userInputDto := &dto.UserInputRegistration{}
	err := json.Read(r, userInputDto)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//Отправляем IN dto в конвертер энтити уровня юзкейса. Возвращает структуру уровня юзкейса
	//Которую мы потом будем рассылать по методам юзкейса
	newUserEntity := models.Converter(*userInputDto)

	//На уровне юзкейса я планирую сделать одну дто для отдачи бизнес-сущности на уровень контроллера
	//То есть регистрация по идее должна будет возвращать юзкейс дто, который мы потом замапим в OUT
	//дто хендлера
	output, err := h.Usecase.Registration(newUserEntity)
	userOutputDto := &dto.UserOutputRegistration{
		Email:  output.Email,
		Name:   output.Name,
		Avatar: output.Avatar,
	}
	if err != nil {
		json.WriteError(w, http.StatusConflict, err.Error()) //Тут бы подумать о том, как прокидывать из юзкейсов ошибки
	}

	json.Write(w, http.StatusCreated, userOutputDto)
}
