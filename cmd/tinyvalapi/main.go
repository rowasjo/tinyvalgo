package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/rowasjo/tinyvalgo/internal/lib"
	"github.com/rowasjo/tinyvalgo/internal/tinyvalapi"
)

func run(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	dataDir := os.Getenv("TINYVAL_DATA_DIR")
	if dataDir == "" {
		return fmt.Errorf("TINYVAL_DATA_DIR environment variable is not set")
	}

	repo := lib.NewDiskRepository(dataDir)

	srv := tinyvalapi.NewServer(repo)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("", "8080"),
		Handler: srv,
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
