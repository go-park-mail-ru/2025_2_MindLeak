package dto

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"time"
)

type ProfileDto struct {
	Phone       string
	Country     string
	Language    string
	Sex         models.Sex
	DateOfBirth string
	Age         int
	CoverURL    string
	CreatedAt   time.Time

	Name   string
	Avatar string
	Email  string
}
