package server

import (
	"github.com/Ad3bay0c/ContactSaver/models"
	"github.com/Ad3bay0c/ContactSaver/server/responses"
	"github.com/Ad3bay0c/ContactSaver/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func (s *Server) CreateContact() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			responses.JSON(c, http.StatusInternalServerError, "", "Internal Server Error", nil)
			return
		}
		ID := userID.(string)
		contact := &models.Contact{}
		errs:= services.Decode(c, contact)
		if errs != nil {
			responses.JSON(c, http.StatusBadRequest, "", errs, nil)
			return
		}
		contact.UserID = ID
		contact.Date = time.Now()
		id, err := s.DB.CreateContact(contact)
		if err != nil {
			responses.JSON(c, http.StatusBadRequest, "", "Error Creating Contact", nil)
			return
		}
		contactID := id.(primitive.ObjectID)
		contact.ID = contactID
		responses.JSON(c, http.StatusCreated, "Contact Created Successfully", "", contact)
		return
	}
}

func (s *Server) GetAllContacts() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			responses.JSON(c, http.StatusInternalServerError, "", "Internal Server Error", nil)
			return
		}
		ID := userID.(string)
		contacts, err := s.DB.GetAllContacts(ID)
		if err != nil {
			s.ErrorLog.Println(err.Error())
			responses.JSON(c, http.StatusInternalServerError, "", "Error Fetching Contacts", nil)
			return
		}
		responses.JSON(c, http.StatusCreated, "Fetched Successfully", "", contacts)
		return
	}
}

func (s *Server) GetAContact() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			responses.JSON(c, http.StatusInternalServerError, "", "Internal Server Error", nil)
			return
		}
		ID := userID.(string)
		contactID := c.Param("contactID")
		contact, err := s.DB.GetAContact(ID, contactID)
		if err != nil {
			s.ErrorLog.Println(err.Error())
			responses.JSON(c, http.StatusInternalServerError, "", "Invalid Contact ID", nil)
			return
		}
		responses.JSON(c, http.StatusCreated, "Fetched Successfully", "", contact)
		return
	}
}

func (s *Server) UpdateContact() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			responses.JSON(c, http.StatusInternalServerError, "", "Internal Server Error", nil)
			return
		}
		ID := userID.(string)
		contact := &models.Contact{}
		errs:= services.Decode(c, contact)
		if errs != nil {
			responses.JSON(c, http.StatusBadRequest, "", errs, nil)
			return
		}
		contact.Date = time.Now()
		contact.UserID = ID
		contactID := c.Param("contactID")

		err := s.DB.UpdateContact(contact, contactID)
		if err != nil {
			responses.JSON(c, http.StatusBadRequest, "", "Error Updating Contact", nil)
			return
		}
		responses.JSON(c, http.StatusCreated, "Contact Updated Successfully", "", nil)
		return
	}
}

func (s *Server) DeleteContact() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.Get("userID")
		if !ok {
			responses.JSON(c, http.StatusInternalServerError, "", "Internal Server Error", nil)
			return
		}
		ID := userID.(string)
		contactID := c.Param("contactID")
		err := s.DB.DeleteContact(ID, contactID)
		if err != nil {
			s.ErrorLog.Println(err.Error())
			responses.JSON(c, http.StatusInternalServerError, "", "Invalid Contact ID", nil)
			return
		}
		responses.JSON(c, http.StatusCreated, "Deleted Successfully", "", nil)
		return
	}
}