package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/akramov1ch/bookstore/internal/db"
	"github.com/akramov1ch/bookstore/internal/models"
	"github.com/akramov1ch/bookstore/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB sql.DB = *db.DbConnect()

func GetAuthors(c *gin.Context) {
	var authors []models.Author

	rows, err := DB.Query("SELECT * FROM authors")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var author models.Author
		var birthdate time.Time
		if err := rows.Scan(&author.AuthorID, &author.Name, &birthdate, &author.Biography, &author.CreatedAt, &author.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		author.BirthDate, err = services.ParseDateStr(birthdate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAuthor.AuthorID = uuid.New()
	newAuthor.CreatedAt = time.Now()
	newAuthor.UpdatedAt = time.Now()
	birthdate, err := services.ParseDateTime(newAuthor.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = DB.Exec("INSERT INTO authors (author_id, name, birth_date, biography, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", newAuthor.AuthorID, newAuthor.Name, birthdate, newAuthor.Biography, newAuthor.CreatedAt, newAuthor.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAuthor)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var updatedAuthor models.Author
	if err := c.ShouldBindJSON(&updatedAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var author models.Author
	_, err := DB.Query("SELECT * FROM authors WHERE author_id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}
	birthdate, err := services.ParseDateTime(updatedAuthor.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedAuthor.AuthorID = author.AuthorID
	updatedAuthor.CreatedAt = author.CreatedAt
	updatedAuthor.UpdatedAt = time.Now()
	_, err = DB.Exec("UPDATE authors SET name = $1, birth_date = $2, biography = $3, updated_at = $4 WHERE author_id = $5", updatedAuthor.Name, birthdate, updatedAuthor.Biography, updatedAuthor.UpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedAuthor)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM authors WHERE author_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "author deleted"})
}
