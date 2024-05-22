package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/akramov1ch/bookstore/internal/models"
	"github.com/google/uuid"
)

func ParseDateTime(datestr string) (time.Time, error) {
	var date time.Time
	if datestr != "" {
		parsedDate, err := time.Parse("2006-01-02", datestr)
		if err != nil {
			return date, err
		}
		date = parsedDate
	}
	return date, nil
}

func ParseDateStr(dateTime time.Time) (string, error) {
	var empty = time.Time{}
	if dateTime != empty {
		return dateTime.Format("2006-01-02"), nil
	}
	return "", fmt.Errorf("Wrong data!")
}

func CheckAutorId(authorid uuid.UUID, DB *sql.DB) (error) {
	var author models.Author
	query := "SELECT author_id, name, birth_date, biography, created_at, updated_at FROM authors WHERE author_id = $1"
	err := DB.QueryRow(query, authorid).Scan(&author.AuthorID, &author.Name, &author.BirthDate, &author.Biography, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	return nil
}

func CheckAuthorsBook(author models.Author, DB *sql.DB) (error) {
	var book models.Book
	query := "SELECT book_id, title, author_id, publication_date, isbn, description, created_at, updated_at FROM books WHERE author_id = $1"
	err := DB.QueryRow(query, author.AuthorID).Scan(&book.BookID, &book.Title, &book.AuthorID, &book.PublicationDate, &book.ISBN, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	return nil
}