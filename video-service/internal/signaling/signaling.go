package signaling

import (
	"encoding/json"
	"log"

	"github.com/Kodnavis/face2face-backend/video-service/internal/matcher"
	"github.com/Kodnavis/face2face-backend/video-service/pkg/models"
)

func HandleSignal(client *models.Client, raw []byte) {
	var msg models.SignalMessage
	if err := json.Unmarshal(raw, &msg); err != nil {
		log.Println("Failed to unmarshal signal message:", err)
		return
	}

	switch msg.Type {
	case "offer", "answer":
		var sdp models.SDPMessage
		if err := json.Unmarshal(msg.Data, &sdp); err != nil || sdp.SDP == "" {
			log.Println("Invalid SDP:", err)
			return
		}
	case "ice":
		var ice models.ICECandidate
		if err := json.Unmarshal(msg.Data, &ice); err != nil || ice.Candidate == "" {
			log.Println("Invalid ICE candidate:", err)
			return
		}
	case "skip":
		matcher.RequeueClient(client)
	}

	if client.Peer == nil {
		log.Println("Client has no peer yet")
		return
	}

	out, err := json.Marshal(msg)
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	client.Peer.Send <- out
}
