package handlers

import (
	"log"
	"net/http"

	"github.com/Kodnavis/face2face-backend/user-service/internal/models"
	"github.com/Kodnavis/face2face-backend/user-service/internal/repositories"
	"github.com/gin-gonic/gin"
)

type User struct {
	Repo *repositories.UserRepository
}

type UserRequest struct {
	Firstname string `json:"firstname" binding:"required,min=2,max=50"`
	Lastname  string `json:"lastname" binding:"required,min=2,max=50"`
	Login     string `json:"login" binding:"required,min=2,max=50"`
	Password  string `json:"password" binding:"required,min=8,max=72"`
}

func (u *User) Create(c *gin.Context) {
	var request UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &models.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Login:     request.Login,
		Password:  request.Password,
	}

	if err := u.Repo.Insert(user); err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *User) List(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "List all users",
	})
}

func (u *User) GetByID(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "Get a user by ID",
	})
}

func (u *User) UpdateById(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a user by ID",
	})
}

func (u *User) DeleteById(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a user by ID",
	})
}
