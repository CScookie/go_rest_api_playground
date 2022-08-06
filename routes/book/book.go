package book

import (
	"github.com/gin-gonic/gin"

	"github.com/go_rest_api_playground/api/book"
	"github.com/go_rest_api_playground/middleware"
	// swagger embed files
)

func Routes(rg *gin.RouterGroup) {
	// grouping book api
	// eg: http://localhost:8080/bookapi/book
	bookapi := rg.Group("/bookapi")
	{
		bookapi.Use(middleware.GroupEndPointMWLog())
		bookapi.GET("/getbooks", middleware.SingleEndPointMWLog(), book.GetBooks)
		bookapi.GET("/getbook", book.GetBook)
		bookapi.POST("/addbook", book.AddBook)
		bookapi.DELETE("/deletebook", book.DeleteBook)

	}
}
