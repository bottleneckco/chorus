package web

import (
	"math/rand"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	hashids "github.com/speps/go-hashids"
)

func generateAccessCode() string {
	hd := hashids.NewData()
	hd.Salt = "random salt"
	hd.MinLength = 5
	h, _ := hashids.NewWithData(hd)

	e, _ := h.Encode([]int{rand.Int()})

	return e
}

func getChannelIDFromParam(c *gin.Context) ChannelID {
	id, _ := strconv.Atoi(c.Param("id"))
	return ChannelID(id)
}

func getNextChannelID() ChannelID {
	return ChannelID(len(channelMap) + 1)
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
