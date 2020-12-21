package main

import (
	"github.com/ederj98/gin-books/controllers"
	"github.com/ederj98/gin-books/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	/*r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})*/
	//Referencia a mi controlador
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id")

	r.Run()
}
