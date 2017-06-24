package web

import (
	"math/rand"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)

const (
	cookieKeyUserID = "user_id"
)

type ChannelID int

type Channel struct {
	ID          ChannelID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	AccessCode  string       `json:"access_code"`
	CreatedBy   int          `json:"created_by"`
	Stream      chan []byte  `json:"-"`
	Users       map[int]User `json:"-"`
	UsersArray  []User       `json:"users"` // UsersArray is only for display
}

type CreateChannelPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
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

func setUserCookie(newUserID int, c *gin.Context) {
	newUserIDStr := strconv.Itoa(newUserID)

	c.SetCookie(
		cookieKeyUserID,
		newUserIDStr,
		0,
		"/channel",
		"",
		false,
		false,
	)
}

func formatUsersForJson(users map[int]User) []User {
	var usersArr []User

	for _, v := range users {
		usersArr = append(usersArr, v)
	}

	sort.Slice(usersArr, func(i, j int) bool {
		return usersArr[i].ID < usersArr[j].ID
	})

	return usersArr
}

func createChannel(c *gin.Context) {
	var json CreateChannelPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to unmarshal json"})
	}

	users := make(map[int]User)
	newUserID := 1
	createdByUser := User{
		ID:       newUserID,
		Nickname: json.CreatedBy,
	}

	users[newUserID] = createdByUser

	channel := Channel{
		ID:          getNextChannelID(),
		CreatedBy:   newUserID,
		Name:        json.Name,
		Description: json.Description,
		AccessCode:  generateAccessCode(),
		Users:       users,
	}

	setUserCookie(newUserID, c)
	Channels[getNextChannelID()] = channel

	// Populate usersArr for view
	channel.UsersArray = formatUsersForJson(users)

	c.JSON(http.StatusOK, channel)
}

func getChannelIDFromParam(c *gin.Context) ChannelID {
	id, _ := strconv.Atoi(c.Param("id"))
	return ChannelID(id)
}

func getChannel(c *gin.Context) {
	channelID := getChannelIDFromParam(c)
	c.JSON(http.StatusOK, Channels[channelID])
}

func addChannelUser(c *gin.Context) {
	channelID := getChannelIDFromParam(c)

	var json CreateUserPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to unmarshal json"})
	}

	channel := Channels[channelID]
	users := channel.Users

	newUserID := len(users) + 1
	newUser := User{
		ID:       newUserID,
		Nickname: json.Nickname,
	}

	users[newUserID] = newUser

	channel.Users = users
	Channels[channelID] = channel

	// Populate usersArr for view
	channel.UsersArray = formatUsersForJson(users)

	c.JSON(http.StatusOK, channel)
}

func getChannelUsers(c *gin.Context) {
	channelID := getChannelIDFromParam(c)
	c.JSON(http.StatusOK, formatUsersForJson(Channels[channelID].Users))
}

func getChannelQueue(c *gin.Context) {
}

func addChannelQueue(c *gin.Context) {

}
