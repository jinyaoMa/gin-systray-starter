package server

import (
	"net/http"
	"this/database"
	"this/server/models"

	"github.com/gin-gonic/gin"
)

func New() *Server {
	r := gin.Default()

	init_router(r)

	database.DB.AutoMigrate(
		&models.TestModel{},
	)

	return &Server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
	}
}
