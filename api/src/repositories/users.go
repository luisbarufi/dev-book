package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository *users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}
