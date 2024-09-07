package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	v1 "multiApp/internal/tutor/http/v1"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run() {
	a.configureLogger()
	a.startHTTPServer()
}

func (a *App) configureLogger() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)
}

func (a *App) startHTTPServer() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()
	r := chi.NewRouter()
	r = v1.Handle(r)
	server := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	g := &errgroup.Group{}

	g.Go(func() error {
		slog.Info("HTTP server is starting...")
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				return fmt.Errorf("failed start server on port %s: %w", ":3000", err)
			}
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		timeoutContext, timeoutCancel := context.WithTimeout(context.Background(), time.Second*5)
		defer timeoutCancel()
		if err := server.Shutdown(timeoutContext); err != nil {
			return fmt.Errorf("failed gracefull shutdown server: %w", err)
		}

		slog.Info("server was gracefully shutdown")

		return nil
	})

	if err := g.Wait(); err != nil {
		slog.Error(err.Error())
	}
}
