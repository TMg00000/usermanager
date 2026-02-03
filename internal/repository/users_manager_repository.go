package repository

import (
	"context"
	"time"
	"usermanager/internal/domain/request"
	"usermanager/internal/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserManagerRepository struct {
	collection *mongo.Collection
}

func NewUserManagerRepository(col *mongo.Collection) *UserManagerRepository {
	return &UserManagerRepository{
		collection: col,
	}
}

func (c *UserManagerRepository) Create(u request.Users) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	HashPassword, errHash := utils.GenerateHashPassword(u.Password)
	if errHash != nil {
		return errHash
	}
	u.Hash = string(HashPassword)

	docUsers := bson.M{
		"name":     u.Username,
		"email":    u.Email,
		"password": u.Hash,
	}

	_, err := c.collection.InsertOne(ctx, docUsers)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserManagerRepository) Get(email, password string) error {
	var u request.Users

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	if err := c.collection.FindOne(ctx, filter).Decode(&u); err != nil {
		return err
	}

	return utils.ComparePassword(u.Password, password)
}
