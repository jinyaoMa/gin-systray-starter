package middlewares

import (
	"errors"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const DEFAULT_SIGNING_KEY = "secret key"

type JWT struct {
	SigningKey []byte
}
type JWTClaims struct {
	jwt.RegisteredClaims
}

var (
	ErrTokenExpired     error = errors.New("token is expired")
	ErrTokenNotValidYet error = errors.New("token not active yet")
	ErrTokenMalformed   error = errors.New("that's not even a token")
	ErrTokenInvalid     error = errors.New("couldn't handle this token")
)

func (j *JWT) CreateToken(claims JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

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
		jwToken := JWT{
			SigningKey: []byte(DEFAULT_SIGNING_KEY),
		}
		claims, err := jwToken.ParseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
