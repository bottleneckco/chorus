package web

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)

type ChannelID int

type Channel struct {
	ID          ChannelID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AccessCode  string    `json:"access_code"`
	CreatedBy   int       `json:"created_by"`
	Users       []User    `json:"users,omitempty"`
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

func getNextChannelID() ChannelID {
	return ChannelID(len(Channels) + 1)
}

func generateAccessCode() string {
	hd := hashids.NewData()
	hd.Salt = "random salt"
	hd.MinLength = 5
	h, _ := hashids.NewWithData(hd)

	e, _ := h.Encode([]int{rand.Int()})

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
		CreatedBy:   json.CreatedBy,
		Name:        json.Name,
		Description: json.Description,
		AccessCode:  generateAccessCode(),
	}

	Channels[getNextChannelID()] = channel
	c.JSON(http.StatusOK, channel)
}

func getChannel(c *gin.Context) {

}

func addChannelUser(c *gin.Context) {
	// var json ChannelUserPayload
	// err != c.BindJSON(&json)

}

func getChannelUsers(c *gin.Context) {

}

func getChannelQueue(c *gin.Context) {

}

func addChannelQueue(c *gin.Context) {

}
