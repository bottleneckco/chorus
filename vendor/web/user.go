package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var AllUsers []User

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type CreateUser struct {
	Nickname string `json:"nickname" binding:"required"`
}

func getNextUserID() int {
	return len(AllUsers) + 1
}

func createUser(c *gin.Context) {
	var json CreateUser
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

	AllUsers = append(AllUsers, u)
	c.JSON(http.StatusOK, u)
}
