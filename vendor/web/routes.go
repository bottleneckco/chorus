package web

import "github.com/gin-gonic/gin"
import "os"
import "fmt"

// StartServer yay
func StartServer() {
	c := gin.Default()
	c.GET("/stream", getStream)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	c.Run(fmt.Sprintf(":%s", port))
}
