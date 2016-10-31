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
	defer r.Body.Close()

	var noteParams jnotes.NoteParams
	err := json.NewDecoder(r.Body).Decode(&noteParams)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
	}

	note, err := nc.NoteRepository.Create(noteParams.Note)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	RenderJSON(w, http.StatusCreated, note)
}

func (nc NoteController) Get(w http.ResponseWriter, r *http.Request) {
	note := r.Context().Value("note").(jnotes.Note)
	RenderJSON(w, http.StatusOK, note)
}
