package post

import "time"

type Post struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	SubTitle    *string
	Published   bool
	Description string
}
