package dto

// сюда кладем query параметры get-запроса
type FeedInputDTO struct {
	Offset int `schema:"offset"`
}

type FeedOutputDTO struct {
	Articles []ArticleOutput `json:"articles"`
}
