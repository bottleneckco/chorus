package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bottleneckco/chorus/model"

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

	user.WSConn = &model.WSConn{Conn: ws}
	channel.AddUser(&user)
	// channel.I.Users[user.ID] = user

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

		fmt.Printf("websocket: received `%s` from client\n", string(data))

		switch string(data) {
		case commandPause:
			channel.IssueCommand(model.ChannelCommandPause)
			channel.BroadcastJSON(websocketCommand{Command: commandPause})
			break
		case commandResume:
			channel.IssueCommand(model.ChannelCommandResume)
			channel.BroadcastJSON(websocketCommand{Command: commandResume})
			break
		case commandSkipCurrent:
			channel.IssueCommand(model.ChannelCommandSkip)
			break
		}
	}
}
