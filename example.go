package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //r for router
	//r.SetTrustedProxies([]string{""}) //add an ip for squid proxy
	r.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "helloWorld",
		})
	})
	go r.Run()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	go r.Run(":8090")
	r.GET("/tacos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "areGood",
		})
	})
	r.Run(":8070")
}
