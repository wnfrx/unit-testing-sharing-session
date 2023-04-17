package models

import "time"

type Post struct {
	ID       int64     `json:"id" db:"id"`
	Title    string    `json:"title" db:"title"`
	Content  string    `json:"content" db:"content"`
	Author   string    `json:"author" db:"author"`
	PostedAt time.Time `json:"-" db:"posted_at"`
}
