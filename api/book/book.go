package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_rest_api_playground/database"
	"github.com/go_rest_api_playground/model"
)

// @Summary      Show all books
// @Description  get all books
// @Tags         books
// @Produce      json
// @Success      200  {object}  []model.Book
// @Failure      400  {object} string
// @Failure      404  {string} string
// @Failure      500  {string} string
// @Router       /bookapi/getbooks [get]
func GetBooks(c *gin.Context) {
	var books []model.Book
	database.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// @Summary      Get a book
// @Description  Get a book by ID
// @Tags         books
// @Produce      json
// @Param 		id query int true "Enter Book ID"
// @Success      200  {object}  model.Book
// @Failure      400  {object} string
// @Failure      404  {string} string
// @Failure      500  {string} string
// @Router       /bookapi/getbook [get]
func GetBook(c *gin.Context) {
	var book model.Book

	id := c.Query("id")

	database.DB.Find(&book, id)

	c.JSON(http.StatusOK, gin.H{"data": book})

}

// @Summary      Add a book
// @Description  Add a book
// @Tags         books
// @Accept 		json
// @Produce      json
// @Param 		title query string true "Enter book title"
// @Param		author query string true "Enter book author"
// @Success      200  {object}  model.Book
// @Failure      400  {object} string
// @Failure      404  {string} string
// @Failure      500  {string} string
// @Router       /bookapi/addbook [post]
func AddBook(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	book := model.Book{Title: title, Author: author}
	database.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// @Summary      Delete a book
// @Description  Delete a book
// @Tags         books
// @Accept 		json
// @Produce      json
// @Param 		id query int true "Enter book id"
// @Success      200  {object}  model.Book
// @Failure      400  {object} string
// @Failure      404  {string} string
// @Failure      500  {string} string
// @Router       /bookapi/deletebook [delete]
func DeleteBook(c *gin.Context) {
	id := c.Query("id")
	database.DB.Delete(&model.Book{}, id)

	c.JSON(http.StatusOK, gin.H{"response": true})

}
