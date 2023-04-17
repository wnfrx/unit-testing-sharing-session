package models

import "errors"

var (
	ErrInvalidId = errors.New("invalid id")
	ErrNotFound  = errors.New("not found")

	// Post
	ErrPostTitleEmpty = errors.New("post title is empty")
)
