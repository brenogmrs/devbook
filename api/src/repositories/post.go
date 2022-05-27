package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {

	statement, err := repository.db.Prepare(
		"insert into posts (title, content, author_id) values (?,?,?)",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthourID)

	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository Posts) GetById(postID uint64) (models.Post, error) {

	row, err := repository.db.Query(`
		select p.*, u.nick 
		from posts p inner join users u 
		on u.id = p.author_id
		where p.id = ?
	`, postID)

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
			&post.AuthourID,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.AuthourNick,
		); err != nil {
			return models.Post{}, err
		}

	}

	return post, nil
}

func (repository Posts) GetAllPosts(userID uint64) ([]models.Post, error) {

	rows, err := repository.db.Query(`
		select distinct p.*, u.nick from posts p 
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc
	`, userID, userID)

	if err != nil {
		return []models.Post{}, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {

		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthourID,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.AuthourNick,
		); err != nil {
			return []models.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare(
		"update posts set title = ?, content = ? where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}
func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from posts where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
