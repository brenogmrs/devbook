package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {

	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?,?,?,?)",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository users) GetAllUsers(query string) ([]models.User, error) {

	query = fmt.Sprintf("%%%s%%", query)

	rows, err := repository.db.Query(
		"select id, name, email, nick, created_at, updated_at FROM users WHERE name LIKE ? or nick LIKE ?",
		query, query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetUserById(id uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"select id, name, email, nick, created_at, updated_at FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return models.User{}, err
		}

	}

	return user, err

}

func (repository users) Update(userID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

func (repositories users) Delete(userID uint64) error {
	statement, err := repositories.db.Prepare("delete from users where id = ?")

	if err != nil {
		return nil
	}

	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}

func (repositories users) GetByEmail(email string) (models.User, error) {

	row, err := repositories.db.Query("select id, password from users where email = ?", email)

	if err != nil {
		return models.User{}, err
	}

	defer row.Close()

	var user models.User

	if row.Next() {

		if err = row.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}

	}

	return user, err
}

func (repositories users) Follow(userID, followerID uint64) error {

	statement, err := repositories.db.Prepare("insert ignore into followers(user_id, follower_id) values (?, ?)")

	if err != nil {
		return nil
	}

	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repositories users) Unfollow(userID, followerID uint64) error {

	statement, err := repositories.db.Prepare("delete from followers where user_id = ? and follower_id = ?")

	if err != nil {
		return nil
	}

	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repositories users) GetUserFollowers(userID uint64) ([]models.User, error) {

	row, err := repositories.db.Query(`
		select u.id, u.name, u.nick, u.email, u.created_at 
		from users u
		inner join followers f on (u.id = f.follower_id) 
		where s.user_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var users []models.User

	for row.Next() {
		var user models.User

		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repositories users) GetUserFollows(userID uint64) ([]models.User, error) {
	row, err := repositories.db.Query(`
		select u.id, u.name, u.nick, u.email, u.created_at 
		from users u
		inner join followers f on (u.id = f.user_id) 
		where s.follower_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var users []models.User

	for row.Next() {
		var user models.User

		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repositories users) GetPassword(userID uint64) (string, error) {
	row, err := repositories.db.Query(`
		select u.password from users u where u.id = ?
	`, userID)

	if err != nil {
		return "", err
	}

	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository users) ChangePassword(userID uint64, newPassword string) error {
	statement, err := repository.db.Prepare(
		"update users set password = ? where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(newPassword, userID); err != nil {
		return err
	}

	return nil
}
