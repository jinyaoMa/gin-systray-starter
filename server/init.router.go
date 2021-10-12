package server

import (
	"net/http"
	"this/server/controllers"
	"this/server/middlewares"

	"github.com/gin-gonic/gin"
)

func init_router(r *gin.Engine) {
	r.Use(middlewares.TestMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": "/",
		})
	})

	api := r.Group("api")
	{
		api.GET("/test", controllers.TestController)
	}
}
