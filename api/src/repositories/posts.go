package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into posts (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.Author_id)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (repository Posts) FindById(postId uint64) (models.Post, error) {
	row, err := repository.db.Query(
		`select p.*, u.nick from posts p inner join users u on u.id = p.author_id where p.id = ?`,
		postId,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Author_id,
			&post.Likes,
			&post.Created_at,
			&post.Author_nick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Posts) ListPosts(userId uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(
		`select distinct p.*, u.nick from posts p inner join users u on u.id = p.author_id inner join followers s on p.author_id = s.user_id 
		where u.id = ? or s.follower_id = ?	order by 1 desc`, userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Author_id,
			&post.Likes,
			&post.Created_at,
			&post.Author_nick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) UpdatePost(postId uint64, post models.Post) error {
	statement, err := repository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postId); err != nil {
		return err
	}

	return nil
}

func (repository Posts) DeletePost(postId uint64) error {
	statement, err := repository.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (repository Posts) ListPostsByUser(userId uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(
		`select p.*, u.nick from posts p join users u on u.id = p.author_id where p.author_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Author_id,
			&post.Likes,
			&post.Created_at,
			&post.Author_nick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) LikePost(postId uint64) error {
	statement, err := repository.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
