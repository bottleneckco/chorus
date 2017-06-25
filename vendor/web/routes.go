package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

var channelMap = make(map[string]*Channel)

// StartServer yay
func StartServer() {
	router := gin.Default()

	router.StaticFS("/sample", http.Dir("./client/sample"))

	router.Use(
		func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
		},
	)

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	apiR := router.Group("/api")
	{
		channelR := apiR.Group("/channels")
		{
			channelR.GET("/:id/search", searchMusic)

			channelR.GET("/:id", getChannel)
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
