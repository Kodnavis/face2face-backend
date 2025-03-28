package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

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
		pdb: db,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	defer func() {
		if err := a.pdb.Close(); err != nil {
			fmt.Println("failed to close PostgreSQL connection:", err)
		}
	}()

	fmt.Println("Starting server on :8080")

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}

		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
