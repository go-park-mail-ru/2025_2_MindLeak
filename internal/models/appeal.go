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

type Appeal struct {
	CreatorID          uuid.UUID `db:"creator_id"`
	EmailRegistered    string    `db:"email_registered"`
	Status             Status    `db:"status" default:"created"`
	ProblemDescription string    `db:"problem_description"`
	Name               string    `db:"name"`
	EmailForConnection string    `db:"email_for_connection"`
	ScreenshotUrl      string    `db:"screenshot_url"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
