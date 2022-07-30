package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rromulos/go-rest-api/database"
	"github.com/rromulos/go-rest-api/models"
)

func GetBook(c *gin.Context) {
	db := database.GetDatabase()
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID has to be integer",
		})
		return
	}

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Book not found: " + err.Error(),
		})

		return
	}
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Book

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error creating the book: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func GetAllBooks(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Book
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot find book with id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}