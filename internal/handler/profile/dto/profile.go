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
	Password    string     `json:"password"`
	Cover       string     `json:"cover_url"`

	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
	Email  string `json:"email"`
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
	CreatedAt   string     `json:"created_at"`

	Name          string `json:"name"`
	AvatarURL     string `json:"avatar_url"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Subscribers   int    `json:"subscribers"`
	Subscriptions int    `json:"subscriptions"`
}
