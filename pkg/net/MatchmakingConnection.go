package net

import "github.com/gorilla/websocket"

type MatchmakingConnection struct {
	WebSocketConn *websocket.Conn
	Entity        EntityInterface
}
