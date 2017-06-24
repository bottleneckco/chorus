package web

const (
	statusOK    = "ok"
	statusError = "error"
)

type response struct {
	Status string `json:"status"`
	Error  error  `json:"error,omitempty"`
}

type channelCreateResponse struct {
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
