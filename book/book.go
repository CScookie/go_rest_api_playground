package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_rest_api_playground/database"
	"github.com/go_rest_api_playground/model"
)

func GetBooks(c *gin.Context) {
	var books []model.Book
	database.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBook(c *gin.Context) {
	var book model.Book

	id := c.Query("id")

	database.DB.Find(&book, id)

	c.JSON(http.StatusOK, gin.H{"data": book})

}

func AddBook(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	book := model.Book{Title: title, Author: author}
	database.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	id := c.Query("id")
	database.DB.Delete(&model.Book{}, id)

	c.JSON(http.StatusOK, gin.H{"response": true})

}
