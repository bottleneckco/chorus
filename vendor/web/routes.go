package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var Channels = make(map[ChannelID]Channel)

// StartServer yay
func StartServer() {
	router := gin.Default()

	router.StaticFS("/static", http.Dir("./client/dist"))

	apiR := router.Group("/api")
	{
		apiR.GET("/search", searchMusic)

		channelR := apiR.Group("/channels")
		{
			channelR.GET("/:id", getChannel)
			channelR.POST("", createChannel)

			channelR.GET("/:id/queue", getChannelQueue)
			channelR.POST("/:id/queue", addChannelQueue)

			channelR.GET("/:id/users", getChannelUsers)
			channelR.POST("/:id/users", addChannelUser)

			channelR.GET("/:id/stream", getStream)
		}
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
