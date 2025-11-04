package dto

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"

type ProfileInputDto struct {
	Phone       string     `json:"phone"`
	Country     string     `json:"country"`
	Language    string     `json:"language"`
	Sex         models.Sex `json:"sex"`
	DateOfBirth string     `json:"date_of_birth"`
	Age         int        `json:"age"`
	Description string     `json:"description"`
	Cover       string     `json:"cover"`

	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ProfileOutputDto struct {
	Phone       string     `json:"phone"`
	Country     string     `json:"country"`
	Language    string     `json:"language"`
	Sex         models.Sex `json:"sex"`
	DateOfBirth string     `json:"date_of_birth"`
	Age         int        `json:"age"`
	Description string     `json:"description"`
	CoverURL    string     `json:"cover_url"`

	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
