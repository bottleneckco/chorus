package web

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getStream(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer ws.Close()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			return
		}
		if err = ws.WriteMessage(websocket.TextMessage, []byte("Hello there, this is Go!")); err != nil {
			log.Println(err)
			break
		}
	}
}
