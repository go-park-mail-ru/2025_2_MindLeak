package appeal

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

var (
	ErrGettingCategories = errors.New("error getting categories")
	ErrGettingAppeals    = errors.New("error getting appeals")
	ErrAppealNotFound    = errors.New("appeal not found")
)

const (
	CreateNewAppealQuery = `
INSERT INTO appeal
    (creator_id, email_registered, category_id, problem_description, name, email_for_connect, screenshot_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING 
    appeal_id, creator_id, email_registered, category_id, status,
    problem_description, name, email_for_connect, screenshot_url,
    created_at, updated_at,
    (SELECT category_id FROM appeal_category WHERE appeal_category.category_id = $3) AS cat_id,
    (SELECT name        FROM appeal_category WHERE appeal_category.category_id = $3) AS cat_name
`

	GetAllAppealsQuery = `
SELECT
    a.appeal_id,
    a.creator_id,
    a.email_registered,
    a.category_id,
    a.status,
    a.problem_description,
    a.name,
    a.email_for_connect,
    a.screenshot_url,
    a.created_at,
    a.updated_at,
    c.category_id,
    c.name
FROM appeal a
JOIN appeal_category c ON a.category_id = c.category_id
ORDER BY a.created_at DESC
`

	GetMyAppealsQuery = `
SELECT
    a.appeal_id,
    a.creator_id,
    a.email_registered,
    a.category_id,
    a.status,
    a.problem_description,
    a.name,
    a.email_for_connect,
    a.screenshot_url,
    a.created_at,
    a.updated_at,
    c.category_id,
    c.name
FROM appeal a
JOIN appeal_category c ON a.category_id = c.category_id
WHERE a.creator_id = $1
ORDER BY a.created_at DESC
`

	GetMyAppealQuery = `
SELECT
    a.appeal_id,
    a.creator_id,
    a.email_registered,
    a.category_id,
    a.status,
    a.problem_description,
    a.name,
    a.email_for_connect,
    a.screenshot_url,
    a.created_at,
    a.updated_at,
    c.category_id,
    c.name
FROM appeal a
JOIN appeal_category c ON a.category_id = c.category_id
WHERE a.appeal_id = $1
`

	UpdateMyAppealQuery = `
UPDATE appeal
SET 
    category_id = $2,
    problem_description = $3,
    name = $4,
    email_for_connect = $5,
    screenshot_url = $6,
    updated_at = NOW()
WHERE appeal_id = $1
RETURNING 
    appeal_id, creator_id, email_registered, category_id, status,
    problem_description, name, email_for_connect, screenshot_url,
    created_at, updated_at
`
	GetAllCategoriesQuery = `
SELECT category_id, name
FROM appeal_category
ORDER BY category_id
`
	GetAppealStatsTotalQuery = `
SELECT COUNT(*) FROM appeal;
`

	GetAppealStatsByCategoryQuery = `
SELECT c.name, COUNT(*)
FROM appeal a
JOIN appeal_category c ON a.category_id = c.category_id
GROUP BY c.name;
`

	GetAppealStatsByStatusQuery = `
SELECT status, COUNT(*)
FROM appeal
GROUP BY status;
`
)

type AppealRepository interface {
	CreateNewAppeal(ctx context.Context, appeal models.Appeal) (models.Appeal, error)
	GetAllAppeals(ctx context.Context) ([]models.Appeal, error)
	GetMyAppeals(ctx context.Context, ownerID uuid.UUID) ([]models.Appeal, error)
	GetMyAppeal(ctx context.Context, appealID uuid.UUID) (models.Appeal, error)
	UpdateAppeal(ctx context.Context, appeal models.Appeal) (models.Appeal, error)
	GetAllCategories(ctx context.Context) ([]models.AppealCategory, error)
	GetStats(ctx context.Context) (AppealStatsRaw, error)
}

type PostgresAppeal struct {
	db *sql.DB
}

func NewPostgresAppeal(db *sql.DB) *PostgresAppeal {
	return &PostgresAppeal{
		db: db,
	}
}

func (p *PostgresAppeal) CreateNewAppeal(ctx context.Context, appeal models.Appeal) (models.Appeal, error) {
	var a models.Appeal

	err := p.db.QueryRowContext(
		ctx,
		CreateNewAppealQuery,
		appeal.CreatorID,
		appeal.EmailRegistered,
		appeal.CategoryID,
		appeal.ProblemDescription,
		appeal.Name,
		appeal.EmailForConnect,
		appeal.ScreenshotURL,
	).Scan(
		&a.AppealID,
		&a.CreatorID,
		&a.EmailRegistered,
		&a.CategoryID,
		&a.Status,
		&a.ProblemDescription,
		&a.Name,
		&a.EmailForConnect,
		&a.ScreenshotURL,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.Category.CategoryID,
		&a.Category.Name,
	)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, err
	}

	return a, nil
}

func (p *PostgresAppeal) GetAllAppeals(ctx context.Context) ([]models.Appeal, error) {
	rows, err := p.db.QueryContext(ctx, GetAllAppealsQuery)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, ErrGettingAppeals
	}
	defer rows.Close()

	appeals := make([]models.Appeal, 0)

	for rows.Next() {
		var appeal models.Appeal

		err := rows.Scan(
			&appeal.AppealID,
			&appeal.CreatorID,
			&appeal.EmailRegistered,
			&appeal.CategoryID,
			&appeal.Status,
			&appeal.ProblemDescription,
			&appeal.Name,
			&appeal.EmailForConnect,
			&appeal.ScreenshotURL,
			&appeal.CreatedAt,
			&appeal.UpdatedAt,
			&appeal.Category.CategoryID,
			&appeal.Category.Name,
		)
		if err != nil {
			logger.Error(ctx, err.Error())
			return nil, ErrGettingAppeals
		}

		appeals = append(appeals, appeal)
	}

	return appeals, nil
}

func (p *PostgresAppeal) GetMyAppeals(ctx context.Context, ownerID uuid.UUID) ([]models.Appeal, error) {
	rows, err := p.db.QueryContext(ctx, GetMyAppealsQuery, ownerID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, ErrGettingAppeals
	}
	defer rows.Close()

	appeals := make([]models.Appeal, 0)

	for rows.Next() {
		var appeal models.Appeal

		err := rows.Scan(
			&appeal.AppealID,
			&appeal.CreatorID,
			&appeal.EmailRegistered,
			&appeal.CategoryID,
			&appeal.Status,
			&appeal.ProblemDescription,
			&appeal.Name,
			&appeal.EmailForConnect,
			&appeal.ScreenshotURL,
			&appeal.CreatedAt,
			&appeal.UpdatedAt,
			&appeal.Category.CategoryID,
			&appeal.Category.Name,
		)
		if err != nil {
			logger.Error(ctx, err.Error())
			return nil, ErrGettingAppeals
		}

		appeals = append(appeals, appeal)
	}

	return appeals, nil
}

func (p *PostgresAppeal) GetMyAppeal(ctx context.Context, appealID uuid.UUID) (models.Appeal, error) {
	var appeal models.Appeal

	err := p.db.QueryRowContext(ctx, GetMyAppealQuery, appealID).Scan(
		&appeal.AppealID,
		&appeal.CreatorID,
		&appeal.EmailRegistered,
		&appeal.CategoryID,
		&appeal.Status,
		&appeal.ProblemDescription,
		&appeal.Name,
		&appeal.EmailForConnect,
		&appeal.ScreenshotURL,
		&appeal.CreatedAt,
		&appeal.UpdatedAt,
		&appeal.Category.CategoryID,
		&appeal.Category.Name,
	)

	if err == sql.ErrNoRows {
		return models.Appeal{}, ErrAppealNotFound
	}
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, err
	}

	return appeal, nil
}

func (p *PostgresAppeal) UpdateAppeal(ctx context.Context, appeal models.Appeal) (models.Appeal, error) {
	var updated models.Appeal

	err := p.db.QueryRowContext(ctx,
		UpdateMyAppealQuery,
		appeal.AppealID,
		appeal.CategoryID,
		appeal.ProblemDescription,
		appeal.Name,
		appeal.EmailForConnect,
		appeal.ScreenshotURL,
	).Scan(
		&updated.AppealID,
		&updated.CreatorID,
		&updated.EmailRegistered,
		&updated.CategoryID,
		&updated.Status,
		&updated.ProblemDescription,
		&updated.Name,
		&updated.EmailForConnect,
		&updated.ScreenshotURL,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, err
	}

	return updated, nil
}

func (p *PostgresAppeal) GetAllCategories(ctx context.Context) ([]models.AppealCategory, error) {
	categories := make([]models.AppealCategory, 0)

	rows, err := p.db.QueryContext(ctx, GetAllCategoriesQuery)
	if err != nil {
		logger.Error(ctx, "Error getting all appeal categories: %v", err)
		return nil, ErrGettingCategories
	}
	defer rows.Close()

	for rows.Next() {
		var category models.AppealCategory
		err = rows.Scan(&category.CategoryID, &category.Name)
		if err != nil {
			logger.Error(ctx, "Error getting all appeal categories: %v", err)
			return nil, ErrGettingCategories
		}
		categories = append(categories, category)
	}

	return categories, nil
}

type AppealStatsRaw struct {
	Total      int64
	ByCategory map[string]int64
	ByStatus   map[string]int64
}

func (p *PostgresAppeal) GetStats(ctx context.Context) (AppealStatsRaw, error) {
	stats := AppealStatsRaw{
		ByCategory: make(map[string]int64),
		ByStatus:   make(map[string]int64),
	}

	err := p.db.QueryRowContext(ctx, GetAppealStatsTotalQuery).Scan(&stats.Total)
	if err != nil {
		logger.Error(ctx, err.Error())
		return stats, err
	}

	rows, err := p.db.QueryContext(ctx, GetAppealStatsByCategoryQuery)
	if err != nil {
		logger.Error(ctx, err.Error())
		return stats, err
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		var count int64
		err := rows.Scan(&category, &count)
		if err != nil {
			logger.Error(ctx, err.Error())
			return stats, err
		}
		stats.ByCategory[category] = count
	}

	rows, err = p.db.QueryContext(ctx, GetAppealStatsByStatusQuery)
	if err != nil {
		logger.Error(ctx, err.Error())
		return stats, err
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		var count int64
		err := rows.Scan(&status, &count)
		if err != nil {
			logger.Error(ctx, err.Error())
			return stats, err
		}
		stats.ByStatus[status] = count
	}

	return stats, nil
}
