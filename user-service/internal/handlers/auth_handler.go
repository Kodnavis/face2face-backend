package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Kodnavis/face2face-backend/common/auth"

	"github.com/Kodnavis/face2face-backend/user-service/internal/data/requests"
	"github.com/Kodnavis/face2face-backend/user-service/internal/repositories"
	"github.com/gin-gonic/gin"
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

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))
	if err != nil {
		log.Panicln(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	token, err := auth.GenerateToken(user.Login, token_lifespan)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Authorization",
		token,
		3600*token_lifespan,
		"/",
		"",
		false,
		true,
	)

	type publicUser struct {
		ID        uint   `json:"id"`
		Login     string `json:"login"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}
	c.JSON(http.StatusOK, publicUser{
		ID:        user.ID,
		Login:     user.Login,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	})
}
