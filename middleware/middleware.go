package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SingleEndPointMWLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Single end point middle ware log is working")
		c.String(http.StatusOK, "Single end point middle ware log is working\n")
	}
}

func GroupEndPointMWLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Println(c.GetHeader("User-Agent"))
		log.Println("Group end point middle ware log is working")

		// Before calling handler
		log.Println("Before c.next()")

		c.Next() // calls the next handler

		// After calling handler
		log.Println("After c.next()")
	}
}
