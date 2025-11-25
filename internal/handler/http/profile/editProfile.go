package profile

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/profile/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (h *Handler) EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		logger.Error(ctx, "GetCookie: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}
	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		logger.Error(ctx, "Parse: %v", err)
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	inputDto := dto.ProfileInputDto{}
	err = json.Read(r, &inputDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}

	userEntity := models.User{
		Name:     inputDto.Name,
		Avatar:   inputDto.Avatar,
		Password: inputDto.Password,
	}

	var dob time.Time
	if inputDto.DateOfBirth != "" {
		dob, err = time.Parse("2006-01-02", inputDto.DateOfBirth)
		if err != nil {
			json.WriteError(w, http.StatusBadRequest, "invalid date_of_birth (expected YYYY-MM-DD)")
			logger.Error(ctx, "parse dob: %v", err)
			return
		}
	}

	profileEntity := models.Profile{
		Phone:       inputDto.Phone,
		Country:     inputDto.Country,
		Language:    inputDto.Language,
		Sex:         inputDto.Sex,
		DateOfBirth: dob,
		Age:         inputDto.Age,
		Description: inputDto.Description,
		CoverURL:    inputDto.Cover,
	}

	newProfile, newUser, err := h.Usecase.EditProfile(ctx, sessionID, profileEntity, userEntity)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	outputDto := dto.ProfileOutputDto{
		Phone:       newProfile.Phone,
		Country:     newProfile.Country,
		Language:    newProfile.Language,
		Sex:         newProfile.Sex,
		DateOfBirth: newProfile.DateOfBirth.Format("2006-01-02"),
		Age:         newProfile.Age,
		Description: newProfile.Description,
		CoverURL:    newProfile.CoverURL,

		Name:      newUser.Name,
		AvatarURL: newUser.Avatar,
		Password:  newUser.Password,
	}

	err = json.Write(w, http.StatusOK, outputDto)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		logger.Error(ctx, err.Error())
		return
	}
}
