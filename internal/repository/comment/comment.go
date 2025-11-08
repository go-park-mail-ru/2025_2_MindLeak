package comment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

var (
	ErrCreatingComment       = errors.New("error creating comment")
	ErrUpdatingComment       = errors.New("error updating comment")
	ErrDeletingComment       = errors.New("error deleting comment")
	ErrGettingComment        = errors.New("error getting profile")
	ErrResolvingAuthor       = errors.New("error resolving author")
	ErrResolvingArticleTitle = errors.New("error resolving article title")
)

const (
	CreateCommentQuery = `
		INSERT INTO comment (article_id, user_id, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING comment_id, article_id, user_id, content, created_at, updated_at
	`

	CreateResponseCommentQuery = `
		INSERT INTO comment (article_id, user_id, content, created_at, updated_at, reply_to)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING comment_id, article_id, user_id, content, created_at, updated_at, reply_to
	`

	DeleteCommentQuery = `
		DELETE FROM comment WHERE comment_id = $1
	`

	UpdateCommentQuery = `
		UPDATE comment SET
		                   content = COALESCE ($2, content),
		                   updated_at = CURRENT_TIMESTAMP
		WHERE comment_id = $1
		RETURNING comment_id, article_id, user_id, content, created_at, updated_at, reply_to
	`
	GetByAuthorQuery = `
		SELECT c.comment_id,
		       c.article_id,
		       c.user_id,
		       c.content,
		       c.created_at,
		       c.updated_at,
		       c.reply_to,
		       u.name,
		       u.avatar,
		       a.title
		FROM comment c JOIN "user" u ON c.user_id = u.user_id JOIN article a ON c.article_id = a.article_id
		WHERE c.user_id = $1
		ORDER BY c.created_at DESC
	`

	GetByArticleQuery = `
		SELECT c.comment_id,
		       c.article_id,
		       c.user_id,
		       c.content,
		       c.created_at,
		       c.updated_at,
		       c.reply_to,
		       u.name,
		       u.avatar,
		       a.title
		FROM comment c JOIN "user" u ON c.user_id = u.user_id JOIN article a ON c.article_id = a.article_id
		WHERE c.article_id = $1
		ORDER BY c.created_at DESC
	`
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment models.Comment) (models.Comment, error)
	GetCommentsByAuthor(ctx context.Context, authorId uuid.UUID) ([]models.Comment, error)
	GetCommentsByArticle(ctx context.Context, articleId uuid.UUID) ([]models.Comment, error)
	UpdateComment(ctx context.Context, comment models.Comment) (models.Comment, error)
	DeleteComment(ctx context.Context, commentId uuid.UUID) (bool, error)
}

type PostgresComment struct {
	db *sql.DB
}

func NewPostgresComment(db *sql.DB) *PostgresComment {
	return &PostgresComment{db: db}
}

func (r *PostgresComment) CreateComment(ctx context.Context, comment models.Comment) (models.Comment, error) {
	var model models.Comment

	//var err error

	var replyTo sql.NullString

	if comment.ReplyTo != nil {
		err := r.db.QueryRowContext(
			ctx,
			CreateResponseCommentQuery,
			comment.ArticleId,
			comment.UserId,
			comment.Content,
			comment.CreatedAt,
			comment.UpdatedAt,
			comment.ReplyTo,
		).Scan(
			&model.Id,
			&model.ArticleId,
			&model.UserId,
			&model.Content,
			&model.CreatedAt,
			&model.UpdatedAt,
			&replyTo,
		)
		if err != nil {
			logger.Error(ctx, err.Error())
			return models.Comment{}, ErrCreatingComment
		}
		if replyTo.Valid {
			u, _ := uuid.Parse(replyTo.String)
			model.ReplyTo = &u
		} else {
			model.ReplyTo = nil
		}
	} else {
		err := r.db.QueryRowContext(
			ctx,
			CreateCommentQuery,
			comment.ArticleId,
			comment.UserId,
			comment.Content,
			comment.CreatedAt,
			comment.UpdatedAt,
		).Scan(
			&model.Id,
			&model.ArticleId,
			&model.UserId,
			&model.Content,
			&model.CreatedAt,
			&model.UpdatedAt,
		)
		if err != nil {
			logger.Error(ctx, err.Error())
			return models.Comment{}, ErrCreatingComment
		}

	}

	if err := r.resolveAuthor(ctx, &model); err != nil {
		logger.Error(ctx, err.Error())
		return models.Comment{}, ErrResolvingAuthor
	}

	if err := r.resolveArticleTitle(ctx, &model); err != nil {
		logger.Error(ctx, err.Error())
		return models.Comment{}, ErrResolvingArticleTitle
	}

	return model, nil
}

func (r *PostgresComment) GetCommentsByAuthor(ctx context.Context, authorId uuid.UUID) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := r.db.QueryContext(ctx, GetByAuthorQuery, authorId)

	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, ErrGettingComment
	}

	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var replyTo sql.NullString
		if err := rows.Scan(
			&comment.Id,
			&comment.ArticleId,
			&comment.UserId,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&replyTo,
			&comment.AuthorName,
			&comment.AuthorAvatar,
			&comment.ArticleTitle,
		); err != nil {
			logger.Error(ctx, err.Error())
			return nil, err
		}

		if replyTo.Valid {
			u, _ := uuid.Parse(replyTo.String)
			comment.ReplyTo = &u
		} else {
			comment.ReplyTo = nil
		}

		comments = append(comments, comment)

	}

	return comments, rows.Err()
}

func (r *PostgresComment) GetCommentsByArticle(ctx context.Context, articleId uuid.UUID) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := r.db.QueryContext(ctx, GetByArticleQuery, articleId)

	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, fmt.Errorf("get comments by article: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var replyTo sql.NullString
		if err := rows.Scan(
			&comment.Id,
			&comment.ArticleId,
			&comment.UserId,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&replyTo,
			&comment.AuthorName,
			&comment.AuthorAvatar,
			&comment.ArticleTitle,
		); err != nil {
			logger.Error(ctx, err.Error())
			return nil, err
		}

		if replyTo.Valid {
			u, _ := uuid.Parse(replyTo.String)
			comment.ReplyTo = &u
		} else {
			comment.ReplyTo = nil
		}

		comments = append(comments, comment)

	}

	return comments, rows.Err()
}

func (r *PostgresComment) UpdateComment(ctx context.Context, comment models.Comment) (models.Comment, error) {
	var updated models.Comment

	var replyTo sql.NullString

	err := r.db.QueryRowContext(ctx, UpdateCommentQuery,
		comment.Id,
		comment.Content,
	).Scan(
		&updated.Id,
		&updated.ArticleId,
		&updated.UserId,
		&updated.Content,
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&replyTo,
	)

	if err != nil {
		logger.Error(ctx, "update comment: %w", err)
		return models.Comment{}, ErrUpdatingComment
	}

	if err := r.resolveAuthor(ctx, &updated); err != nil {
		logger.Error(ctx, err.Error())
		return models.Comment{}, ErrResolvingAuthor
	}

	if err := r.resolveArticleTitle(ctx, &updated); err != nil {
		logger.Error(ctx, err.Error())
		return models.Comment{}, ErrResolvingArticleTitle
	}

	if replyTo.Valid {
		u, _ := uuid.Parse(replyTo.String)
		updated.ReplyTo = &u
	} else {
		comment.ReplyTo = nil
	}

	return updated, nil
}

func (r *PostgresComment) DeleteComment(ctx context.Context, commentId uuid.UUID) (bool, error) {

	_, err := r.db.ExecContext(ctx, DeleteCommentQuery, commentId)
	if err != nil {
		return false, ErrDeletingComment
	}

	return true, nil
}

func (r *PostgresComment) resolveAuthor(ctx context.Context, model *models.Comment) error {
	query := `SELECT name, avatar FROM "user" WHERE user_id = $1`
	return r.db.QueryRowContext(ctx, query, model.UserId).Scan(&model.AuthorName, &model.AuthorAvatar)
}

func (r *PostgresComment) resolveArticleTitle(ctx context.Context, model *models.Comment) error {
	query := `SELECT title FROM article WHERE article_id = $1`
	return r.db.QueryRowContext(ctx, query, model.ArticleId).Scan(&model.ArticleTitle)
}
