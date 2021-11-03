package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Working Perfectly"})
	}
}
