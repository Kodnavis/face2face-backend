package handler

import (
	"log"
	"net/http"

	"github.com/Kodnavis/face2face-backend/video-service/internal/matcher"
	"github.com/Kodnavis/face2face-backend/video-service/internal/signaling"
	"github.com/Kodnavis/face2face-backend/video-service/pkg/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// authHeader := r.Header.Get("Authorization")
	// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
	// 	http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
	// 	return
	// }
	// jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

	// userID, err := auth.ExtractToken(jwtToken, os.Getenv("JWT_SECRET"))
	// if err != nil {
	// 	http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
	// 	return
	// }

	userID := "user_" + uuid.NewString()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
		return
	}

	client := &models.Client{
		UserID: userID,
		Conn:   conn,
		Send:   make(chan []byte),
	}

	go readLoop(client)
	go writeLoop(client)

	matcher.RegisterClient(client)
}

func readLoop(client *models.Client) {
	defer func() {
		matcher.RemoveClient(client)
		if client.Peer != nil {
			client.Peer.Peer = nil
		}
		client.Conn.Close()
		log.Printf("Client %s disconnected\n", client.UserID)
	}()

	for {
		_, raw, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		signaling.HandleSignal(client, raw)
	}
}

func writeLoop(client *models.Client) {
	defer client.Conn.Close()

	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
