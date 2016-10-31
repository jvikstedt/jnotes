package router

import (
	"github.com/jvikstedt/jnotes/jnotes"
	"github.com/pressly/chi"
	"net/http"
)

type Router struct {
	NoteController jnotes.NoteController
}

func (r *Router) Handler() http.Handler {
	chiRouter := chi.NewRouter()
	chiRouter.Post("/notes", r.NoteController.Create)
	return chiRouter
}
