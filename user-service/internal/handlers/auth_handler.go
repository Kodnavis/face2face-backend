package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Kodnavis/face2face-backend/user-service/internal/data/requests"
	"github.com/Kodnavis/face2face-backend/user-service/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) Login(c *gin.Context) {
	var login_request requests.LoginRequest

	if err := c.ShouldBindJSON(&login_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := u.Repo.FindOne(login_request.Login)
	if err != nil {
		if errors.Is(err, repositories.ErrNotExist) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "invalid login or password",
			})
			return
		}

		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login_request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid login or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	token_string, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token_string, 3600, "", "", false, true)

	c.Status(http.StatusOK)
}
