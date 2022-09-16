package mysqldb

import (
	"context"

	"github.com/fazarrahman/blogbe/domain/gallery/entity"
	"github.com/jmoiron/sqlx"
)

type Mysqldb struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Mysqldb {
	return &Mysqldb{db: db}
}

func (m *Mysqldb) GetGallery(ctx context.Context) ([]*entity.Gallery, error) {
	var galleries []*entity.Gallery
	err := m.db.SelectContext(ctx,
		&galleries,
		`select id, source from gallery where is_active = 1`)
	if err != nil {
		return nil, err
	}
	return galleries, err
}
