package repository

import (
	"context"

	ue "github.com/fazarrahman/blogbe/domain/user/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Repository ..
type Repository interface {
	InsertUser(ctx context.Context, user *ue.User) error
	GetUsers(ctx context.Context, filter bson.M) ([]*ue.User, error)
	GetUser(ctx context.Context, fieldName string, value interface{}) (*ue.User, error)
}
