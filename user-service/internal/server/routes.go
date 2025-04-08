package server

import (
	"os"

	"github.com/Kodnavis/face2face-backend/common/auth"
	"github.com/Kodnavis/face2face-backend/user-service/internal/handlers"
	"github.com/Kodnavis/face2face-backend/user-service/internal/repositories"
	"github.com/gin-gonic/gin"
)

func (a *App) loadRoutes() {
	router := gin.Default()

	userGroup := router.Group("/users")
	a.loadUserRoutes(userGroup)

	a.router = router
}

var jwt_secret string = os.Getenv("TOKEN_LIFESPAN")

func (a *App) loadUserRoutes(router *gin.RouterGroup) {
	userHandler := &handlers.User{
		Repo: &repositories.UserRepository{
			DB: a.pdb,
		},
	}

	router.POST("/", userHandler.Create)
	router.GET("/", userHandler.List)
	router.GET("/:login", userHandler.Get)
	router.PUT("/:login", auth.JwtAuthMiddleware(jwt_secret), userHandler.Update)
	router.DELETE("/:login", auth.JwtAuthMiddleware(jwt_secret), userHandler.Delete)

	router.POST("/login", userHandler.Login)
}
