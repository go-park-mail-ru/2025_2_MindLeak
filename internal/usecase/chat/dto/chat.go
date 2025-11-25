package dto

import (
	"encoding/json"
	"time"
)

type ServerEvent struct {
	Type string    `json:"type"`
	Data any       `json:"data"`
	Ts   time.Time `json:"ts"`
}

type ClientMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
