package dto

type TopBlogsDto struct {
	Name             string `json:"name"`
	Avatar           string `json:"avatar"`
	SubscribersCount int    `json:"subscribers_count"`
}
