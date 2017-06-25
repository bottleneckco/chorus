package web

import "youtube"
import "github.com/gorilla/websocket"

const (
	statusOK      = "ok"
	statusError   = "error"
	commandPause  = "pause"
	commandResume = "resume"
	commandPing   = "ping"
)

// Payload models
type createChannelPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

type createUserPayload struct {
	Nickname string `json:"nickname" binding:"required"`
}

type addToChannelQueuePayload struct {
	URL string `json:"url"`
}

// Response models

type response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type channelResponse struct {
	response
	Channel Channel `json:"channel"`
}

type channelAddUserResponse struct {
	response
	User User `json:"user"`
}

type channelListUsersResponse struct {
	response
	Users []User `json:"users"`
}

type channelListQueueResponse struct {
	response
	Count int           `json:"count"`
	Queue []videoResult `json:"queue"`
}

type searchResponse struct {
	response
	Count   int           `json:"count"`
	Results []videoResult `json:"results"`
}

// Data models

type Channel struct {
	ID                string                          `json:"id"`
	Name              string                          `json:"name"`
	Description       string                          `json:"description"`
	CreatedBy         int                             `json:"created_by"`
	VideoResultsCache map[string]youtube.YoutubeVideo `json:"-"`
	Queue             []youtube.YoutubeVideo          `json:"-"`
	Users             map[int]User                    `json:"-"`
	UsersArray        []User                          `json:"users"` // UsersArray is only for display
	SkipCurrent       bool                            `json:"-"`     // Set this to true so the manager will abort distributing the current song on its next loop
}

func (c Channel) BroadcastMessage(messageType int, data []byte) error {
	var err error
	for _, user := range c.Users {
		err = user.WSConn.WriteMessage(messageType, data)
		if err != nil {
			return err
		}
	}
	return nil
}

type User struct {
	ID       int             `json:"id"`
	Nickname string          `json:"nickname"`
	WSConn   *websocket.Conn `json:"-"`
}

// WebSocket Command Interface models
type websocketCommand struct {
	Command string `json:"command"`
}
