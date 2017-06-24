package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// StartServer yay
func StartServer() {
	router := gin.Default()

	router.StaticFS("/static", http.Dir("./client/dist"))
	router.GET("/stream", getStream)

	// Users
	router.POST("/api/users", createUser)

	// Channels
	router.POST("/api/channels", createChannel)
	router.POST("/api/add-channel-user", addChannelUser)
	router.POST("/api/remove-channel-user", removeChannelUser)

	// Users & Channels
	// router.POST("/api/channels/:id/users")

	// Music
	router.POST("/api/search", SearchMusic)

	// Query string of channel={channel_id}
	router.GET("/api/queue", GetQueue)
	router.POST("/api/queue", AddToQueue)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
