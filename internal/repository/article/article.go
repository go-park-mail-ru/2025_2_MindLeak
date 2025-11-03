package article

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
		RETURNING article_id, author_id, title, content, created_at, updated_at, status
	`

	var a models.Article
	err := r.db.QueryRowContext(ctx, query, authorID, title, content).Scan(
		&a.ID, &a.AuthorID, &a.Title, &a.Content, &a.CreatedAt, &a.UpdatedAt, &a.Status,
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
		SELECT a.article_id, a.author_id, a.title, a.content, a.image_url, a.created_at, a.updated_at, a.status,
		       u.name, u.avatar
		FROM article a
		JOIN "user" u ON a.author_id = u.user_id
		WHERE a.article_id = $1
	`

	var a models.Article
	var authorName, avatar string
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.AuthorID, &a.Title, &a.Content, &a.ImageURL, &a.CreatedAt, &a.UpdatedAt, &a.Status,
		&authorName, &avatar,
	)
	if err == pgx.ErrNoRows {
		return nil, ErrArticleNotFound
	}
	if err != nil {
		return nil, err
	}

	a.AuthorName = authorName
	a.AuthorAvatar = avatar

	return &a, nil
}

func (r *ArticleRepo) GetArticlesByAuthorId(ctx context.Context, authorID uuid.UUID) ([]*models.Article, error) {
	query := `
		SELECT a.article_id, a.title, a.content, a.image_url, a.created_at, a.updated_at, a.status,
		       u.name, u.avatar
		FROM article a
		JOIN "user" u ON a.author_id = u.user_id
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
		var authorName, avatar string
		if err := rows.Scan(
			&a.ID, &a.Title, &a.Content, &a.ImageURL, &a.CreatedAt, &a.UpdatedAt, &a.Status,
			&authorName, &avatar,
		); err != nil {
			return nil, err
		}

		a.AuthorID = authorID
		a.AuthorName = authorName
		a.AuthorAvatar = avatar

		articles = append(articles, &a)
	}

	return articles, rows.Err()
}

func (r *ArticleRepo) GetFeedArticles(ctx context.Context, feed models.Feed) ([]*models.Article, error) {
	query := `
		SELECT a.article_id, a.author_id, a.title, a.content, a.image_url, a.created_at, a.updated_at, a.status,
		       u.name, u.avatar
		FROM article a
		JOIN "user" u ON a.author_id = u.user_id
		WHERE a.status = 'published'
		ORDER BY a.updated_at DESC NULLS LAST, a.created_at DESC
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
		var authorName, avatar string
		if err := rows.Scan(
			&a.ID, &a.AuthorID, &a.Title, &a.Content, &a.ImageURL, &a.CreatedAt, &a.UpdatedAt, &a.Status,
			&authorName, &avatar,
		); err != nil {
			return nil, err
		}

		a.AuthorName = authorName
		a.AuthorAvatar = avatar
		articles = append(articles, &a)
	}

	return articles, rows.Err()
}

func (r *ArticleRepo) Update(ctx context.Context, article *models.Article) (*models.Article, error) {
	query := `
        UPDATE article
        SET 
            title = COALESCE($1, title),
            content = COALESCE($2, content),
            image_url = COALESCE($3, image_url),
            status = COALESCE($4, status),
            updated_at = CURRENT_TIMESTAMP
        WHERE article_id = $5 AND author_id = $6
        RETURNING 
            article_id,
            author_id,
            title,
            content,
            image_url,
            status,
            created_at,
            updated_at
    `

	var updated models.Article
	err := r.db.QueryRowContext(ctx, query,
		article.Title,
		article.Content,
		article.ImageURL,
		article.Status,
		article.ID,
		article.AuthorID,
	).Scan(
		&updated.ID,
		&updated.AuthorID,
		&updated.Title,
		&updated.Content,
		&updated.ImageURL,
		&updated.Status,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrArticleNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("update article: %w", err)
	}

	if err := r.loadAuthor(ctx, &updated); err != nil {
		return nil, fmt.Errorf("load author: %w", err)
	}

	return &updated, nil
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
		SELECT u.name, u.avatar
		FROM "user" u
		WHERE u.user_id = $1
	`

	var authorName, avatar string
	err := r.db.QueryRowContext(ctx, query, a.AuthorID).Scan(&authorName, &avatar)
	if err != nil {
		return err
	}

	a.AuthorName = authorName
	a.AuthorAvatar = avatar

	return nil
}
