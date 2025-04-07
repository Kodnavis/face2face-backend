package jwt

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(jwt_secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token_string, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user_id, err := ExtractToken(token_string, jwt_secret)
		if err != nil {
			if errors.Is(err, ErrInvalidToken) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			c.Status(http.StatusInternalServerError)
		}

		c.Set("user_login", user_id)
		c.Next()
	}
}
