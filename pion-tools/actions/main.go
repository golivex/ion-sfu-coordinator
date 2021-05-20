package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manishiitg/actions/mirror-sfu/client"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/syncsfu/:session1/:session2/:addr1", func(c *gin.Context) {
		go client.Init(c.Param("session1"), c.Param("addr1"), c.Param("session2"), c.Param("addr1"))
		c.String(http.StatusOK)
	})
	r.GET("/syncsfu/:session1/:session2/:addr1/*addr2", func(c *gin.Context) {
		go client.Init(c.Param("session1"), c.Param("addr1"), c.Param("session2"), c.Param("addr2"))
		c.String(http.StatusOK)
	})
	r.Run(":3050")
}
