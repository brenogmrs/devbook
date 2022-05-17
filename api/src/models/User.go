package models

import (
	"api/src/helpers"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User model
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *User) Prepare(method string) error {
	if err := user.validateFields(method); err != nil {
		return err
	}

	if err := user.formatData(method); err != nil {
		return err
	}

	return nil
}

func (user *User) validateFields(method string) error {
	if user.Name == "" {
		return errors.New("field `name` is required")
	}

	if user.Email == "" {
		return errors.New("field `email` is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email format")
	}

	if user.Nick == "" {
		return errors.New("field `nick` is required")
	}

	if method == "create" && user.Password == "" {
		return errors.New("field `password` is required")
	}

	return nil
}

func (user *User) formatData(method string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if method == "create" {
		hashedPassword, err := helpers.Hash(user.Password)

		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
