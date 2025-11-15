package dto

type AppealDTO struct {
	AppealID           string `json:"appealId"`
	Status             string `json:"status"`
	CategoryName       string `json:"category"`
	ProblemDescription string `json:"problemDescription"`
	CreatedAt          int64  `json:"createdAt"`
}

type AppealsStatisticsDTO struct {
	Total      int64            `json:"total"`
	ByCategory map[string]int64 `json:"byCategory"`
	ByStatus   map[string]int64 `json:"byStatus"`
	Appeals    []AppealDTO      `json:"appeals"`
}
