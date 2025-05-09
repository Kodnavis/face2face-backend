package models

import "encoding/json"

type SignalMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type SDPMessage struct {
	SDP string `json:"sdp"`
}

type ICECandidate struct {
	Candidate string `json:"candidate"`
}
