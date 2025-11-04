package models

import "github.com/google/uuid"

type Topic struct {
	TopicId uuid.UUID `db:"topic_id"`
	Title   string    `db:"title"`
}
