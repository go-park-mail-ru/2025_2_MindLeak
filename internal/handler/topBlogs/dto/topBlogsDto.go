package dto

type TopBlogDto struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Subscribers int    `json:"subscribers"`
}

type TopBlogsDto struct {
	Blogs []TopBlogDto
}
