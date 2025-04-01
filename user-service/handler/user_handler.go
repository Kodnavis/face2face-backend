package handler

import (
	"net/http"

	"github.com/Kodnavis/face2face-backend/user-service/repository"
	"github.com/gin-gonic/gin"
)

type User struct {
	Repo *repository.UserRepo
}

func (u *User) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create a user",
	})
}

func (u *User) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all users",
	})
}

func (u *User) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get a user by ID",
	})
}

func (u *User) UpdateById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a user by ID",
	})
}

func (u *User) DeleteById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a user by ID",
	})
}
