package user

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
	ErrUserExists   = errors.New("this user is already registered")
	ErrUserNotFound = errors.New("user not found")
	ErrCreatingUser = errors.New("error creating user")
	ErrGettingUser  = errors.New("error getting user")
	ErrDeletingUser = errors.New("error deleting user")
	ErrUpdatingUser = errors.New("error updating user")
)

const (
	CreateUserQuery     = `INSERT INTO "user" (email, password, name, avatar) VALUES ($1, $2, $3, $4)`
	GetUserByIdQuery    = `SELECT user_id, email, password, name, avatar FROM "user" WHERE user_id=$1`
	GetUserByEmailQuery = `SELECT user_id, email, password, name, avatar FROM "user" WHERE email=$1`
	GetAllUsersQuery    = `SELECT user_id, email, password, name, avatar FROM "user"`
	DeleteUserQuery     = `DELETE FROM "user" WHERE user_id=$1`
	UpdateUserQuery     = `
    UPDATE "user"
    SET
        name = $2,
        avatar = $3,
        updated_at = NOW()
    WHERE user_id = $1
    RETURNING user_id, email, password, name, avatar;`
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string, name string) (models.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateUser(ctx context.Context, oldUser models.User) (models.User, error)
}

type PostgresUser struct {
	db *sql.DB
}

func NewPostgresUser(db *sql.DB) *PostgresUser {
	return &PostgresUser{db: db}
}

func (p *PostgresUser) CreateUser(ctx context.Context, email string, password string, name string) (models.User, error) {
	var user models.User

	defaultAvatar := minio.DefaultAvatarURL

	err := p.db.QueryRowContext(ctx,
		CreateUserQuery+" RETURNING user_id, email, password, name, avatar",
		email, password, name, defaultAvatar,
	).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Avatar)

	if err != nil {
		logger.Error(ctx, "Error creating user: %v", err)
		return models.User{}, ErrCreatingUser
	}

	return user, nil
}

func (p *PostgresUser) GetUserById(ctx context.Context, userID uuid.UUID) (models.User, error) {
	var user models.User

	err := p.db.QueryRowContext(ctx, GetUserByIdQuery, userID).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Avatar,
	)

	if err != nil {
		logger.Error(ctx, "Error getting user: %v", err)
		return models.User{}, ErrGettingUser
	}

	return user, nil
}

func (p *PostgresUser) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	err := p.db.QueryRowContext(ctx, GetUserByEmailQuery, email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Avatar,
	)

	if err != nil {
		logger.Error(ctx, "Error getting user: %v", err)
		return models.User{}, ErrGettingUser
	}

	return user, nil
}

func (p *PostgresUser) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := p.db.QueryContext(ctx, GetAllUsersQuery)
	if err != nil {
		logger.Error(ctx, "Error getting all users: %v", err)
		return nil, ErrGettingUser
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Avatar)
		if err != nil {
			logger.Error(ctx, "Error getting all users: %v", err)
			return nil, ErrGettingUser
		}
		users = append(users, user)
	}

	return users, nil

}

func (p *PostgresUser) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	_, err := p.db.ExecContext(ctx, DeleteUserQuery, id)
	if err != nil {
		logger.Error(ctx, "Error deleting user: %v", err)
		return false, ErrDeletingUser
	}

	return true, nil
}

func (p *PostgresUser) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	var updated models.User

	err := p.db.QueryRowContext(ctx, UpdateUserQuery,
		user.Id,
		user.Name,
		user.Avatar,
	).Scan(
		&updated.Id,
		&updated.Email,
		&updated.Password,
		&updated.Name,
		&updated.Avatar,
	)

	if err != nil {
		logger.Error(ctx, "Error updating user: %v", err)
		return models.User{}, ErrUpdatingUser
	}

	return updated, nil
}
