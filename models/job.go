package models

type Job struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
}
