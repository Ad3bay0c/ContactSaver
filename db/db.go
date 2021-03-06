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
	CreateContact(contact *models.Contact) (interface{}, error)
	GetAllContacts(userID string) ([]models.Contact, error)
	GetAContact(userID, contactID string) (models.Contact, error)
	DeleteContact(userID, contactID string) error
	UpdateContact(contact *models.Contact, contactID string) error
}
