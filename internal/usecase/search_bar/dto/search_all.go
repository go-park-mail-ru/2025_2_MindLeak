package dto

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"

type SearchAll struct {
	Users    []models.User
	Articles []models.Article
}
