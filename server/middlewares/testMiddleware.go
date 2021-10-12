package middlewares

import (
	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
