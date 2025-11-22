package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/profile"
	repoProfile "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/profile"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	repoSession "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
	repoUser "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/minio_client"
	"github.com/microcosm-cc/bluemonday"
)

var (
	UserExists        = errors.New("user is already registered")
	UserNotFound      = errors.New("user not found")
	ServerError       = errors.New("internal apiserver error")
	SessionNotFound   = errors.New("session not found")
	SessionNotCreated = errors.New("session not created")
	SessionNotSet     = errors.New("session not set")
	UserNotCreated    = errors.New("user not created")
	UserNotUpdated    = errors.New("user not updated")
	UserNotDeleted    = errors.New("user not deleted")
	UserNotGet        = errors.New("user not get")
	ProfileNotCreated = errors.New("profile not created")
	ProfileNotUpdated = errors.New("profile not updated")
	ProfileNotDeleted = errors.New("profile not deleted")
	ProfileNotGet     = errors.New("profile not get")
)

type Usecase struct {
	sessionRepo session.SessionRepository
	userRepo    user.UserRepository
	profileRepo profile.ProfileRepository
	minioClient *minio_client.Client
	sanitizer   *bluemonday.Policy
}

func NewProfileUsecase(sessionRepo session.SessionRepository, userRepo user.UserRepository, profileRepo profile.ProfileRepository, minioClient *minio_client.Client) *Usecase {
	sanitizer := bluemonday.UGCPolicy()
	return &Usecase{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		profileRepo: profileRepo,
		minioClient: minioClient,
		sanitizer:   sanitizer,
	}
}

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, repoUser.ErrUserExists):
		return UserExists
	case errors.Is(err, repoUser.ErrUserNotFound):
		return UserNotFound
	case errors.Is(err, repoSession.ErrSessionNotFound):
		return SessionNotFound
	case errors.Is(err, repoSession.ErrCreatingSession):
		return SessionNotCreated
	case errors.Is(err, repoSession.ErrSettingSession):
		return SessionNotSet
	case errors.Is(err, repoUser.ErrCreatingUser):
		return UserNotCreated
	case errors.Is(err, repoUser.ErrDeletingUser):
		return UserNotDeleted
	case errors.Is(err, repoUser.ErrGettingUser):
		return UserNotGet
	case errors.Is(err, repoUser.ErrUpdatingUser):
		return UserNotUpdated
	case errors.Is(err, repoProfile.ErrCreatingProfile):
		return ProfileNotCreated
	case errors.Is(err, repoProfile.ErrGettingProfile):
		return ProfileNotUpdated
	case errors.Is(err, repoProfile.ErrUpdatingProfile):
		return ProfileNotGet
	case errors.Is(err, repoProfile.ErrDeletingProfile):
		return ProfileNotDeleted
	default:
		return ServerError
	}
}
