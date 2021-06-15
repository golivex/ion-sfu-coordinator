package coordinator

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/pion/ion-log"
)

func (e *etcdCoordinator) InitApi() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/stats", func(c *gin.Context) {
		//TODO add infor about machines also here
		c.JSON(200, gin.H{
			"hosts":       e.hosts,
			"sessions":    e.sessions,
			"actionhosts": e.actionhosts,
		})
	})

	r.GET("/session/:id", func(c *gin.Context) {
		id := c.Param("id") //session name
		if id == "" {
			c.JSON(http.StatusOK, gin.H{
				// "Host": "", TODO
			})
			return
		}
		capacity := c.Query("capacity")
		cap := -1
		role := c.Query("role")

		if len(capacity) > 0 {
			x, err := strconv.Atoi(capacity)
			if err == nil {
				cap = x
			}
		}

		host := e.FindHost(id, cap, role)
		c.JSON(200, host)
	})
	r.GET("/action/status/:session/:action", func(c *gin.Context) {
		host, ac := e.queryActionStatus(c.Param("session"), c.Param("action"))
		if host != nil {
			c.JSON(http.StatusOK, gin.H{
				"host":         host.String(),
				"actionstatus": ac,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Response": "",
			})
		}
	})
	r.GET("/load/stats/:host/:port", func(c *gin.Context) {
		hosts := e.statsLoad(c.Param("host"), c.Param("port"))
		c.JSON(http.StatusOK, gin.H{
			"Response": hosts,
		})
	})
	r.GET("/stopload/:host/:port", func(c *gin.Context) {
		resp := e.stopSimLoad(c.Param("host"), c.Param("port"))
		c.JSON(http.StatusOK, gin.H{
			"Response": resp,
		})
	})

	r.GET("/load/stats", func(c *gin.Context) {
		hosts := e.statsLoadAll()
		c.JSON(http.StatusOK, gin.H{
			"Response": hosts,
		})
	})
	r.GET("/stopload", func(c *gin.Context) {
		hosts := e.stopAllSimLoad()
		c.JSON(http.StatusOK, gin.H{
			"Response": hosts,
		})
	})
	r.GET("/load/:session", func(c *gin.Context) {
		clients := c.Query("clients")
		no := 1
		if clients != "" {
			x, err := strconv.Atoi(clients)
			if err == nil {
				no = x
			}
		}
		role := c.Query("role")
		if len(role) == 0 || role == "pubsub" {
			role = "pubsub"
		} else {
			role = "sub"
		}

		file := c.Query("file")
		if len(file) == 0 {
			file = "default"
		} else {
			file = c.Query("file")
		}

		qcycle := c.Query("cycle")
		cycle := 0
		if len(qcycle) != 0 {
			x, err := strconv.Atoi(qcycle)
			if err == nil {
				cycle = x
			}
		}

		qrooms := c.Query("rooms")
		rooms := -1
		if len(qrooms) != 0 {
			x, err := strconv.Atoi(qrooms)
			if err == nil {
				rooms = x
			}
		}

		resp := e.simLoad(c.Param("session"), no, role, cycle, rooms, file)
		c.String(http.StatusOK, resp)
	})

	r.GET("/disk/:session/:filename", func(c *gin.Context) {
		session := c.Param("session")
		filename := c.Param("filename")
		resp := e.saveSessionToDisk(session, filename)
		c.String(http.StatusOK, resp)
	})
	r.GET("/stopdisk/:session", func(c *gin.Context) {
		resp := e.stopSessionToDisk(c.Param("session"))
		c.String(http.StatusOK, resp)
	})

	r.GET("/stream/:session", func(c *gin.Context) {
		session := c.Param("session")
		resp := e.startStream(session, c.Query("rtmp"))
		c.String(http.StatusOK, resp)
	})
	r.GET("/stopstream/:session", func(c *gin.Context) {
		resp := e.stopStream(c.Param("session"))
		c.String(http.StatusOK, resp)
	})

	r.GET("/rtmp/:session", func(c *gin.Context) {
		session := c.Param("session")
		resp := e.startRtmp(session, c.Query("rtmp"))
		c.String(http.StatusOK, resp)
	})
	r.GET("/stoprtmp/:session", func(c *gin.Context) {
		resp := e.stopRtmp(c.Param("session"))
		c.String(http.StatusOK, resp)
	})

	// r.GET("/load/:session/:host/:port", func(c *gin.Context) {
	// 	clients := c.Query("clients")
	// 	no := 1
	// 	if clients != "" {
	// 		x, err := strconv.Atoi(clients)
	// 		if err == nil {
	// 			log.Errorf("error string to int ", err)
	// 		} else {
	// 			no = x
	// 		}
	// 	}
	// 	role := c.Query("role")
	// 	if len(role) == 0 || role == "pubsub" {
	// 		role = "pubsub"
	// 	} else {
	// 		role = "sub"
	// 	}

	// 	qcycle := c.Query("cycle")
	// 	cycle := 0
	// 	if len(qcycle) != 0 {
	// 		x, err := strconv.Atoi(qcycle)
	// 		if err == nil {
	// 			cycle = x
	// 		}
	// 	}
	// 	qrooms := c.Query("rooms")
	// 	rooms := -1
	// 	if len(qrooms) != 0 {
	// 		x, err := strconv.Atoi(qrooms)
	// 		if err == nil {
	// 			rooms = x
	// 		}
	// 	}

	// 	file := c.Query("file")
	// 	if len(file) == 0 {
	// 		file = "default"
	// 	} else {
	// 		file = c.Query("file")
	// 	}

	// 	go e.simLoadForHost(c.Param("session"), c.Param("host"), c.Param("port"), no, role, cycle, rooms, file, 1, -1)
	// 	c.Status(http.StatusOK)
	// })
	// /session/test/node/5.9.18.28:7002/peer/ckoy35usg00080110qpo13b3v
	// r.GET("/clearsession", func(c *gin.Context) {
	// 	e.cli.Delete(context.Background(), "/session/", clientv3.WithPrefix())
	// })
	// r.GET("/simulate/session/:id/:host/:port", func(c *gin.Context) {
	// 	kvc := clientv3.NewKV(e.cli)
	// 	peerid := fmt.Sprintf("%v", rand.Intn(1000000000000000))
	// 	id := c.Param("id")
	// 	host := c.Param("host")
	// 	port := c.Param("port")
	// 	session := "/session/" + id + "/node/" + host + ":" + port + "/peer/" + peerid + "/track/" + peerid + "/"
	// 	kvc.Put(context.Background(), "/session/"+id+"/node/"+host+":"+port, "")
	// 	kvc.Put(context.Background(), session+"/video", "")
	// 	kvc.Put(context.Background(), session+"/audio", "")
	// 	c.String(200, session)
	// })
	log.Infof("Starting api server")
	r.Run(":4000")
}
