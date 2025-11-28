package dto

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

type UserDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Subscribers int       `json:"subscribers"`
}

type ArticleDto struct {
	ID            uuid.UUID    `json:"id"`
	AuthorID      uuid.UUID    `json:"author_id"`
	Title         string       `json:"title"`
	Content       string       `json:"content"`
	MediaURL      string       `json:"media_url,omitempty"`
	TopicID       int          `json:"topic_id"`
	Status        string       `json:"status"`
	CommentsCount int          `json:"comments_count"`
	RepostsCount  int          `json:"reposts_count"`
	ViewsCount    int          `json:"views_count"`
	Topic         models.Topic `db:"-"`
	AuthorName    string       `json:"author_name"`
	AuthorAvatar  string       `json:"author_avatar"`
}

type UsersDto struct {
	Users []UserDto `json:"users"`
}

type ArticlesDto struct {
	Articles []ArticleDto `json:"articles"`
}

//type SearchAllOutputDto struct {
//	Users    []UserDto    `json:"users"`
//	Articles []ArticleDto `json:"articles"`
//}
