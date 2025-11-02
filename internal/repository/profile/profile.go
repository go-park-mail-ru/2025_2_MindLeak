package profile

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/minio"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

var (
	ErrCreatingProfile = errors.New("error creating profile")
	ErrUpdatingProfile = errors.New("error updating profile")
	ErrDeletingProfile = errors.New("error deleting profile")
	ErrGettingProfile  = errors.New("error getting profile")
)

const (
	CreateProfileQuery = `INSERT INTO profile (UserID, CoverURL) VALUES ($1, $2)`
	GetProfileQuery    = `
        SELECT
            user_id,
            phone,
            country,
            language,
            sex,
            date_of_birth,
            age,
            cover_url,
            created_at,
            updated_at
        FROM profile
        WHERE user_id = $1
    `
	DeleteProfileQuery = `DELETE FROM profile WHERE UserID = $1`
	UpdateProfileQuery = `
    UPDATE profile
    SET
        phone = $2,
        country = $3,
        language = $4,
        sex = $5,
        date_of_birth = $6,
        cover_url = $7,
        age = $8,
        updated_at = NOW()
    WHERE user_id = $1
    RETURNING user_id, phone, country, language, sex, date_of_birth, cover_url, age, created_at, updated_at;`
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
		return models.Profile{}, ErrCreatingProfile
	}

	return profile, nil
}

func (p *PostgresProfile) GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	var prof models.Profile
	err := p.db.QueryRowContext(ctx, GetProfileQuery, userID).Scan(
		&prof.UserID, &prof.Phone, &prof.Country, &prof.Language, &prof.Sex,
		&prof.DateOfBirth, &prof.Age, &prof.CoverURL, &prof.CreatedAt, &prof.UpdatedAt,
	)
	if err != nil {
		logger.Error(ctx, "Error getting profile: %v", err)
		return models.Profile{}, ErrGettingProfile
	}
	return prof, nil
}

func (p *PostgresProfile) UpdateProfile(ctx context.Context, profile models.Profile) (models.Profile, error) {
	var updated models.Profile

	err := p.db.QueryRowContext(ctx, UpdateProfileQuery,
		profile.UserID,
		profile.Phone,
		profile.Country,
		profile.Language,
		profile.Sex,
		profile.DateOfBirth,
		profile.CoverURL,
		profile.Age,
	).Scan(
		&updated.UserID,
		&updated.Phone,
		&updated.Country,
		&updated.Language,
		&updated.Sex,
		&updated.DateOfBirth,
		&updated.CoverURL,
		&updated.Age,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		logger.Error(ctx, "Error updating profile: %v", err)
		return models.Profile{}, ErrUpdatingProfile
	}

	return updated, nil
}

//func (p *PostgresProfile) DeleteProfile(ctx context.Context, uuid uuid.UUID) error {
//
//}
