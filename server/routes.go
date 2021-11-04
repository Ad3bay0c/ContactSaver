package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Server) Routes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	AuthRouter := router.Group("/api/auth")
	UserRouter := router.Group("/api/user")
	//ContactRouter := UserRouter.Group("/contact")
	AuthRouter.POST("/", s.SignUp())
	AuthRouter.POST("/login", s.Login())

	UserRouter.GET("/", s.GetUser())
}
