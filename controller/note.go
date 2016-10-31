package controller

import (
	"encoding/json"
	"github.com/jvikstedt/jnotes/jnotes"
	"net/http"
)

type NoteController struct {
	NoteRepository jnotes.NoteRepository
}

func (nc NoteController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	note := jnotes.Note{}
	err := decoder.Decode(&note)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
	}
	note, err = nc.NoteRepository.Create(note)
	if err != nil {
		http.Error(w, err.Error(), 422)
	}
}
