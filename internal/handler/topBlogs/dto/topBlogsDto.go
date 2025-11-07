package dto

import "github.com/google/uuid"

type TopBlogDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Subscribers int       `json:"subscribers"`
}

type TopBlogsDto struct {
	Blogs []TopBlogDto
}
