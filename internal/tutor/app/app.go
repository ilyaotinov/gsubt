package app

import (
	"context"
	"fmt"
	"log/slog"
	"multiApp/internal/tutor/config"
	"multiApp/internal/tutor/handler/ihttp"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"golang.org/x/sync/errgroup"
)

type App struct {
	Container Container
	Cfg       config.Config
}

func New() (*App, error) {
	a := &App{}
	var err error
	a.Cfg, err = config.New()
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

	serv := ihttp.New(a.Cfg.HTTPConfig)

	r := ihttp.NewRouter()
	serv.RegisterHandler(r)

	g := &errgroup.Group{}

	g.Go(func() error {
		if err := serv.StartAndListen(); err != nil {
			return fmt.Errorf("failed start http server: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		if err := serv.Shutdown(); err != nil {
			return fmt.Errorf("failed shutdown server: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		slog.Error(err.Error())
	}
}
