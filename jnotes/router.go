package jnotes

import "net/http"

type Router interface {
	Handler() http.Handler
}
