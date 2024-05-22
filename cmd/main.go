package main

import(
	"github.com/gin-gonic/gin"
	"github.com/akramov1ch/bookstore/internal/handlers"
)

func main(){
	router := gin.Default()
	router.GET("/authors",handlers.GetAuthors)
	router.GET("/authors/:authorid", handlers.GetAuthorsById)
	router.POST("/authors", handlers.CreateAuthor)
	router.PUT("/authors/:authorid", handlers.UpdateAuthor)
	router.DELETE("/authors/:authorid", handlers.DeleteAuthor)
	
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:bookid", handlers.GetBooksById)
	router.POST("/books", handlers.CreateBook)
	router.PUT("/books/:bookid", handlers.UpdateBook)
	router.DELETE("/books/:bookid", handlers.DeleteBook)

	router.Run(":8080")
}
// 85f0b898-f073-4030-853e-d9ea7f23e583