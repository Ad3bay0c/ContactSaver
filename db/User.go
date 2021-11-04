package db

import (
	"context"
	"github.com/Ad3bay0c/ContactSaver/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (mongodb *MongoDB) CreateUser(user *models.User) (interface{}, error) {
	collection := mongodb.InitializeCollection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	result, err := collection.InsertOne(ctx, user)
	return result.InsertedID, err
}

func (mongodb *MongoDB) GetUserByEmail(email string) (*models.User, bool) {
	user := &models.User{}
	collection := mongodb.InitializeCollection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	return user, err == nil
}
