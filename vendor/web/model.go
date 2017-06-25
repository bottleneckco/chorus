package web

import "model"

const (
	statusOK           = "ok"
	statusError        = "error"
	commandPause       = "pause"
	commandResume      = "resume"
	commandPing        = "ping"
	commandSkipCurrent = "skipCurrent"
	commandUsersLeft   = "usersLeft"
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
	Channel model.Channel `json:"channel"`
}

type channelAddUserResponse struct {
	response
	User model.User `json:"user"`
}

type channelListUsersResponse struct {
	response
	Users []*model.User `json:"users"`
}

type channelListQueueResponse struct {
	response
	Count int                    `json:"count"`
	Queue []channelListQueueItem `json:"queue"`
}

type channelListQueueItem struct {
	ID          int         `json:"id"`
	User        model.User  `json:"user"`
	VideoResult videoResult `json:"video"`
}

type searchResponse struct {
	response
	Count   int           `json:"count"`
	Results []videoResult `json:"results"`
}

// WebSocket Command Interface models
type websocketCommand struct {
	Command string `json:"command"`
}

type websocketDataCommand struct {
	websocketCommand
	Data interface{} `json:"data"`
}
