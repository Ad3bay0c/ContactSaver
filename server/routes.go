package server

import "github.com/gin-gonic/gin"

func (s *Server) Routes(router *gin.Engine) {
	router.GET("/", s.SignUp())
}