package matcher

import (
	"fmt"
	"log"
	"sync"

	"github.com/Kodnavis/face2face-backend/video-service/pkg/models"
)

var (
	queueMu        sync.Mutex
	waitingClients []*models.Client
)

func RegisterClient(client *models.Client) {
	queueMu.Lock()
	defer queueMu.Unlock()

	if len(waitingClients) > 0 {
		peer := waitingClients[0]
		waitingClients = waitingClients[1:]

		client.Peer = peer
		peer.Peer = client

		log.Printf("Matched client %s with %s\n", client.UserID, peer.UserID)
		go notifyMatched(peer, client)
	} else {
		waitingClients = append(waitingClients, client)
		log.Printf("Client %s added to queue\n", client.UserID)
	}
}

func notifyMatched(caller, receiver *models.Client) {
	callerMsg := []byte(fmt.Sprintf(`{"type":"match_found","role":"caller","userName":"%s"}`, caller.UserID))
	receiverMsg := []byte(fmt.Sprintf(`{"type":"match_found","role":"receiver","userName":"%s"}`, receiver.UserID))

	caller.Send <- callerMsg
	receiver.Send <- receiverMsg
}

func RemoveClient(target *models.Client) {
	queueMu.Lock()
	defer queueMu.Unlock()

	for i, client := range waitingClients {
		if client == target {
			waitingClients = append(waitingClients[:i], waitingClients[i+1:]...)
			log.Printf("Removed disconnected client %s from queue\n", client.UserID)
			return
		}
	}
}

func RequeueClient(client *models.Client) {
	if client.Peer != nil {
		client.Peer.Peer = nil
		client.Peer = nil
	}

	RemoveClient(client)
	RegisterClient(client)
}
