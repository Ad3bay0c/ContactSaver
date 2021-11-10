package db

import (
	"context"
	"github.com/Ad3bay0c/ContactSaver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (mongodb *MongoDB) CreateContact(contact *models.Contact) (interface{}, error) {
	collection := mongodb.InitializeCollection("contactss")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelFunc()
	result, err := collection.InsertOne(ctx, contact)

	return result.InsertedID, err
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


func (mongodb *MongoDB) DeleteContact(userID, contactID string) error {
	var contact models.Contact
	collection := mongodb.InitializeCollection("contactss")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelFunc()
	ID, _ := primitive.ObjectIDFromHex(contactID)
	err := collection.FindOneAndDelete(ctx, bson.M{"user_id": userID, "_id": ID}).Decode(&contact)
	return err
}

func (mongodb *MongoDB) UpdateContact(contact *models.Contact, contactID string) error {
	collection := mongodb.InitializeCollection("contactss")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelFunc()
	ID, _ := primitive.ObjectIDFromHex(contactID)
	contact.ID = ID
	err := collection.
		FindOneAndUpdate(ctx, bson.M{"_id": ID}, bson.D{{"$set", *contact}}).
		Decode(contact)
	return err
}