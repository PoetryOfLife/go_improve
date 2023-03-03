package web

import "net/http"

type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandleBasedOnMap()
	var root Filter = func(c *Context) {
		handler.ServerHTTP(c)
	}
	for i := len(builders); i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: NewHandleBasedOnMap(),
		root:    root,
	}
}
