package web

import (
	"math/rand"
	"sort"
	"strconv"

	"github.com/bottleneckco/chorus/model"

	"github.com/gin-gonic/gin"
	hashids "github.com/speps/go-hashids"
)

func generateAccessCode() string {
	hd := hashids.NewData()
	hd.Salt = "random salt"
	hd.MinLength = 5
	h := hashids.NewWithData(hd)

	e, _ := h.Encode([]int{rand.Int()})

	return e
}

func getUserByCookies(c *gin.Context) (model.User, error) {
	var user model.User
	userIDStr, err := c.Cookie(cookieKeyUserID)
	if err != nil {
		return user, err
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return user, err
	}
	user.ID = userID
	userNick, err := c.Cookie(cookieKeyUserNickname)
	if err != nil {
		return user, err
	}
	user.Nickname = userNick
	return user, nil
}

func setUserCookie(user model.User, c *gin.Context) {
	c.SetCookie(
		cookieKeyUserID,
		strconv.Itoa(user.ID),
		0,
		"/",
		"",
		false,
		false,
	)

	c.SetCookie(
		cookieKeyUserNickname,
		user.Nickname,
		0,
		"/",
		"",
		false,
		false,
	)
}

func formatUsersForJson(users map[int]*model.User) []*model.User {
	var usersArr = make([]*model.User, 0)

	for _, v := range users {
		usersArr = append(usersArr, v)
	}

	sort.Slice(usersArr, func(i, j int) bool {
		return usersArr[i].ID < usersArr[j].ID
	})

	return usersArr
}
