package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Kodnavis/face2face-backend/user-service/database"
	_ "github.com/lib/pq"
)

type App struct {
	router http.Handler
	pdb    *sql.DB
}

func New() *App {
	db := database.Connect()

	app := &App{
		router: loadRoutes(),
		pdb:    db,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	fmt.Println("Starting server on :8080")

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
