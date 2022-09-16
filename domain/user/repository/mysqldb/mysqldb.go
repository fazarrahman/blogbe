package mysqldb

import (
	"context"

	ue "github.com/fazarrahman/blogbe/domain/user/entity"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
)

type Mysqldb struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Mysqldb {
	return &Mysqldb{db: db}
}

func (m *Mysqldb) InsertUser(ctx context.Context, user *ue.User) error {
	return nil
}

func (m *Mysqldb) GetUsers(ctx context.Context, filter bson.M) ([]*ue.User, error) {
	return nil, nil
}

func (m *Mysqldb) GetUser(ctx context.Context, fieldName string, value interface{}) (*ue.User, error) {
	return nil, nil
}
