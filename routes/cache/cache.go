package cache

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func Routes(rg *gin.RouterGroup) {
	cacheapi := rg.Group("/cacheapi")
	{
		store := persistence.NewInMemoryStore(60 * time.Second)
		// Cached Page

		cacheapi.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
		})

		cacheapi.GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *gin.Context) {
			c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
		}))
	}
}
