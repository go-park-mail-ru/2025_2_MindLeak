package models

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusSolved  Status = "solved"
	StatusInWork  Status = "in_work"
	StatusCreated Status = "created"
)

type AppealCategory struct {
	CategoryID uuid.UUID `db:"category_id"`
	Name       string    `db:"name"`
}

type Appeal struct {
	AppealID           uuid.UUID `db:"appeal_id"`
	CreatorID          uuid.UUID `db:"creator_id"`
	EmailRegistered    string    `db:"email_registered"`
	CategoryID         uuid.UUID `db:"category_id"`
<<<<<<< HEAD
=======
	Category           AppealCategory
>>>>>>> 61b5f755d6a7c08963173028677fed4146ac3e1a
	Status             Status    `db:"status"`
	ProblemDescription string    `db:"problem_description"`
	Name               string    `db:"name"`
	EmailForConnect    string    `db:"email_for_connect"`
	ScreenshotURL      string    `db:"screenshot_url"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
