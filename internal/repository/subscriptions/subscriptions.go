package subscriptions

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

var (
	ErrNoTopBlogs  = errors.New("no blogs found")
	ErrQueryFailed = errors.New("failed to execute GetTopBlogs query")
	ErrScanFailed  = errors.New("failed to scan top blogs row")
)

const (
	GetTopBlogs = `
		SELECT u.user_id, u.name, u.avatar, COUNT(s.follower_id) AS followers_count
		FROM "user" u
		LEFT JOIN subscription s ON s.followed_id = u.user_id
		GROUP BY u.user_id
		ORDER BY followers_count DESC
		LIMIT 5
	`
)

type SubscriptionRepository interface {
	GetTopBlogs(ctx context.Context) ([]models.User, error)
}

type PostgresSubscription struct {
	db *sql.DB
}

func NewPostgresSubscription(db *sql.DB) *PostgresSubscription {
	return &PostgresSubscription{db: db}
}

func (p *PostgresSubscription) GetTopBlogs(ctx context.Context) ([]models.User, error) {
	rows, err := p.db.QueryContext(ctx, GetTopBlogs)
	if err != nil {
		return nil, ErrQueryFailed
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		user := models.User{}

		if err := rows.Scan(&user.Id, &user.Name, &user.Avatar, &user.Subscribers); err != nil {
			return nil, ErrScanFailed
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, ErrNoTopBlogs
	}

	return users, nil
}
