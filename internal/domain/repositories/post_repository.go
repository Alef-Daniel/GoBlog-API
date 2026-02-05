package repositories

import (
	"context"

	"github.com/goblog-api/internal/domain/entities"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *entities.Post) (*entities.Post, error)
	GetByID(ctx context.Context, id int64) (*entities.Post, error)
	GetAllPost(ctx context.Context) (*[]entities.Post, error)
	UpdatePost(ctx context.Context, post *entities.Post) (*entities.Post, error)
	DeletePost(ctx context.Context, post *entities.Post) error
}
