package main

import(
	"github.com/gin-gonic/gin"
	"github.com/akramov1ch/bookstore/internal/handlers"
)

func main(){
	router := gin.Default()
	router.GET("/authors",handlers.GetAuthors)
	router.POST("/authors", handlers.CreateAuthor)
	router.PUT("/authors/:authorid", handlers.UpdateAuthor)
	router.DELETE("/authors/:authorid", handlers.DeleteAuthor)
	
	router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:bookid", handlers.UpdateBook)
	router.DELETE("/books/:bookid", handlers.DeleteBook)

	router.Run(":8080")
}