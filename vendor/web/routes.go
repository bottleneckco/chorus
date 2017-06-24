package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// StartServer yay
func StartServer() {
	c := gin.Default()
	c.StaticFS("/static", http.Dir("./client/dist"))
	c.GET("/stream", getStream)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	c.Run(fmt.Sprintf(":%s", port))
}
