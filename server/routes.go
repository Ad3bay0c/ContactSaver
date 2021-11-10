package server

import (
	"github.com/Ad3bay0c/ContactSaver/server/middleware"
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
	//ContactRouter := UserRouter.Group("/contact")
	AuthRouter.POST("/", s.SignUp())
	AuthRouter.POST("/login", s.Login())

	UserRouter := router.Group("/api/user")
	UserRouter.Use(middleware.Authorize())
	UserRouter.GET("/", s.GetUser())
	ContactRouter := UserRouter.Group("/contact")
	{
		ContactRouter.POST("/", s.CreateContact())
		ContactRouter.GET("/", s.GetAllContacts())
		ContactRouter.GET("/:contactID", s.GetAContact())
		ContactRouter.DELETE("/:contactID", s.DeleteContact())
		ContactRouter.PUT("/:contactID", s.UpdateContact())
	}
}
