package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Channels []Channel

type Channel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AccessCode  string `json:"access_code"`
	CreatedBy   string `json:"created_by"`
}

type CreateChannelPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

func (cc *CreateChannelPayload) validate() {
}

func getNextChannelID() int {
	return len(Channels) + 1
}

func createChannel(c *gin.Context) {
	var json CreateChannelPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to unmarshal json"})
	}

	// channel := Channel{
	// 	ID:          getNextChannelID(),
	// 	Name:        json.Name,
	// 	Description: json.Description,
	// 	CreatedBy:   json.CreatedBy,
	// }

	channel := Channel{}
	Channels = append(Channels, channel)

	// AllUsers = append(AllUsers, u)
	// c.JSON(http.StatusOK, u)
}
