package application

import (
	"github.com/Kodnavis/face2face-backend/user-service/handler"
	"github.com/Kodnavis/face2face-backend/user-service/repository"
	"github.com/gin-gonic/gin"
)

func (a *App) loadRoutes() {
	router := gin.Default()

	userGroup := router.Group("/users")
	a.loadUserRoutes(userGroup)

	a.router = router
}

func (a *App) loadUserRoutes(router *gin.RouterGroup) {
	userHandler := &handler.User{
		Repo: &repository.UserRepo{
			Db: a.pdb,
		},
	}

	router.POST("/", userHandler.Create)
	router.GET("/", userHandler.List)
	router.GET("/:id", userHandler.GetByID)
	router.PUT("/:id", userHandler.UpdateById)
	router.DELETE("/:id", userHandler.DeleteById)
}
