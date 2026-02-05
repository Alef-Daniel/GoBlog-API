package repositories

import "errors"

var (
	ErrAlreadyExists       = errors.New("post already exists")
	ErrNotFound            = errors.New("post not found")
	ErrRepository          = errors.New("error creating post")
	ErrRetrieve            = errors.New("error retrieving posts")
	ErrUpdate              = errors.New("error updating post")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrDelete              = errors.New("error deleting post")
)
