package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

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

type UserResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Login     string `json:"login"`
	CreatedAt string `json:"created_at"`
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
			"error": "failed to create user",
		})
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Login:     user.Login,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}

func (u *User) List(c *gin.Context) {
	var query_params repositories.FindAllQueryParams

	if err := c.ShouldBindQuery(&query_params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid query parameters",
		})
		return
	}

	if query_params.Size <= 0 {
		query_params.Size = 10
	}
	if query_params.Offset < 0 {
		query_params.Offset = 0
	}

	users, err := u.Repo.FindAll(repositories.FindAllQueryParams{
		Size:   query_params.Size,
		Offset: query_params.Offset,
	})

	if err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list users",
		})
		return
	}

	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"meta": gin.H{
			"size":   query_params.Size,
			"offset": query_params.Offset,
		},
	})
}

func (u *User) Get(c *gin.Context) {
	login := c.Param("login")
	user, err := u.Repo.FindOne(login)

	if err != nil {
		if errors.Is(err, repositories.ErrNotExist) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Login:     user.Login,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}

func (u *User) Update(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a user by ID",
	})
}

func (u *User) Delete(c *gin.Context) {
	login := c.Param("login")
	err := u.Repo.Delete(login)

	if err != nil {
		if errors.Is(err, repositories.ErrNotExist) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.Status(http.StatusOK)
}
