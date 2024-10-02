package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"like"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (post *Post) PrepareValidation() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("o título é obrigatório e não pode estar em branco")
	}

	if post.Content == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
