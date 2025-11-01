package dto

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"

type ProfileInputDto struct {
	Name     string     `json:"name"`
	Phone    string     `json:"phone"`
	Country  string     `json:"country"`
	Language string     `json:"language"`
	Sex      models.Sex `json:"sex"`
}

type ProfileOutputDto struct {
	Phone       string     `json:"phone"`
	Country     string     `json:"country"`
	Language    string     `json:"language"`
	Sex         models.Sex `json:"sex"`
	DateOfBirth string     `json:"date_of_birth"`
	Age         int        `json:"age"`
	CoverURL    string     `json:"cover_url"`

	//User-model fields
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
