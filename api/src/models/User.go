package models

import (
	"errors"
	"strings"
	"time"
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

func (user *User) Prepare() error {
	if err := user.validateFields(); err != nil {
		return err
	}

	user.formatData()
	return nil
}

func (user *User) validateFields() error {
	if user.Name == "" {
		return errors.New("field `name` is required")
	}
	if user.Email == "" {
		return errors.New("field `email` is required")
	}
	if user.Nick == "" {
		return errors.New("field `nick` is required")
	}
	if user.Password == "" {
		return errors.New("field `password` is required")
	}

	return nil
}

func (user *User) formatData() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
