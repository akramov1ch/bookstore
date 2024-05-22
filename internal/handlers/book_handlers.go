package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/akramov1ch/bookstore/internal/models"
	"github.com/akramov1ch/bookstore/internal/services"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		var publicationDate time.Time
		if err := rows.Scan(&book.BookID, &book.Title, &book.AuthorID, &publicationDate, &book.ISBN, &book.Description, &book.CreatedAt, &book.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		book.PublicationDate, err = services.ParseDateStr(publicationDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBooksById(c *gin.Context) {
	var book models.Book
	var publicationDate time.Time
	bookid := c.Param("bookid")
	err := DB.QueryRow("SELECT * FROM books WHERE book_id = $1", bookid).Scan(&book.BookID, &book.Title, &book.AuthorID, &publicationDate, &book.ISBN, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	book.PublicationDate, err = services.ParseDateStr(publicationDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.CreatedAt = time.Now()
	newBook.UpdatedAt = time.Now()
	publicationDate, err := services.ParseDateTime(newBook.PublicationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.CheckAutorId(newBook.AuthorID, &DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = DB.Exec("INSERT INTO books (book_id, title, author_id, publication_date, isbn, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", newBook.BookID, newBook.Title, newBook.AuthorID, publicationDate, newBook.ISBN, newBook.Description, newBook.CreatedAt, newBook.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var book models.Book
	var publicationDate time.Time
	publicationDate, err := services.ParseDateTime(updatedBook.PublicationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.CheckAutorId(updatedBook.AuthorID, &DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	query := "SELECT book_id, title, author_id, publication_date, isbn, description, created_at, updated_at FROM books WHERE book_id = $1"
	err = DB.QueryRow(query, id).Scan(&book.BookID, &book.Title, &book.AuthorID, &publicationDate, &book.ISBN, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	updatedBook.BookID = book.BookID
	updatedBook.CreatedAt = book.CreatedAt
	updatedBook.UpdatedAt = time.Now()
	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM books WHERE book_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}
