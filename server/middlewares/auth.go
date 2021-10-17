package middlewares

import (
	"net/http"
	"regexp"
	"this/gate"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	bearerRegexp, err := regexp.Compile(`^Bearer (.+)$`)
	if err != nil {
		panic("Failed to compile regexp in middleware Auth")
	}

	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		matches := bearerRegexp.FindStringSubmatch(authorization)
		if matches == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := matches[1]
		claims, err := gate.ParseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
