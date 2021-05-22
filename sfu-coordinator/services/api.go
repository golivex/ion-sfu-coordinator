package coordinator

import (
	"net/http"

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
		e.mu.Lock()
		defer e.mu.Unlock()
		c.JSON(200, gin.H{
			"hosts":    e.hosts,
			"sessions": e.sessions,
		})
	})
	r.GET("/session/:id", func(c *gin.Context) {
		id := c.Param("id") //session name
		host := e.FindHost(id)
		c.JSON(200, host)
	})
	log.Infof("Starting api server")
	r.Run(":4000")
}
