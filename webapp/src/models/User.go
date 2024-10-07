package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedAt time.Time `json:"created_at"`
	Followers []User    `json:"followers"`
	Follower  []User    `json:"follower"`
	Posts     []Post    `json:"posts"`
}
