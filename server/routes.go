package server

import (
	"github.com/gin-gonic/gin"
	"github.com/williampiv/venstar-cli/api"
	"net/http"
)

func initializeRoutes(ip string) {
	router.GET("/", showIndexPage)
	router.GET("/acOn", turnOnAC(ip))
}

func showIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
	})
}

func turnOnAC(ip string) gin.HandlerFunc {
	if ip != "" {
		info := api.GetThermostatInfo(ip)
		api.SetCoolTemp(ip, 58, info)
		api.SetThermostatMode(ip, api.ConvertThermostatMode("cool"), info)
	}
	return showIndexPage
}
