package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"this/database"

	"this/server/controllers"
	"this/server/middlewares"
	"this/server/models"
)

func New() *Server {
	router := gin.Default()

	router.Use(middlewares.TestMiddleware())

	api := router.Group("api")
	{
		api.GET("/test", controllers.TestController)
	}

	database.DB.AutoMigrate(
		&models.TestModel{},
	)

	return &Server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}
