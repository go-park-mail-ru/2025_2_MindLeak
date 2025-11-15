package appeal

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
)

const (
	CreateNewAppelQuery = ``
	GetAllAppealsQuery  = ``
	GetMyAppealsQuery   = ``
)

type AppealRepository interface {
	CreateNewAppeal(ctx context.Context, appeal models.Appeal) (models.Appeal, error)
	GetAllAppeals(ctx context.Context)
	GetMyAppeals(ctx context.Context)
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

	err := p.db.QueryRowContext(ctx, CreateNewAppelQuery, appeal.CreatorID, appeal.EmailRegistered, appeal.Status, appeal.ProblemDescription, appeal.Name, appeal.EmailForConnection, appeal.ScreenshotUrl).Scan(
		&a.CreatorID,
		&a.EmailRegistered,
		&a.Status,
		&a.ProblemDescription,
		&a.Name,
		&a.EmailForConnection,
		&a.ScreenshotUrl,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, err
	}

	return appeal, nil

}

func (p *PostgresAppeal) GetMyAppeals(ctx context.Context) {

}

func (p *PostgresAppeal) GetAllAppeals(ctx context.Context) {

}

func (p *PostgresAppeal) GetMyAppeal(ctx context.Context) {}

func (p *PostgresAppeal) Get
