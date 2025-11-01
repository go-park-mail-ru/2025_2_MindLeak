package profile

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/profile/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"net/http"
)

func EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	var NewInput dto.ProfileInputDto
	json.Read(r, &NewInput)
}
