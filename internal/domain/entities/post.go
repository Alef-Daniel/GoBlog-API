package entities

import "time"

type Post struct {
	ID         int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Title      string
	CategoryID int64
	Author     string
	Content    string
}
