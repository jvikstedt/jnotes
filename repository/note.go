package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/jnotes/jnotes"
)

type NoteRepository struct {
	DB *sqlx.DB
}

func (nr NoteRepository) Create(note jnotes.Note) (jnotes.Note, error) {
	fmt.Println(note)
	rows, err := nr.DB.NamedQuery(`INSERT INTO notes (title,body,createdAt) VALUES (:title,:body,CURRENT_TIMESTAMP) RETURNING *`, note)
	if rows.Next() {
		err = rows.StructScan(&note)
	}
	return note, err
}

func (nr NoteRepository) DeleteByID(id int) (note jnotes.Note, err error) {
	rows, err := nr.DB.Queryx("DELETE FROM notes where id=$1 RETURNING *", id)
	if rows.Next() {
		err = rows.StructScan(&note)
	}
	return note, err
}

func (nr NoteRepository) Delete(note jnotes.Note) (jnotes.Note, error) {
	return nr.DeleteByID(int(note.ID))
}

func (nr NoteRepository) Update(note jnotes.Note) (jnotes.Note, error) {
	rows, err := nr.DB.NamedQuery(`UPDATE notes SET (title,body,updatedAt) = (:title,:body,CURRENT_TIMESTAMP) WHERE id=:id RETURNING *`, note)
	if rows.Next() {
		err = rows.StructScan(&note)
	}
	return note, err
}

func (nr NoteRepository) FindByID(id int) (note jnotes.Note, err error) {
	err = nr.DB.Get(&note, "SELECT * FROM notes WHERE id=$1", id)
	return
}

func (nr NoteRepository) GetAll() (notes []jnotes.Note, err error) {
	err = nr.DB.Select(&notes, "SELECT * FROM notes")
	return
}
