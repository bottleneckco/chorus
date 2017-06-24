package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Users []User

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type CreateUserPayload struct {
	Nickname string `json:"nickname" binding:"required"`
}

func getNextUserID() int {
	return len(Users) + 1
}

func createUser(c *gin.Context) {
	var json CreateUserPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to unmarshal json"})
	}

	if len(json.Nickname) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nickname is empty"})
	}

	u := User{
		ID:       getNextUserID(),
		Nickname: json.Nickname,
	}

	Users = append(Users, u)
	c.JSON(http.StatusOK, u)
}
