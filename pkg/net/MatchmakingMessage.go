package net

import "github.com/gorilla/websocket"

type MatchmakingMessage struct {
	WebSocketConn *websocket.Conn
	Message       []byte
}
