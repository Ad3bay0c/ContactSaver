package db

import (
	"context"
	"github.com/Ad3bay0c/ContactSaver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (mongodb *MongoDB) CreateContact(contact *models.Contact) error {
	collection := mongodb.InitializeCollection("contactss")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelFunc()
	_, err := collection.InsertOne(ctx, contact)
	return err
}
func (mongodb *MongoDB) GetAllContacts(userID string) ([]models.Contact, error) {
	var contacts []models.Contact
	collection := mongodb.InitializeCollection("contactss")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelFunc()
	Opt := options.Find().SetSort(bson.D{{"date", -1}})
	res, err := collection.Find(ctx, bson.M{"user_id": userID}, Opt)
	if err != nil {
		return nil, err
	}
	err = res.All(ctx, &contacts)
	return contacts, err
}

func (mongodb *MongoDB) GetAContact(userID, contactID string) (models.Contact, error) {
	var contact models.Contact
	collection := mongodb.InitializeCollection("contactss")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelFunc()
	ID, _ := primitive.ObjectIDFromHex(contactID)
	err := collection.FindOne(ctx, bson.M{"user_id": userID, "_id": ID}).Decode(&contact)
	return contact, err
}


