package controllers

import (
	"net/http"
	"this/server/middlewares"

	"github.com/gin-gonic/gin"
)

// @Summary TestJwtCheckToken
// @Description Test Jwt, check token
// @Tags After Authorization
// @accept plain
// @Produce json
// @Security BearerIdAuth
// @param Authorization header string false "Authorization"
// @Success 200 "{ ok , claims }"
// @Failure 404 "{ error }"
// @Router /test/checkToken/ [get]
func TestJwtCheckToken(c *gin.Context) {
	temp, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Middleware Auth not exists!",
		})
		return
	}
	claims := temp.(*middlewares.JWTClaims)

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"claims": claims,
	})
}
