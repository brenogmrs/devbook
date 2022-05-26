package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	AuthourID   uint64    `json:"author_id,omitempty"`
	AuthourNick string    `json:"author_nick,omitempty"`
	Likes       uint64    `json:"likes"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (post *Post) Prepare() error {
	if err := post.validateFields(); err != nil {
		return err
	}

	post.formatFields()
	return nil
}

func (post *Post) validateFields() error {
	if post.Title == "" {
		return errors.New("field `title` is required")
	}

	if post.Content == "" {
		return errors.New("field `content` is required")
	}

	return nil
}

func (post *Post) formatFields() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
