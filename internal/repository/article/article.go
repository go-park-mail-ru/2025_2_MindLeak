package article

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

var (
	ErrArticleExists   = errors.New("this article already exists")
	ErrArticleNotFound = errors.New("article not found")
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, authorId uuid.UUID, title, content string, topicId int) (models.Article, error)
	GetArticleById(ctx context.Context, id uuid.UUID) (models.Article, error)
	GetArticlesByAuthorId(ctx context.Context, authorId uuid.UUID) ([]models.Article, error)
	GetFeedArticles(ctx context.Context, feed models.Feed) ([]models.Article, error)
	DeleteArticle(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateArticle(ctx context.Context, article models.Article) (models.Article, error)
	GetArticlesByTopic(ctx context.Context, topicTitle string, offset int) ([]models.Article, error)
}

type ArticleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) *ArticleRepo {
	return &ArticleRepo{db: db}
}

func (r *ArticleRepo) CreateArticle(ctx context.Context, authorID uuid.UUID, title, content string, topicID int) (models.Article, error) {
	query := `
		INSERT INTO article (author_id, title, content, topic_id, status)
		VALUES ($1, $2, $3, $4, 'draft')
		RETURNING article_id, author_id, title, content, topic_id, status, created_at, updated_at
	`

	var a models.Article
	err := r.db.QueryRowContext(ctx, query, authorID, title, content, topicID).Scan(
		&a.ID, &a.AuthorID, &a.Title, &a.Content, &a.Topic.TopicId,
		&a.Status, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Article{}, fmt.Errorf("create article: %w", err)
	}

	if err := r.loadTopic(ctx, &a); err != nil {
		logger.Error(ctx, err.Error())

		return models.Article{}, fmt.Errorf("load topic: %w", err)
	}
	if err := r.loadAuthor(ctx, &a); err != nil {
		logger.Error(ctx, err.Error())
		return models.Article{}, fmt.Errorf("load author: %w", err)
	}

	return a, nil
}

func (r *ArticleRepo) GetArticleById(ctx context.Context, id uuid.UUID) (models.Article, error) {
	query := `
		SELECT a.article_id, a.author_id, a.title, a.content, 
		       a.status, a.created_at, a.updated_at,
		       t.topic_id, t.title AS topic_title,
		       u.name, u.avatar
		FROM article a
		JOIN topic t ON a.topic_id = t.topic_id
		JOIN "user" u ON a.author_id = u.user_id
		WHERE a.article_id = $1
	`

	var a models.Article
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.AuthorID, &a.Title, &a.Content,
		&a.Status, &a.CreatedAt, &a.UpdatedAt,
		&a.Topic.TopicId, &a.Topic.Title,
		&a.AuthorName, &a.AuthorAvatar,
	)
	if errors.Is(err, sql.ErrNoRows) {
		logger.Error(ctx, err.Error())
		return models.Article{}, ErrArticleNotFound
	}
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Article{}, fmt.Errorf("get article by id: %w", err)
	}

	return a, nil
}

func (r *ArticleRepo) GetArticlesByAuthorId(ctx context.Context, authorID uuid.UUID) ([]models.Article, error) {
	query := `
		SELECT a.article_id, a.title, a.content, 
		       a.status, a.created_at, a.updated_at,
		       t.topic_id, t.title AS topic_title,
		       u.name, u.avatar
		FROM article a
		JOIN topic t ON a.topic_id = t.topic_id
		JOIN "user" u ON a.author_id = u.user_id
		WHERE a.author_id = $1
		ORDER BY a.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, authorID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, fmt.Errorf("get articles by author: %w", err)
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		if err := rows.Scan(
			&a.ID, &a.Title, &a.Content,
			&a.Status, &a.CreatedAt, &a.UpdatedAt,
			&a.Topic.TopicId, &a.Topic.Title,
			&a.AuthorName, &a.AuthorAvatar,
		); err != nil {
			logger.Error(ctx, err.Error())
			return nil, err
		}
		a.AuthorID = authorID
		articles = append(articles, a)
	}

	return articles, rows.Err()
}

func (r *ArticleRepo) GetFeedArticles(ctx context.Context, feed models.Feed) ([]models.Article, error) {
	query := `
		SELECT a.article_id, a.author_id, a.title, a.content, a.media_url,
		       a.status, a.created_at, a.updated_at,
		       t.topic_id, t.title AS topic_title,
		       u.name, u.avatar
		FROM article a
		JOIN topic t ON a.topic_id = t.topic_id
		JOIN "user" u ON a.author_id = u.user_id
		WHERE a.status = 'draft'
		ORDER BY a.updated_at DESC NULLS LAST, a.created_at DESC
		OFFSET $1 LIMIT 5
	`

	rows, err := r.db.QueryContext(ctx, query, feed.Offset)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, fmt.Errorf("get feed articles: %w", err)
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		var mediaURL sql.NullString

		if err := rows.Scan(
			&a.ID, &a.AuthorID, &a.Title, &a.Content, &mediaURL,
			&a.Status, &a.CreatedAt, &a.UpdatedAt,
			&a.Topic.TopicId, &a.Topic.Title,
			&a.AuthorName, &a.AuthorAvatar,
		); err != nil {
			logger.Error(ctx, err.Error())
			return nil, err
		}
		if mediaURL.Valid {
			a.MediaURL = mediaURL.String
		} else {
			a.MediaURL = ""
		}
		articles = append(articles, a)
	}

	return articles, rows.Err()
}

func (r *ArticleRepo) GetArticlesByTopic(ctx context.Context, topicTitle string, offset int) ([]models.Article, error) {
	query := `
		SELECT a.article_id, a.author_id, a.title, a.content, 
		       a.status, a.created_at, a.updated_at,
		       t.topic_id, t.title AS topic_title,
		       u.name, u.avatar
		FROM article a
		JOIN topic t ON a.topic_id = t.topic_id
		JOIN "user" u ON a.author_id = u.user_id
		WHERE t.title = $1
		ORDER BY a.created_at DESC
		OFFSET $2 LIMIT 5
	`

	rows, err := r.db.QueryContext(ctx, query, topicTitle, offset)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, fmt.Errorf("get articles by topic: %w", err)
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		if err := rows.Scan(
			&a.ID, &a.AuthorID, &a.Title, &a.Content,
			&a.Status, &a.CreatedAt, &a.UpdatedAt,
			&a.Topic.TopicId, &a.Topic.Title,
			&a.AuthorName, &a.AuthorAvatar,
		); err != nil {
			logger.Error(ctx, err.Error())
			return nil, err
		}
		articles = append(articles, a)
	}

	return articles, rows.Err()
}

func (r *ArticleRepo) DeleteArticle(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `DELETE FROM article WHERE article_id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error(ctx, err.Error())
		return false, fmt.Errorf("delete article: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error(ctx, err.Error())
		return false, fmt.Errorf("rows affected: %w", err)
	}

	return rowsAffected > 0, nil
}

func (r *ArticleRepo) UpdateArticle(ctx context.Context, article models.Article) (models.Article, error) {
	query := `
		UPDATE article
		SET 
		    title = COALESCE($1, title),
		    content = COALESCE($2, content),
		    media_url = COALESCE($3, media_url),
		    status = COALESCE($4, status),
		    updated_at = CURRENT_TIMESTAMP
		WHERE article_id = $5 AND author_id = $6
		RETURNING 
		    article_id, author_id, title, content, media_url, topic_id, status,
		    comments_count, reposts_count, views_count, created_at, updated_at
	`

	var updated models.Article
	err := r.db.QueryRowContext(ctx, query,
		article.Title, article.Content, article.MediaURL, article.Status,
		article.ID, article.AuthorID,
	).Scan(
		&updated.ID, &updated.AuthorID, &updated.Title, &updated.Content, &updated.MediaURL,
		&updated.TopicID, &updated.Status,
		&updated.CommentsCount, &updated.RepostsCount, &updated.ViewsCount,
		&updated.CreatedAt, &updated.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Article{}, ErrArticleNotFound
	}
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Article{}, fmt.Errorf("update article: %w", err)
	}

	if err := r.loadTopic(ctx, &updated); err != nil {
		logger.Error(ctx, err.Error())
		return models.Article{}, fmt.Errorf("load topic: %w", err)
	}
	if err := r.loadAuthor(ctx, &updated); err != nil {
		logger.Error(ctx, err.Error())
		return models.Article{}, fmt.Errorf("load author: %w", err)
	}

	return updated, nil
}

func (r *ArticleRepo) loadAuthor(ctx context.Context, a *models.Article) error {
	query := `SELECT name, avatar FROM "user" WHERE user_id = $1`
	return r.db.QueryRowContext(ctx, query, a.AuthorID).Scan(&a.AuthorName, &a.AuthorAvatar)
}

func (r *ArticleRepo) loadTopic(ctx context.Context, a *models.Article) error {
	query := `SELECT title FROM topic WHERE topic_id = $1`
	return r.db.QueryRowContext(ctx, query, a.Topic.TopicId).Scan(&a.Topic.Title)
}
