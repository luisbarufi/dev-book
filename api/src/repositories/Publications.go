package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func NewPublicationsRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into publications (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.Author_id)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (repository Publications) FindById(publicationId uint64) (models.Publication, error) {
	row, err := repository.db.Query(
		`select p.*, u.nick from publications p inner join users u on u.id = p.author_id where p.id = ?`,
		publicationId,
	)
	if err != nil {
		return models.Publication{}, err
	}
	defer row.Close()

	var publication models.Publication

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.Author_id,
			&publication.Likes,
			&publication.Created_at,
			&publication.Author_nick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}
