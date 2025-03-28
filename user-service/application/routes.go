package application

import (
	"net/http"

	"github.com/Kodnavis/face2face-backend/user-service/handler"
	"github.com/Kodnavis/face2face-backend/user-service/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	})

	router.Route("/users", a.loadUserRoutes)

	a.router = router
}

func (a *App) loadUserRoutes(router chi.Router) {
	userHandler := &handler.User{
		Repo: &repository.UserRepo{
			Db: a.pdb,
		},
	}

	router.Post("/", userHandler.Create)
	router.Get("/", userHandler.List)
	router.Get("/{id}", userHandler.GetByID)
	router.Put("/{id}", userHandler.UpdateById)
	router.Delete("/{id}", userHandler.DeleteById)
}
