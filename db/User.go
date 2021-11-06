package db

import (
	"context"
	"github.com/Ad3bay0c/ContactSaver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (mongodb *MongoDB) GetAuthUser(id string) (*models.User, error){
	user := &models.User{}
	collection := mongodb.InitializeCollection("user")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	newId, _ := primitive.ObjectIDFromHex(id) //create a new ObjectID from the string
	err := collection.FindOne(ctx, bson.M{"_id": newId}).Decode(user)
	cancelFunc()
	return user, err
}
