package server

import (
	"context"
	"net/http"
	"product-feedback/provider"
	"time"
)

type Server struct {
	handlers   provider.Handler
	httpServer *http.Server
}

func NewServer(port string, h provider.Handler) *Server {
	router := h.InitRoutes()
	svr := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return &Server{handlers: h, httpServer: svr}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
