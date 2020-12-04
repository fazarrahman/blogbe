package entity

import "time"

// User ...
type User struct {
	ID          string     `bson:"_id,omitempty"`
	Email       string     `bson:"email"`
	Username    string     `bson:"username"`
	FirstName   string     `bson:"first_name"`
	LastName    string     `bson:"last_name"`
	Password    []byte     `bson:"password"`
	Status      int8       `bson:"status"`
	CreatedDate time.Time  `bson:"created_date"`
	UpdatedDate *time.Time `bson:"updated_date"`
	DeletedDate *time.Time `bson:"deleted_date"`
}
