package ihttp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"multiApp/internal/tutor/config"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server   *http.Server
	httpConf config.HTTPConfig
}

func New(conf config.HTTPConfig) *Server {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	}
	return &Server{
		server:   srv,
		httpConf: conf,
	}
}

func (s *Server) RegisterHandler(r *chi.Mux) {
	s.server.Handler = r
}

func (s *Server) StartAndListen() error {
	slog.Info("HTTP server is starting...")
	if err := s.server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed start server on port %s: %w", ":3000", err)
		}
	}

	return nil
}

func (s *Server) Shutdown() error {
	timeoutContext, timeoutCancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(s.httpConf.TimeoutOnStop),
	)
	defer timeoutCancel()
	if err := s.server.Shutdown(timeoutContext); err != nil {
		return fmt.Errorf("failed gracefull shutdown server: %w", err)
	}

	slog.Info("server was gracefully shutdown")

	return nil
}
