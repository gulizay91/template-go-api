package models

// swagger:model Template
type Template struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
