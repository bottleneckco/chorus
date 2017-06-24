package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var channelMap = make(map[ChannelID]Channel)

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.AbortWithStatus(200)
}

// StartServer yay
func StartServer() {
	router := gin.Default()

	router.StaticFS("/sample", http.Dir("./client/sample"))

	router.Use(
		func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
		},
	)

	apiR := router.Group("/api")
	{
		channelR := apiR.Group("/channels")
		{
			channelR.GET("/:id/search", searchMusic)

			channelR.GET("/:id", getChannel)
			channelR.OPTIONS("", preflight)
			channelR.POST("", createChannel)

			channelR.GET("/:id/queue", getChannelQueue)
			channelR.POST("/:id/queue", addToChannelQueue)
			channelR.DELETE("/:id/queue/:index", skipInChannelQueue)

			channelR.GET("/:id/users", getUsersInChannel)
			channelR.POST("/:id/users", addUserToChannel)

			channelR.GET("/:id/stream", getStream)
		}
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
