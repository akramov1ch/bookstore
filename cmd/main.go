package main

import(
	"github.com/gin-gonic/gin"
	"github.com/akramov1ch/bookstore/internal/handlers"
)

func main(){
	router := gin.Default()
	router.GET("/authors", func(c *gin.Context){
		handlers.GetAuthors(c)
	})
	router.POST("/authors", func(c *gin.Context){
		handlers.CreateAuthor(c)
	})
	router.PUT("/authors/:authorid", func(c *gin.Context){
		handlers.UpdateAuthor(c)
	})
	router.DELETE("/authors/:authorid", func(c *gin.Context){
		handlers.DeleteAuthor(c)
	})
	router.GET("/books", func(c *gin.Context){
		handlers.GetBooks(c)
	})
	router.POST("/books", func(c *gin.Context){
		handlers.CreateBook(c)
	})
	router.PUT("/books/:bookid", func(c *gin.Context){
		handlers.UpdateBook(c)
	})
	router.DELETE("/books/:bookid", func(c *gin.Context){
		handlers.DeleteBook(c)
	})

	router.Run(":8080")
}