package db

import (
	"github.com/Ad3bay0c/ContactSaver/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	CreateUser(user *models.User) (interface{}, error)
	GetUserByEmail(email string) (*models.User, bool)
	InitializeCollection(collection string) *mongo.Collection
	GetAuthUser(id string) (*models.User, error)
	CreateContact(contact *models.Contact) error

}
