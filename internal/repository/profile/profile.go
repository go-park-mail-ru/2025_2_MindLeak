package profile

import (
	"context"
	"database/sql"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/minio"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

const (
	CreateProfileQuery = `INSERT INTO profile (UserID, CoverURL) VALUES ($1, $2)`
	GetProfileQuery    = `SELECT id, UserID, CoverURL FROM profile WHERE UserID = $1`
	DeleteProfileQuery = `DELETE FROM profile WHERE UserID = $1`
	UpdateProfileQuery = ``
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, UserID uuid.UUID) (models.Profile, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error)
	UpdateProfile(ctx context.Context, profile models.Profile) (models.Profile, error)
	//DeleteProfile(ctx context.Context, uuid uuid.UUID) error
}

type PostgresProfile struct {
	db *sql.DB
}

func NewPostgresProfile(db *sql.DB) *PostgresProfile {
	return &PostgresProfile{db: db}
}

func (p *PostgresProfile) CreateProfile(ctx context.Context, UserID uuid.UUID) (models.Profile, error) {
	var profile models.Profile

	defaultCover := minio.DefaultCoverURL

	err := p.db.QueryRowContext(ctx, CreateProfileQuery+" RETURNING UserID, CoverURL",
		UserID, defaultCover).Scan(&profile.UserID, &profile.CoverURL)
	if err != nil {
		logger.Error(ctx, "Error creating profile: %v", err)
		return models.Profile{}, err
	}

	return profile, nil
}

func (p *PostgresProfile) GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	var profile models.Profile

	err := p.db.QueryRowContext(ctx, GetProfileQuery, userID).Scan(&profile.UserID, &profile.CoverURL)
	if err != nil {
		logger.Error(ctx, "Error getting profile: %v", err)
		return models.Profile{}, err
	}

	return profile, nil
}

func (p *PostgresProfile) UpdateProfile(ctx context.Context, profile models.Profile) (models.Profile, error) {
	return profile, nil
}

//func (p *PostgresProfile) DeleteProfile(ctx context.Context, uuid uuid.UUID) error {
//
//}
