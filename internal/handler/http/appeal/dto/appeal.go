package dto

import "github.com/google/uuid"

type Status string

const (
	StatusSolved  Status = "solved"
	StatusInWork  Status = "in_work"
	StatusCreated Status = "created"
)

type AppealInputDto struct {
	EmailRegistered    string `json:"email_registered"`
	Status             Status `json:"status" default:"created"`
	ProblemDescription string `json:"problem_description"`
	Name               string `json:"name"`
	CategoryID         string `json:"category_id"`
	EmailForConnect    string `json:"email_for_connection"`
	ScreenshotURL      string `json:"screenshot_url"`
}

type AppealOutputDto struct {
	Id                 uuid.UUID `json:"appeal_id"`
	EmailRegistered    string    `json:"email_registered"`
	Status             Status    `json:"status" default:"created"`
	ProblemDescription string    `json:"problem_description"`
	Name               string    `json:"name"`
	CategoryID         int       `json:"category_id"`
	EmailForConnect    string    `json:"email_for_connection"`
	ScreenshotURL      string    `json:"screenshot_url"`
}

type CategoryDto struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}

type CategoryOutputDto struct {
	Categories []CategoryDto `json:"categories"`
}
