package main

import (
	"log"

	"github.com/Kodnavis/face2face-backend/user-service/application"
)

func init() {
	application.LoadEnvVariables()
}

func main() {
	app := application.New()

	err := app.Start()
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
