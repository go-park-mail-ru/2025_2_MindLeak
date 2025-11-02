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
	Id          uuid.UUID
	UserID      uuid.UUID
	Phone       string
	Country     string
	Language    string
	Sex         Sex
	DateOfBirth time.Time
	Age         int
	CoverURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
