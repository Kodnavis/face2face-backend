package application

import (
	"net/http"

	"github.com/Kodnavis/face2face-backend/user-service/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	})

	router.Route("/users", loadUserRoutes)

	return router
}

func loadUserRoutes(router chi.Router) {
	userHandler := &handler.User{}

	router.Post("/", userHandler.Create)
	router.Get("/", userHandler.List)
	router.Get("/{id}", userHandler.GetByID)
	router.Put("/{id}", userHandler.UpdateById)
	router.Delete("/{id}", userHandler.DeleteById)
}
