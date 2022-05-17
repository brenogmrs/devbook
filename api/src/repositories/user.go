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
