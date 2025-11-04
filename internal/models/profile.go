package models

import (
	"github.com/google/uuid"
	"time"
)

type Sex string

const (
	SexMale      Sex = "male"
	SexFemale    Sex = "female"
	SexUndefined Sex = "undefined"
)

type Profile struct {
	Id          uuid.UUID `db:"profile_id"`
	UserID      uuid.UUID `db:"user_id"`
	Phone       string    `db:"phone"`
	Country     string    `db:"country"`
	Language    string    `db:"language"`
	Sex         Sex       `db:"sex"`
	DateOfBirth time.Time `db:"date_of_birth"`
	Age         int       `db:"age"`
	description string    `db:"description"`
	CoverURL    string    `db:"cover_url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
