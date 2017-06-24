package web

import "youtube"

const (
	statusOK    = "ok"
	statusError = "error"
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
	Error  error  `json:"error,omitempty"`
}

type channelResponse struct {
	response
	Channel Channel `json:"channel"`
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

type ChannelID int

type Channel struct {
	ID                ChannelID                       `json:"id"`
	Name              string                          `json:"name"`
	Description       string                          `json:"description"`
	AccessCode        string                          `json:"access_code"`
	CreatedBy         int                             `json:"created_by"`
	VideoResultsCache map[string]youtube.YoutubeVideo `json:"-"`
	Queue             []youtube.YoutubeVideo          `json:"-"`
	Stream            chan []byte                     `json:"-"`
	Users             map[int]User                    `json:"-"`
	UsersArray        []User                          `json:"users"` // UsersArray is only for display
}

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}
