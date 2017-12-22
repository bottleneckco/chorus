package model

import "github.com/gorilla/websocket"

const (
	ping     = "ping"
	userLeft = "userLeft"
)

type WSConn struct {
	Conn *websocket.Conn
}

func (ws *WSConn) WriteMessage(messageType int, data []byte) error {
	return ws.Conn.WriteMessage(messageType, data)
}

func (ws *WSConn) WriteText(data []byte) error {
	return ws.WriteMessage(websocket.TextMessage, data)
}

func (ws *WSConn) WriteJSON(data interface{}) error {
	return ws.Conn.WriteJSON(data)
}

func (ws *WSConn) WriteBinary(data []byte) error {
	return ws.WriteMessage(websocket.BinaryMessage, data)
}

func (ws *WSConn) Test() error {
	return ws.WriteJSON(map[string]string{"command": ping})
}
