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

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

type App struct {
	Cfg       Config
	Container Container
}

func New() (*App, error) {
	a := &App{}
	var err error
	a.Cfg, err = newConfig()
	if err != nil {
		return nil, fmt.Errorf("failed init configuration: %w", err)
	}

	pgConnect, err := a.makePgConnection()
	if err != nil {
		return nil, fmt.Errorf("failed perform connection to database: %w", err)
	}

	container := Container{
		pgConnect: pgConnect,
	}
	a.Container = container

	return a, nil
}

func (a *App) makePgConnection() (*sqlx.DB, error) {
	sqlConfig := a.Cfg.SQLConfig
	conn, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s",
			sqlConfig.Username,
			sqlConfig.DatabaseName,
			sqlConfig.Password,
			sqlConfig.GetStringSSLMode(),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed connect to postgresql datbase: %w", err)
	}

	return conn, nil
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
		Addr:    fmt.Sprintf("%s:%d", a.Cfg.HTTPConfig.Host, a.Cfg.HTTPConfig.Port),
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
