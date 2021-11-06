package db

import (
	"context"
	"github.com/Ad3bay0c/ContactSaver/models"
	"time"
)

func (mongodb *MongoDB) CreateContact(contact *models.Contact) error {
	collection := mongodb.InitializeCollection("contactss")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelFunc()
	_, err := collection.InsertOne(ctx, contact)
	return err
}