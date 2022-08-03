package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"

	"github.com/go_rest_api_playground/book"
	"github.com/go_rest_api_playground/database"
	"github.com/go_rest_api_playground/middleware"
)

func main() {

	r := gin.Default()
	database.ConnectDatabase()

	// grouping book api
	// eg: http://localhost:8080/bookapi/book
	bookapi := r.Group("/bookapi")
	{
		bookapi.Use(middleware.GroupEndPointMWLog())
		bookapi.GET("/books", middleware.SingleEndPointMWLog(), book.GetBooks)
		bookapi.GET("/book", book.GetBook)
		bookapi.POST("/book", book.AddBook)
		bookapi.DELETE("/book", book.DeleteBook)

	}

	cacheapi := r.Group("/cacheapi")
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

	r.Run(":8080") // listen and serve on "localhost:8080"
}
