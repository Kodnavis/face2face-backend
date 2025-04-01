package application

import (
	"log"

	"github.com/Kodnavis/face2face-backend/user-service/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	router *gin.Engine
	pdb    *gorm.DB
}

func New() *App {
	db := database.Connect()

	app := &App{
		pdb: db,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start() error {
	err := a.router.Run()
	if err != nil {
		return err
	}

	log.Println("listen and serve on :8080")

	return nil
}
