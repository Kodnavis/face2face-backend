package middlewares

import (
	"errors"
	"net/http"

	"github.com/Kodnavis/face2face-backend/user-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token_string, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user_id, err := utils.ExtractToken(token_string)
		if err != nil {
			if errors.Is(err, utils.ErrInvalidToken) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			c.Status(http.StatusInternalServerError)
		}

		c.Set("user_login", user_id)
		c.Next()
	}
}
