package web

import (
	"log"
	"net/http"

	"encoding/json"
	"time"

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
	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  "Channel does not exist",
		})
		return
	}

	user, err := getUserByCookies(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response{
			Status: statusError,
			Error:  "Error retrieving cookies",
		})
		return
	}

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

	user.WSConn = ws
	channel.Users[user.ID] = user

	// Test alive function
	go func(ws *websocket.Conn) {
		// Copy from channel stream
		ticker := time.NewTicker(time.Second * 10)
		for range ticker.C {
			pingData, _ := json.Marshal(websocketCommand{
				Command: commandPing,
			})
			if err = ws.WriteMessage(websocket.TextMessage, pingData); err != nil {
				// Error writing, probably user disconnected
				log.Println(err)
				delete(channel.Users, user.ID)
				break
			}
		}
	}(ws)

	for {
		messageType, data, err := ws.ReadMessage()
		if err != nil {
			return
		}

		// Command-response code
		if messageType != websocket.TextMessage {
			log.Println("Client sent a non-textual message, ignored")
			return
		}

		switch string(data) {
		case commandPause:
			jsonData, _ := json.Marshal(websocketCommand{
				Command: commandPause,
			})
			channel.BroadcastMessage(websocket.TextMessage, jsonData)
			break
		case commandResume:
			jsonData, _ := json.Marshal(websocketCommand{
				Command: commandResume,
			})
			channel.BroadcastMessage(websocket.TextMessage, jsonData)
			break
		}
	}
}
