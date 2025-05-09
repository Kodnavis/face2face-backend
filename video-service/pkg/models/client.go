package models

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID string
	Conn   *websocket.Conn
	Send   chan []byte
	Peer   *Client
}
