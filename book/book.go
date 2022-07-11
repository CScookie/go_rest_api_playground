package book

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `josn:"rating`
}

var books = []Book{
	{ID: "1", Title: "harry potter", Author: "abc", Rating: 5},
	{ID: "2", Title: "GoT", Author: "def", Rating: 5},
}

func GetBooks(c *gin.Context) {
	log.Println("Getting books")
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	log.Println("Getting book")

	for _, book := range books {
		if id == book.ID {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book of ID: " + id + " not found"})
}

func AddBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		log.Fatal(err)
	}

	books = append(books, newBook)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Book added",
		"response": newBook,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if id == book.ID {
			books = append(books[:index], books[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message":  "Book removed",
				"response": books,
			})

			return

		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book of ID: " + id + " not found"})

}
