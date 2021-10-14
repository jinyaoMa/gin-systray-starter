package controllers

import (
	"net/http"
	"this/server/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// @Summary TestJwtGetToken
// @Description Test Jwt, get token
// @Tags Before Authorization
// @accept plain
// @Produce json
// @Success 200 "{ ok , token }"
// @Failure 404 "{ error }"
// @Router /test/getToken/ [get]
func TestJwtGetToken(c *gin.Context) {
	now := time.Now()
	claims := middlewares.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       "1",
			Issuer:   "gin-systray-starter",
			Subject:  "user",
			Audience: []string{"gin-systray-starter", "user"},
			ExpiresAt: &jwt.NumericDate{
				Time: now.AddDate(0, 0, 1), // 1 day
			},
			NotBefore: &jwt.NumericDate{
				Time: now,
			},
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
		},
	}

	jwToken := middlewares.JWT{
		SigningKey: []byte(middlewares.DEFAULT_SIGNING_KEY),
	}
	token, err := jwToken.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "TestJwtGetToken() error!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"token": token,
	})
}
