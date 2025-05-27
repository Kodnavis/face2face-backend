package main

import (
	"log"
	"net/http"

	"github.com/Kodnavis/face2face-backend/video-service/internal/handler"
)

func main() {
	http.HandleFunc("/ws", handler.WebSocketHandler)

	log.Println("Video service started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
