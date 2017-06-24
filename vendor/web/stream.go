package web

import (
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getStream(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, http.Header{
		"Sec-Websocket-Protocol": []string{
			c.Request.Header.Get("Sec-Websocket-Protocol"),
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	defer ws.Close()

	audioFile, err := ioutil.ReadFile("free.mp3")
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			return
		}
		if err = ws.WriteMessage(websocket.BinaryMessage, audioFile); err != nil {
			log.Println(err)
			break
		}
	}
}
