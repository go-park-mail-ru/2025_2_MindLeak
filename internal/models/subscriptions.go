package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	FollowerID uuid.UUID `db:"follower_id"`
	FoolowedID uuid.UUID `db:"followed_id"`
	CreatedAt  time.Time `db:"created_at"`
}
