package models

import (
	"time"
)

type Post struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	Author_id   uint64    `json:"author_id,omitempty"`
	Author_nick string    `json:"author_nick,omitempty"`
	Likes       uint64    `json:"like"`
	Created_at  time.Time `json:"created_at,omitempty"`
}
