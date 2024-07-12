package storage

import (
	"context"
	"service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func NewUserRepo(c *mongo.Collection) *UserRepo {
	return &UserRepo{Collection: c}
}

func (db *UserRepo) SignUp(req *models.User) error {

	_, err := db.Collection.InsertOne(context.Background(), bson.M{
		"full_name": req.Fullname,
		"email":     req.Email,
		"password":  req.Password,
	})

	if err != nil {
		return err
	}

	return nil
}
