package main

import (
	"github.com/gin-gonic/gin"

	"github.com/go_rest_api_playground/book"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/book", book.GetBooks)
	r.GET("/api/v1/book/:id", book.GetBook)
	r.POST("/api/v1/book", book.AddBook)
	r.DELETE("/api/v1/book/:id", book.DeleteBook)

	r.Run() // listen and serve on "localhost:8080"
}
