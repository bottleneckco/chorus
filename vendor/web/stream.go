package web

import (
	"log"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	commandPause  = "pause"
	commandResume = "resume"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getStream(c *gin.Context) {
	channelID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotAcceptable)
		return
	}

	channel, isChannelExists := Channels[ChannelID(channelID)]
	if !isChannelExists {
		c.Status(http.StatusInternalServerError)
		return
	}

	userIDStr, err := c.GetCookie(cookieKeyUserID)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotAcceptable)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	_, isUserExists := channel.Users[userID]
	if !isUserExists {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
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

	go func(ws *websocket.Conn) {
		// Copy from channel stream
		for data := range channel.Stream {
			log.Printf("Sending out %d bytes\n", len(data))
			if err = ws.WriteMessage(websocket.BinaryMessage, data); err != nil {
				// Error writing, probably user disconnected

				// TODO: Delete the user from the channel
				// TODO: Need a better way to handle this. User refreshes page = lost forever.
				// delete(channel.Users, userID)
				log.Println(err)
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
			break
		case commandResume:
			break
		}
	}
}
