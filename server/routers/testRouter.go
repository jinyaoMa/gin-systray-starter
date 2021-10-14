package routers

import (
	"this/server/controllers"
	"this/server/middlewares"

	"github.com/gin-gonic/gin"
)

func TestRouter(engine *gin.Engine) {
	g := engine.Group("/test")
	{
		g.Use(middlewares.TestMiddleware())

		g.GET("/", controllers.TestController)
		g.GET("/getToken", controllers.TestJwtGetToken)
		g.GET("/checkToken", middlewares.Auth(), controllers.TestJwtCheckToken)
	}
}
