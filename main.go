package main

import (
	"context"
	"log"
	"main/package/handler"
	"net/http"
)

func main() {
	srv := new(Server)
	router := new(handler.Handler)
	err := srv.Start("8080", router.ServeRoutes())
	if err != nil {
		log.Fatalf("Error: %e", err)
	}
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(context context.Context) error {
	return s.httpServer.Shutdown(context)
}
