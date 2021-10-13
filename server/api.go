package server

import (
	"this/server/routers"
	_ "this/swagger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title gin-systray-starter
// @version 1.0.0
// @description "Starter code for Go application with Gin, System Tray, Gorm, Air, Swagger"

// @contact.name Github Issues
// @contact.url https://github.com/jinyaoMa/gin-systray-starter/issues

// @license.name MIT
// @license.url https://github.com/jinyaoMa/gin-systray-starter/blob/main/LICENSE

// @securityDefinitions.apikey BearerIdAuth
// @in header
// @name Authorization

func Apis(engine *gin.Engine, withSwagger bool) {
	routers.TestRouter(engine)

	if withSwagger {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
