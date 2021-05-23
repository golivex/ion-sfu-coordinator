package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	loadtest "github.com/manishiitg/actions/loadtest"
	mirrorsfu "github.com/manishiitg/actions/mirror-sfu"
	log "github.com/pion/ion-log"
)

var mu sync.Mutex
var loadTestCancel = [](chan struct{}){}
var mirrorCancel = [](chan struct{}){}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/stopsyncsfu", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		for _, cancel := range mirrorCancel {
			close(cancel)
		}
		mirrorCancel = [](chan struct{}){}
	})
	r.GET("/syncsfu/:session1/:session2", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		cancel := make(chan struct{})
		go mirrorsfu.Init(c.Param("session1"), c.Param("session2"), cancel)
		mirrorCancel = append(mirrorCancel, cancel)
		c.Status(http.StatusOK)
	})
	r.GET("/syncsfu/:session1/:session2/:addr1/:addr2", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		cancel := make(chan struct{})
		go mirrorsfu.InitWithAddress(c.Param("session1"), c.Param("session2"), c.Param("addr1"), c.Param("addr2"), cancel)
		mirrorCancel = append(mirrorCancel, cancel)
		c.Status(http.StatusOK)
	})

	r.GET("/stopload", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		for _, cancel := range loadTestCancel {
			close(cancel)
		}
		loadTestCancel = [](chan struct{}){}

	})
	r.GET("/load/:session1", func(c *gin.Context) {
		clients := c.Query("clients")
		no := 1
		if clients != "" {
			x, err := strconv.Atoi(clients)
			if err != nil {
				log.Errorf("error string to int ", err)
			} else {
				no = x
			}
		}

		mu.Lock()
		defer mu.Unlock()
		cancel := make(chan struct{})
		go loadtest.InitApi(c.Param("session1"), no, cancel)
		loadTestCancel = append(loadTestCancel, cancel)
		c.Status(http.StatusOK)
	})

	r.Run(":3050")
}
