package jnotes

import "time"

type Note struct {
	ID        uint       `gorm:"primary_key",json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
}

type NoteRepository interface {
	Create(Note) (Note, error)
	DeleteByID(int) (Note, error)
	Delete(Note) (Note, error)
	Update(Note) (Note, error)
	FindByID(int) (Note, error)
}
