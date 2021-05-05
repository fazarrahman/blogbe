package mongodb

import (
	ue "blogbe/domain/user/entity"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongodb struct {
	usersCollection *mongo.Collection
}

func New(_mdb *mongo.Database) *Mongodb {
	return &Mongodb{usersCollection: _mdb.Collection("users")}
}

func (m *Mongodb) InsertUser(ctx context.Context, user *ue.User) error {
	_, err := m.usersCollection.InsertOne(ctx, user)
	return err
}

func (m *Mongodb) GetUsers(ctx context.Context, filter bson.M) ([]*ue.User, error) {
	var users []*ue.User
	res, err := m.usersCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for res.Next(ctx) {
		var user ue.User
		err = res.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (m *Mongodb) GetUser(ctx context.Context, fieldName string, value interface{}) (*ue.User, error) {
	var res *mongo.SingleResult
	if fieldName == "_id" {
		objectId, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", value))
		if err != nil {
			log.Println("Invalid id")
		}
		res = m.usersCollection.FindOne(ctx, bson.M{fieldName: objectId})
	} else {
		res = m.usersCollection.FindOne(ctx, bson.M{fieldName: value})
	}
	var user ue.User
	err := res.Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	if user.Username == "" {
		return nil, nil
	}
	return &user, nil
}
