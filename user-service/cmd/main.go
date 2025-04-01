package main

import (
	"log"

	"github.com/Kodnavis/face2face-backend/user-service/internal/server"
)

func main() {
	app := server.New()

	err := app.Start()
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
