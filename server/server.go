package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

// Entry is the entrypoint for the server, to run in server mode
func Entry(ip string) {
	log.Println("Starting Server Engine...")
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes(ip)

	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}

}
