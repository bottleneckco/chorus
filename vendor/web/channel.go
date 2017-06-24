package web

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)

var Channels []Channel

type Channel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AccessCode  string `json:"access_code"`
	CreatedBy   int    `json:"created_by"`
	Users       []User `json:"users"`
}

type CreateChannelPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
}

type ChannelUserPayload struct {
	ChannelID int `json:"channel_id"`
	UserID    int `json:"user_id"`
}

func (cc *CreateChannelPayload) validate() {

}

func getNextChannelID() int {
	return len(Channels) + 1
}

func generateAccessCode() string {
	hd := hashids.NewData()
	hd.Salt = "random salt"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{getNextChannelID(), rand.Int()})

	return e
}

func createChannel(c *gin.Context) {
	var json CreateChannelPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to unmarshal json"})
	}

	channel := Channel{
		ID:          getNextChannelID(),
		Name:        json.Name,
		Description: json.Description,
		CreatedBy:   json.CreatedBy,
		AccessCode:  generateAccessCode(),
	}

	Channels = append(Channels, channel)
	c.JSON(http.StatusOK, channel)
}

func addChannelUser(c *gin.Context) {
	// var json ChannelUserPayload
	// err != c.BindJSON(&json)

}

func removeChannelUser(c *gin.Context) {

}
