package jnotes

import "net/http"

type Router interface {
	Handler() http.Handler
	GetURLParameter(r *http.Request, key string) string
}
