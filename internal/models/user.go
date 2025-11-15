package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `db:"user_id"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Name     string    `db:"name"`
	Avatar   string    `db:"avatar"`

	Subscribers   int `db:"subscriptions"`
	Subscriptions int `db:"subscriptions"`
}
