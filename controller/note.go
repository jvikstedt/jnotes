package controller

import (
	"context"
	"encoding/json"
	"github.com/jvikstedt/jnotes/jnotes"
	"net/http"
	"strconv"
)

type NoteController struct {
	NoteRepository jnotes.NoteRepository
	Router         jnotes.Router
}

func (nc NoteController) BeforeFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noteIDStr := nc.Router.GetURLParameter(r, "noteID")
		noteID, err := strconv.Atoi(noteIDStr)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		note, err := nc.NoteRepository.FindByID(noteID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "note", note)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (nc NoteController) Create(w http.ResponseWriter, r *http.Request) {
	note := jnotes.Note{}
	defer r.Body.Close()
	var data struct {
		*jnotes.Note
		OmitID interface{} `json:"id,omitempty"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, http.StatusText(422), 422)
	}
	note, err = nc.NoteRepository.Create(note)
	if err != nil {
		http.Error(w, err.Error(), 422)
	}
	bytes, _ := json.Marshal(note)
	w.Write(bytes)
}

func (nc NoteController) Get(w http.ResponseWriter, r *http.Request) {
	note := r.Context().Value("note").(jnotes.Note)
	bytes, _ := json.Marshal(note)
	w.Write(bytes)
}
