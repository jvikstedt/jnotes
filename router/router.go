package router

import (
	"github.com/jvikstedt/jnotes/jnotes"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"net/http"
)

type Router struct {
	NoteController jnotes.NoteController
}

func (router Router) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/v1", router.ApiV1())

	return r
}

func (router Router) GetURLParameter(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func (router Router) ApiV1() http.Handler {
	r := chi.NewRouter()
	r.Route("/notes", func(r chi.Router) {
		r.Get("/", router.NoteController.GetAll)
		r.Post("/", router.NoteController.Create)
		r.Route("/:noteID", func(r chi.Router) {
			r.Use(router.NoteController.BeforeFilter)
			r.Get("/", router.NoteController.Get)
		})
	})
	return r
}
