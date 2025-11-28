package dto

import "github.com/google/uuid"

type SubscriberDto struct {
	Id          uuid.UUID `json:"id"`
	Avatar      string    `json:"avatar"`
	Name        string    `json:"name"`
	Subscribers int       `json:"subscriptions"`
}

type SubscriptionDto struct {
	Id          uuid.UUID `json:"id"`
	Avatar      string    `json:"avatar"`
	Name        string    `json:"name"`
	Subscribers int       `json:"subscriptions"`
}

type SubscriptionsOutputDto struct {
	Subsccriptions []SubscriptionDto `json:"subscriptions"`
}

type SubscribersOutputDto struct {
	Subscribers []SubscriberDto `json:"subscriptions"`
}
