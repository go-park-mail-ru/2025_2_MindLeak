package dto

//сюда кладем query параметры get-запроса
type FeedInputDTO struct {
	Offset int `json:"offset"`
}

type FeedOutputDTO struct {
	Articles []ArticleOutputDTO `json:"articles"`
}
