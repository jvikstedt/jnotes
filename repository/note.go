package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/jvikstedt/jnotes/jnotes"
)

type NoteRepository struct {
	DB *gorm.DB
}

func (nr NoteRepository) Create(note jnotes.Note) (jnotes.Note, error) {
	r := nr.DB.Create(&note)
	return note, r.Error
}

func (nr NoteRepository) DeleteByID(id int) (jnotes.Note, error) {
	note, err := nr.FindByID(id)
	if err != nil {
		return note, err
	}
	r := nr.DB.Delete(&note)
	return note, r.Error
}

func (nr NoteRepository) Delete(note jnotes.Note) (jnotes.Note, error) {
	r := nr.DB.Delete(&note)
	return note, r.Error
}

func (nr NoteRepository) Update(note jnotes.Note) (jnotes.Note, error) {
	r := nr.DB.Save(&note)
	return note, r.Error
}

func (nr NoteRepository) FindByID(id int) (jnotes.Note, error) {
	note := jnotes.Note{}
	r := nr.DB.First(&note, id)
	return note, r.Error
}
