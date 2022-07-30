package controllers

import "github.com/gin-gonic/gin"

func ShowBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID has to be integer",
		})
		return
	}

	database := database.getDatabase()
	var book models.Book
	err = database.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Book not found: " + err.Error(),
		})

		return
	}
	c.JSON(200, book)
}