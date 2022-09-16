package repository

import (
	"context"

	"github.com/fazarrahman/blogbe/domain/gallery/entity"
)

type Repository interface {
	GetGallery(ctx context.Context) ([]*entity.Gallery, error)
}
