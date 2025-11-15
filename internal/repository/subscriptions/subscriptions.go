package subscriptions

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

var (
	ErrNoTopBlogs  = errors.New("no blogs found")
	ErrQueryFailed = errors.New("failed to execute GetTopBlogs query")
	ErrScanFailed  = errors.New("failed to scan top blogs row")
)

const (
	GetTopBlogsQuery = `
		SELECT u.user_id, u.name, u.avatar, COUNT(s.follower_id) AS followers_count
		FROM "user" u
		LEFT JOIN subscription s ON s.followed_id = u.user_id
		GROUP BY u.user_id
		ORDER BY followers_count DESC
		LIMIT 5
	`

	GetSubscribersQuery = `        
		SELECT 
            u.user_id,
            u.email,
            u.password,
            u.name,
            u.avatar,
            COALESCE(subs.count, 0) AS subscriptions,
            COALESCE(follow.count, 0) AS subscriptions
        FROM subscription s
        JOIN "user" u ON u.user_id = s.follower_id
        LEFT JOIN (
            SELECT followed_id, COUNT(*) AS count
            FROM subscription
            GROUP BY followed_id
        ) subs ON subs.followed_id = u.user_id
        LEFT JOIN (
            SELECT follower_id, COUNT(*) AS count
            FROM subscription
            GROUP BY follower_id
        ) follow ON follow.follower_id = u.user_id
        WHERE s.followed_id = $1`

	GetSubscriptionsQuery = `SELECT 
			u.user_id,
			u.email,
			u.password,
			u.name,
			u.avatar,
			COALESCE(subs.count, 0) AS subscriptions,
			COALESCE(follow.count, 0) AS subscriptions
		FROM subscription s
		JOIN "user" u ON u.user_id = s.followed_id
		LEFT JOIN (
			SELECT followed_id, COUNT(*) AS count
			FROM subscription
			GROUP BY followed_id
		) subs ON subs.followed_id = u.user_id
		LEFT JOIN (
			SELECT follower_id, COUNT(*) AS count
			FROM subscription
			GROUP BY follower_id
		) follow ON follow.follower_id = u.user_id
		WHERE s.follower_id = $1`

	subscribeQuery = `
		INSERT INTO subscription (follower_id, followed_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
		`

	existsQuery = `
		SELECT 1 FROM subscription
		WHERE follower_id=$1 AND followed_id=$2
		`

	unsubscribeQuery = `DELETE FROM subscription WHERE follower_id=$1 AND followed_id=$2;`
)

type SubscriptionRepository interface {
	GetTopBlogs(ctx context.Context) ([]models.User, error)
	Unsubscribe(ctx context.Context, followerID uuid.UUID, targetID uuid.UUID) error
	GetSubscribers(ctx context.Context, userID uuid.UUID) ([]models.User, error)
	GetSubscriptions(ctx context.Context, userID uuid.UUID) ([]models.User, error)
	Subscribe(ctx context.Context, followerID, targetID uuid.UUID) error
	Exists(ctx context.Context, followerID, targetID uuid.UUID) (bool, error)
}

type PostgresSubscription struct {
	db *sql.DB
}

func NewPostgresSubscription(db *sql.DB) *PostgresSubscription {
	return &PostgresSubscription{db: db}
}

func (p *PostgresSubscription) GetTopBlogs(ctx context.Context) ([]models.User, error) {
	rows, err := p.db.QueryContext(ctx, GetTopBlogsQuery)
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

func (p *PostgresSubscription) Subscribe(ctx context.Context, followerID, targetID uuid.UUID) error {
	_, err := p.db.ExecContext(ctx, subscribeQuery, followerID, targetID)
	if err != nil {
		return ErrQueryFailed
	}

	return err
}

func (p *PostgresSubscription) Unsubscribe(ctx context.Context, followerID uuid.UUID, targetID uuid.UUID) error {
	_, err := p.db.ExecContext(ctx, unsubscribeQuery, followerID, targetID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresSubscription) GetSubscribers(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	rows, err := p.db.QueryContext(ctx, GetSubscribersQuery, userID)
	if err != nil {
		return nil, ErrQueryFailed
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		user := models.User{}

		err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.Avatar,
			&user.Subscribers,
			&user.Subscriptions,
		)
		if err != nil {
			return nil, ErrScanFailed
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return []models.User{}, nil
	}

	return users, nil
}

func (p *PostgresSubscription) GetSubscriptions(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	rows, err := p.db.QueryContext(ctx, GetSubscriptionsQuery, userID)
	if err != nil {
		return nil, ErrQueryFailed
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		user := models.User{}

		err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.Name,
			&user.Avatar,
			&user.Subscribers,
			&user.Subscriptions,
		)
		if err != nil {
			return nil, ErrScanFailed
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return []models.User{}, nil
	}

	return users, nil
}

func (p *PostgresSubscription) Exists(ctx context.Context, followerID, targetID uuid.UUID) (bool, error) {
	var isExists int
	err := p.db.QueryRowContext(ctx, existsQuery, followerID, targetID).Scan(&isExists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
