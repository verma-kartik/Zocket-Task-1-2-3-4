package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
)

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to CRUD API with Golang"})
}

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  string `json:"price"`
}

var books = []Book{
	{
		ID:     "1",
		Name:   "Frankenstein",
		Author: "Mark Shelly",
		Price:  "7.99",
	},
	{
		ID:     "2",
		Name:   "Jane Eyre",
		Author: "Charlotte Bronte",
		Price:  "4.99",
	},
}

func NewBookHandler(ctx *gin.Context) {
	var newBook Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newBook.ID = xid.New().String()
	books = append(books, newBook)
	ctx.JSON(http.StatusCreated, newBook)
}

func GetBooksHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, books)
}

func UpdateBookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var book Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Company not found",
		})
		return
	}
	books[index] = book
	ctx.JSON(http.StatusOK, book)
}

func DeleteBookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	index := -1
	for i := 0; i < len(books); i++ {
		if books[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}
	books = append(books[:index], books[index+1:]...)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book has been deleted",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", HomepageHandler)
	router.GET("/books", GetBooksHandler)
	router.POST("/book", NewBookHandler)
	router.PUT("/book/:id", UpdateBookHandler)
	router.DELETE("/book/:id", DeleteBookHandler)
	router.Run(":5555")

}
