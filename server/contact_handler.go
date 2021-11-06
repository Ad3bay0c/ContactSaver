package server

import (
	"github.com/Ad3bay0c/ContactSaver/models"
	"github.com/Ad3bay0c/ContactSaver/server/responses"
	"github.com/Ad3bay0c/ContactSaver/services"
	"github.com/gin-gonic/gin"
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
		err := s.DB.CreateContact(contact)
		if err != nil {
			responses.JSON(c, http.StatusBadRequest, "", "Error Creating Contact", nil)
			return
		}
		responses.JSON(c, http.StatusCreated, "Contact Created Successfully", "", nil)
		return
	}
}
