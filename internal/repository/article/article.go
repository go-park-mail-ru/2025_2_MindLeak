package article

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

var (
	ErrArticleExists   = errors.New("this article is already exists")
	ErrArticleNotFound = errors.New("article not found")
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, authorId uuid.UUID, title, content string) (*models.Article, error)
	GetArticleById(ctx context.Context, id uuid.UUID) (*models.Article, error)
	GetArticlesByAuthorId(ctx context.Context, authorId uuid.UUID) ([]*models.Article, error)
	GetFeedArticles(ctx context.Context, feed models.Feed) ([]*models.Article, error)
	DeleteArticle(ctx context.Context, id uuid.UUID) (bool, error)
}

type Article struct {
	Id           uuid.UUID
	AuthorId     uuid.UUID
	Title        string
	Content      string
	CreatedAt    time.Time
	Image        string
	AuthorName   string
	AuthorAvatar string
}

type ArticleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) *ArticleRepo {
	return &ArticleRepo{db: db}
}

func (r *ArticleRepo) CreateArticle(ctx context.Context, authorID uuid.UUID, title, content string) (*models.Article, error) {
	query := `
		INSERT INTO article (author_id, title, content, status)
		VALUES ($1, $2, $3, 'draft')
		RETURNING article_id, author_id, title, content, created_at, published_at, status
	`

	var a models.Article
	err := r.db.QueryRowContext(ctx, query, authorID, title, content).Scan(
		&a.ID, &a.AuthorID, &a.Title, &a.Content, &a.PublishedAt, &a.Status,
	)
	if err != nil {
		return nil, err
	}

	if err := r.loadAuthor(ctx, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *ArticleRepo) GetArticleById(ctx context.Context, id uuid.UUID) (*models.Article, error) {
	query := `
		SELECT a.article_id, a.author_id, a.title, a.content, a.created_at, a.published_at, a.status,
		       u.login, up.display_name, up.avatar_url
		FROM article a
		JOIN user u ON a.author_id = u.user_id
		LEFT JOIN user_profile up ON u.user_id = up.user_id
		WHERE a.article_id = $1
	`

	var a models.Article
	var authorName, displayName, avatarURL string
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.AuthorID, &a.Title, &a.Content, &a.PublishedAt, &a.PublishedAt, &a.Status,
		&authorName, &displayName, &avatarURL,
	)
	if err == pgx.ErrNoRows {
		return nil, ErrArticleNotFound
	}
	if err != nil {
		return nil, err
	}

	a.AuthorName = displayName
	if a.AuthorName == "" {
		a.AuthorName = authorName
	}
	a.AuthorAvatar = avatarURL

	return &a, nil
}

func (r *ArticleRepo) GetArticlesByAuthorId(ctx context.Context, authorID uuid.UUID) ([]*models.Article, error) {
	query := `
		SELECT a.article_id, a.title, a.content, a.created_at, a.published_at, a.status,
		       u.login, up.display_name, up.avatar_url
		FROM article a
		JOIN user u ON a.author_id = u.user_id
		LEFT JOIN user_profile up ON u.user_id = up.user_id
		WHERE a.author_id = $1
		ORDER BY a.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*models.Article
	for rows.Next() {
		var a models.Article
		var authorName, displayName, avatarURL string
		if err := rows.Scan(
			&a.ID, &a.Title, &a.Content, &a.PublishedAt, &a.PublishedAt, &a.Status,
			&authorName, &displayName, &avatarURL,
		); err != nil {
			return nil, err
		}

		a.AuthorID = authorID
		a.AuthorName = displayName
		if a.AuthorName == "" {
			a.AuthorName = authorName
		}
		a.AuthorAvatar = avatarURL

		articles = append(articles, &a)
	}

	return articles, rows.Err()
}

func (r *ArticleRepo) GetFeedArticles(ctx context.Context, feed models.Feed) ([]*models.Article, error) {
	query := `
		SELECT a.article_id, a.author_id, a.title, a.content, a.created_at, a.published_at, a.status,
		       u.login, up.display_name, up.avatar_url
		FROM article a
		JOIN user u ON a.author_id = u.user_id
		LEFT JOIN user_profile up ON u.user_id = up.user_id
		WHERE a.status = 'published'
		ORDER BY a.published_at DESC NULLS LAST, a.created_at DESC
		OFFSET $1 LIMIT 5
	`

	rows, err := r.db.QueryContext(ctx, query, feed.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*models.Article
	for rows.Next() {
		var a models.Article
		var authorName, displayName, avatarURL string
		if err := rows.Scan(
			&a.ID, &a.AuthorID, &a.Title, &a.Content, &a.PublishedAt, &a.Status,
			&authorName, &displayName, &avatarURL,
		); err != nil {
			return nil, err
		}

		a.AuthorName = displayName
		if a.AuthorName == "" {
			a.AuthorName = authorName
		}
		a.AuthorAvatar = avatarURL

		articles = append(articles, &a)
	}

	return articles, rows.Err()
}

func (r *ArticleRepo) Update(ctx context.Context, article *models.Article) error {
	query := `
        UPDATE article
        SET 
            title = COALESCE($1, title),
            content = COALESCE($2, content),
            image_url = COALESCE($3, image_url),
            status = COALESCE($4, status),
            updated_at = CURRENT_TIMESTAMP
        WHERE article_id = $5 AND author_id = $6
        RETURNING updated_at
    `

	var updatedAt time.Time
	err := r.db.QueryRowContext(ctx, query,
		article.Title,
		article.Content,
		article.ImageURL,
		article.Status,
		article.ID,
		article.AuthorID,
	).Scan(&updatedAt)

	if err == pgx.ErrNoRows {
		return ErrArticleNotFound
	}
	if err != nil {
		return fmt.Errorf("update article: %w", err)
	}

	article.UpdatedAt = updatedAt
	return nil
}

func (r *ArticleRepo) DeleteArticle(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM article WHERE article_id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("delete article: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("delete article rows affected: %w", err)
	}

	return rowsAffected > 0, nil
}

func (r *ArticleRepo) loadAuthor(ctx context.Context, a *models.Article) error {
	query := `
		SELECT u.login, up.display_name, up.avatar_url
		FROM user u
		LEFT JOIN user_profile up ON u.user_id = up.user_id
		WHERE u.user_id = $1
	`

	var authorName, displayName, avatarURL string
	err := r.db.QueryRowContext(ctx, query, a.AuthorID).Scan(&authorName, &displayName, &avatarURL)
	if err != nil {
		return err
	}

	a.AuthorName = displayName
	if a.AuthorName == "" {
		a.AuthorName = authorName
	}
	a.AuthorAvatar = avatarURL

	return nil
}
