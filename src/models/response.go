package models

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
type EventList struct {
	Events []Event `json:"events"`
}
