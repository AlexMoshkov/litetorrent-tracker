package server

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewServer(port string, handler *mux.Router) *Server {
	return &Server{
		&http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
