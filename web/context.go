package web

import "net/http"

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: writer,
		R: request,
	}
}
