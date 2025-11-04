package profile

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/minio_client"
	"github.com/google/uuid"
	"time"
)

var (
	ErrCreatingProfile = errors.New("error creating profile")
	ErrUpdatingProfile = errors.New("error updating profile")
	ErrDeletingProfile = errors.New("error deleting profile")
	ErrGettingProfile  = errors.New("error getting profile")
)

const (
	CreateProfileQuery = `
    INSERT INTO profile (user_id, phone, country, language, sex, date_of_birth, age, description, cover_url)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING profile_id, user_id, phone, country, language, sex, date_of_birth, age, description, cover_url, created_at, updated_at
	`
	GetProfileQuery = `
	SELECT 
	    profile_id,
	    user_id,
	    phone,
	    country,
	    language,
	    sex,
	    date_of_birth,
	    age,
	    description,
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
		phone = COALESCE($2, phone),
		country = COALESCE($3, country),
		language = COALESCE($4, language),
		sex = COALESCE($5, sex),
		date_of_birth = COALESCE($6, date_of_birth),
		age = COALESCE($7, age),
		description = COALESCE($8, description),
		cover_url = COALESCE($9, cover_url),
		updated_at = CURRENT_TIMESTAMP
	WHERE user_id = $1
	RETURNING 
	    profile_id,
	    user_id,
	    phone,
	    country,
	    language,
	    sex,
	    date_of_birth,
	    age,
	    description,
	    cover_url,
	    created_at,
	    updated_at
	`
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, UserID uuid.UUID) (models.Profile, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error)
	UpdateProfile(ctx context.Context, profile models.Profile) (models.Profile, error)
	//DeleteProfile(ctx context.Context, uuid uuid.UUID) error
}

type PostgresProfile struct {
	db    *sql.DB
	minio *minio_client.Client
}

func NewPostgresProfile(db *sql.DB, minio *minio_client.Client) *PostgresProfile {
	return &PostgresProfile{db: db, minio: minio}
}

func (p *PostgresProfile) CreateProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	var profile models.Profile

	defaultCover := p.minio.GetDefaultCover()
	defaultSex := models.SexUndefined
	defaultDate := time.Time{}
	defaultPhone := ""
	defaultCountry := ""
	defaultLanguage := ""
	dafaultDescription := ""
	defaultAge := 0

	err := p.db.QueryRowContext(ctx, CreateProfileQuery, userID, defaultPhone, defaultCountry, defaultLanguage, defaultSex, defaultDate, defaultAge, dafaultDescription, defaultCover).Scan(
		&profile.Id,
		&profile.UserID,
		&profile.Phone,
		&profile.Country,
		&profile.Language,
		&profile.Sex,
		&profile.DateOfBirth,
		&profile.Age,
		&profile.Description,
		&profile.CoverURL,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if err != nil {
		return models.Profile{}, fmt.Errorf("create profile: %w", err)
	}

	return profile, nil
}

func (p *PostgresProfile) GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	var profile models.Profile
	err := p.db.QueryRowContext(ctx, GetProfileQuery, userID).Scan(
		&profile.Id,
		&profile.UserID,
		&profile.Phone,
		&profile.Country,
		&profile.Language,
		&profile.Sex,
		&profile.DateOfBirth,
		&profile.Age,
		&profile.Description,
		&profile.CoverURL,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Profile{}, ErrGettingProfile
	}
	if err != nil {
		return models.Profile{}, fmt.Errorf("get profile: %w", err)
	}
	return profile, nil
}

func (p *PostgresProfile) UpdateProfile(ctx context.Context, prof models.Profile) (models.Profile, error) {
	var updated models.Profile
	err := p.db.QueryRowContext(ctx, UpdateProfileQuery,
		prof.UserID,
		prof.Phone,
		prof.Country,
		prof.Language,
		prof.Sex,
		prof.DateOfBirth,
		prof.Age,
		prof.Description,
		prof.CoverURL,
	).Scan(
		&updated.Id,
		&updated.UserID,
		&updated.Phone,
		&updated.Country,
		&updated.Language,
		&updated.Sex,
		&updated.DateOfBirth,
		&updated.Age,
		&updated.Description,
		&updated.CoverURL,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Profile{}, ErrUpdatingProfile
	}
	if err != nil {
		return models.Profile{}, fmt.Errorf("update profile: %w", err)
	}
	return updated, nil
}

//func (p *PostgresProfile) DeleteProfile(ctx context.Context, uuid uuid.UUID) error {
//
//}
