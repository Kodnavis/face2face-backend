package main

import (
	"log"

	"github.com/Kodnavis/face2face-backend/user-service/internal/database"
	"github.com/Kodnavis/face2face-backend/user-service/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connect()

	db.AutoMigrate(&models.User{})

	log.Println("migrate the schemas finished")
}
