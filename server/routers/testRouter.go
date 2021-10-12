package routers

import (
	"this/server/controllers"
	"this/server/middlewares"

	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	r.Use(middlewares.TestMiddleware())

	r.GET("/test", controllers.TestController)
}
