package models

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	AuthorID  uuid.UUID `json:"author_id"`
	Name      string    `json:"name"`
	BirthDate string `json:"birth_date"`
	Biography string    `json:"biography"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
