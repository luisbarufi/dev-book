package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) PrepareValidation(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("o apelido é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if step == "register" && user.Password == "" {
		return errors.New("a senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
