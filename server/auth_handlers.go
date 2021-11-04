package server

import (
	"github.com/Ad3bay0c/ContactSaver/models"
	"github.com/Ad3bay0c/ContactSaver/server/responses"
	"github.com/Ad3bay0c/ContactSaver/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SignUp Sign up a user
// Access Public
// @route POST
func (s *Server) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &models.User{}
		errs:= services.Decode(c, user)
		if errs != nil {
			responses.JSON(c, http.StatusBadRequest, "", errs, nil)
			return
		}
		_, ok := s.DB.GetUserByEmail(user.Email)
		if ok {
			responses.JSON(c, http.StatusBadRequest, "", "Email Already Taken", nil)
			return
		}
		hashPassword, err := services.HashPassword(user.Password)
		if err != nil {
			s.ErrorLog.Println(err.Error())
			responses.JSON(c, http.StatusBadRequest, "", "Internal Server Error", nil)
			return
		}
		user.Password = string(hashPassword)
		user.CreatedAt = time.Now()
		id, err := s.DB.CreateUser(user)
		if err != nil {
			responses.JSON(c, http.StatusBadRequest, "", "Internal Server Error", nil)
			return
		}
		token, err := services.GenerateToken(id)
		if err != nil {
			s.ErrorLog.Println(err.Error())
			responses.JSON(c, http.StatusBadRequest, "", "Internal Server Error", nil)
			return
		}
		responses.JSON(c, http.StatusCreated, "User Created Successfully", nil, token)
	}
}

// Login Log in a user
// Access Public
// @route POST
func (s *Server) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &models.User{}

		errs:= services.Decode(c, user)
		if errs != nil {
			responses.JSON(c, http.StatusBadRequest, "", errs, nil)
			return
		}
		DbUser, ok := s.DB.GetUserByEmail(user.Email)
		if !ok {
			responses.JSON(c, http.StatusBadRequest, "", "Email Does not Exist", nil)
			return
		}
		err := services.ComparePassword(user.Password, DbUser.Password)
		if err != nil {
			responses.JSON(c, http.StatusBadRequest, "", "Incorrect Password", nil)
			return
		}
		token, err := services.GenerateToken(DbUser.ID)
		if err != nil {
			s.ErrorLog.Println(err.Error())
			responses.JSON(c, http.StatusBadRequest, "", "Internal Server Error", nil)
			return
		}
		responses.JSON(c, http.StatusCreated, "User Created Successfully",
			nil, gin.H{"token": token, "user": DbUser})
	}
}

// GetUser Get a logged in User
// Access Private
// @route GET
func (s *Server) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		responses.JSON(c, 200, "Fetched Successfully", "", nil)
	}

}
