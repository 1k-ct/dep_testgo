package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func router(gae bool) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello go"})
	})
	return r
}

// Init server initialization
func Init(gae bool) {
	r := router(gae)
	if err := r.Run(); err != nil {
		log.Println(err)
	}
	if gae {
		appengine.Main()
	}
}
