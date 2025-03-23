package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/rowasjo/tinyvalgo/internal/lib"
	"github.com/rowasjo/tinyvalgo/internal/tinyvalapi"
)

func run(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg, err := env.ParseAs[tinyvalapi.Config]()
	if err != nil {
		return err
	}

	repo := lib.NewDiskRepository(cfg.DataDir)

	app := tinyvalapi.NewApp(repo)

	port := strconv.Itoa(int(cfg.Port))
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: app,
	}

	go func() {
		slog.Info("listening", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error listening and serving", "err", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Error("error shutting down http server", "err", err)
		}
	}()
	wg.Wait()
	return nil
}

func main() {

	ctx := context.Background()
	if err := run(ctx, os.Args); err != nil {
		slog.Error("fatal error", "err", err)
		os.Exit(1)
	}
}
