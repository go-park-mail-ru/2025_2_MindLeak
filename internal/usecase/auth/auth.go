package auth

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/entities"
)

type AuthUsecase interface {
	Registration(User entities.User) (dto.RegisteredUserDto, error)
}
