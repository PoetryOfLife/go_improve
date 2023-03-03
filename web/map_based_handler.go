package web

import "net/http"

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}

type Handler interface {
	ServerHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) Route(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServerHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("NOT FOUND"))
	}
}

func (h *HandlerBasedOnMap) key(method, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandleBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context), 128),
	}
}
