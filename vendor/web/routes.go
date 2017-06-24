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

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
