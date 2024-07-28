package models

// swagger:model Message
type Message struct {
	Body string `json:"body" example:"{\"message\": \"template message body\"}"`
}
