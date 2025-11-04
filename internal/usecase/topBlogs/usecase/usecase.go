package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/subscriptions"
	subsRepo "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/subscriptions"
)

var (
	NoTopBlogs  = errors.New("no top blogs found")
	QueryFailed = errors.New("query failed")
	ScanFailed  = errors.New("scan failed")
	ServerError = errors.New("internal apiserver error")
)

type Usecase struct {
	subscriptionsRepo subscriptions.SubscriptionRepository
}

func NewTopBlogsUsecase(subsRepo subscriptions.SubscriptionRepository) *Usecase {
	return &Usecase{subscriptionsRepo: subsRepo}
}

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, subsRepo.ErrNoTopBlogs):
		return NoTopBlogs
	case errors.Is(err, subsRepo.ErrQueryFailed):
		return QueryFailed
	case errors.Is(err, subsRepo.ErrScanFailed):
		return ScanFailed
	default:
		return ServerError
	}
}
