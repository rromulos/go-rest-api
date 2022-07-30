package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rromulos/go-rest-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.GetBook)
			books.GET("/", controllers.GetAllBooks)
			books.POST("/", controllers.CreateBook)
		}
	}

	return router
}
