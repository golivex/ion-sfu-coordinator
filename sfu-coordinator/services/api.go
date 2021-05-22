package coordinator

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/pion/ion-log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (e *etcdCoordinator) InitApi() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/stats", func(c *gin.Context) {
		e.mu.Lock()
		defer e.mu.Unlock()
		c.JSON(200, gin.H{
			"hosts":    e.hosts,
			"sessions": e.sessions,
		})
	})
	r.GET("/session/:id", func(c *gin.Context) {
		id := c.Param("id") //session name
		isaction := c.Query("action")
		host := e.FindHost(id, len(isaction) > 0)
		c.JSON(200, host)
	})
	r.GET("/stopload", func(c *gin.Context) {
		go StopSimLoad("0.0.0.0")
		c.Status(http.StatusOK)
	})
	r.GET("/stopload/:host", func(c *gin.Context) {
		go StopSimLoad(c.Param("host"))
		c.Status(http.StatusOK)
	})
	r.GET("/load/:session", func(c *gin.Context) {
		go SimLoad(c.Param("session"), "0.0.0.0")
		c.Status(http.StatusOK)
	})
	r.GET("/load/:session/:host", func(c *gin.Context) {
		go SimLoad(c.Param("session"), c.Param("host"))
		c.Status(http.StatusOK)
	})
	// /session/test/node/5.9.18.28:7002/peer/ckoy35usg00080110qpo13b3v
	r.GET("/clearsession", func(c *gin.Context) {
		e.cli.Delete(context.Background(), "/session/", clientv3.WithPrefix())
	})
	r.GET("/simulate/session/:id/:host/:port", func(c *gin.Context) {
		kvc := clientv3.NewKV(e.cli)
		peerid := fmt.Sprintf("%v", rand.Intn(1000000000000000))
		id := c.Param("id")
		host := c.Param("host")
		port := c.Param("port")
		session := "/session/" + id + "/node/" + host + ":" + port + "/peer/" + peerid + "/track/" + peerid + "/"
		kvc.Put(context.Background(), "/session/"+id+"/node/"+host+":"+port, "")
		kvc.Put(context.Background(), session+"/video", "")
		kvc.Put(context.Background(), session+"/audio", "")
		c.String(200, session)

	})
	log.Infof("Starting api server")
	r.Run(":4000")
}
