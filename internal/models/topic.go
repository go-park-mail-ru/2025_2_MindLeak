package models

type Topic struct {
	TopicId int    `db:"topic_id"`
	Title   string `db:"title"`
}
