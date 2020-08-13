package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

// StartApp is the entrypoint of this application
func StartApp() {
	mapUrls()
	if err := router.Run(":3000"); err != nil {
		panic(err)
	}
	fmt.Println("Server up and running on port 3000")
}
