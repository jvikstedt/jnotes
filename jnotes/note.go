package jnotes

import (
	"net/http"
	"time"
)

type Note struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

type NoteParams struct {
	Note
	OmitID        interface{} `json:"id,omitempty"`
	OmitCreatedAt interface{} `json:"created_at,omitempty"`
	OmitUpdatedAt interface{} `json:"updated_at,omitempty"`
}

type NoteRepository interface {
	Create(Note) (Note, error)
	DeleteByID(int) (Note, error)
	Delete(Note) (Note, error)
	Update(Note) (Note, error)
	FindByID(int) (Note, error)
	GetAll() ([]Note, error)
}

type NoteController interface {
	BeforeFilter(next http.Handler) http.Handler
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
