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
)

const (
	CreateUserQuery     = `INSERT INTO user (email, password, name, avatar) VALUES ($1, $2, $3, $4)`
	GetUserByIdQuery    = `SELECT id, email, password, name, avatar FROM user WHERE id=$1`
	GetUserByEmailQuery = `SELECT id, email, password, name, avatar FROM user WHERE email=$1`
	GetAllUsersQuery    = `SELECT id, email, password, name, avatar FROM user`
	DeleteUserQuery     = `DELETE FROM user WHERE id=$1`
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string, name string) (models.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (bool, error)
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
		CreateUserQuery+" RETURNING id, email, password, name, avatar",
		email, password, name, defaultAvatar,
	).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Avatar)

	if err != nil {
		logger.Error(ctx, "Error creating user: %v", err)
		return models.User{}, err
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
		return models.User{}, err
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
		return models.User{}, err
	}

	return user, nil
}

func (p *PostgresUser) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := p.db.QueryContext(ctx, GetAllUsersQuery)
	if err != nil {
		logger.Error(ctx, "Error getting all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Avatar)
		if err != nil {
			logger.Error(ctx, "Error getting all users: %v", err)
		}
		users = append(users, user)
	}

	return users, nil

}

func (p *PostgresUser) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	_, err := p.db.ExecContext(ctx, DeleteUserQuery, id)
	if err != nil {
		logger.Error(ctx, "Error deleting user: %v", err)
		return false, err
	}

	return true, nil
}
