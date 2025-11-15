package models

import (
	"time"

	"github.com/google/uuid"
)

type Appeal struct {
	CreatorID          uuid.UUID `db:"creator_id"`
	EmailRegistered    string    `db:"email_registered"`
	Status             string    `db:"status" default:"created"`
	ProblemDescription string    `db:"problem_description"`
	Name               string    `db:"name"`
	EmailForConnection string    `db:"email_for_connection"`
	ScreenshotUrl      string    `db:"screenshot_url"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
