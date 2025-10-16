package auth

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/models"
)

type AuthUsecase interface {
	Registration(User models.User) (dto.RegisteredUserDto, error)
}
