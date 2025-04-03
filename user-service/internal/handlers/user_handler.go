package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Kodnavis/face2face-backend/user-service/internal/data/requests"
	"github.com/Kodnavis/face2face-backend/user-service/internal/data/responses"
	"github.com/Kodnavis/face2face-backend/user-service/internal/models"
	"github.com/Kodnavis/face2face-backend/user-service/internal/repositories"
	"github.com/gin-gonic/gin"
)

type User struct {
	Repo *repositories.UserRepository
}

func (u *User) Create(c *gin.Context) {
	var request requests.CreateUserRequest
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

	response := responses.UserResponse{
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

	var response []responses.UserResponse
	for _, user := range users {
		response = append(response, responses.UserResponse{
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

	response := responses.UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Login:     user.Login,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}

func (u *User) Update(c *gin.Context) {
	login := c.Param("login")

	existing_user, err := u.Repo.FindOne(login)
	if err != nil {
		if errors.Is(err, repositories.ErrNotExist) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update user",
		})
		return
	}

	var request requests.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	existing_user.Firstname = request.Firstname
	existing_user.Lastname = request.Lastname
	existing_user.Login = request.Login

	if err := u.Repo.Update(login, &existing_user); err != nil {
		if errors.Is(err, repositories.ErrNotExist) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update user",
		})
		return
	}

	response := responses.UserResponse{
		ID:        existing_user.ID,
		Firstname: existing_user.Firstname,
		Lastname:  existing_user.Lastname,
		Login:     existing_user.Login,
		CreatedAt: existing_user.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
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
