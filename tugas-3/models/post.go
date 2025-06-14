package models

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
