package web

import (
	"fmt"
	"os"

	"net/http"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/gin-contrib/static"
)

var channelMap = make(map[string]*Channel)

// StartServer yay
func StartServer() {
	router := gin.Default()

	router.Use(
		func(c *gin.Context) {
			urlPath := c.Request.URL.EscapedPath()
			if len(urlPath) == 0 {
				c.Next()
				return
			}
			if _, isChannelExists := channelMap[urlPath[1:]]; isChannelExists {
				data, err := ioutil.ReadFile("./client/dist/index.html")
				if err != nil {
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}
				c.Data(http.StatusOK, gin.MIMEHTML, data)
			}
		},
	)

	router.Use(static.Serve("/", static.LocalFile("./client/dist/", true)))
	// router.StaticFS("/", http.Dir("./client/dist/"))

	router.Use(
		func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
		},
	)

	// reverseProxy := httputil.NewSingleHostReverseProxy(&url.URL{
	// 	Scheme: "http",
	// 	Host:   "localhost:8000",
	// })
	//
	// router.Use(
	// 	func(c *gin.Context) {
	// 		if strings.HasPrefix(c.Request.URL.EscapedPath(), "/api") {
	// 			c.Next()
	// 		} else {
	// 			reverseProxy.ServeHTTP(c.Writer, c.Request)
	// 		}
	// 	},
	// )

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
