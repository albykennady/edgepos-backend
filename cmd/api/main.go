package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/albykennady/edgepos-backend/internal/config"
	"github.com/albykennady/edgepos-backend/internal/httpapi"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	handler := httpapi.NewRouter()

	// TODO 1: Log that EdgePOS is starting on port 8080.
	address := ":" + cfg.Port
	slog.Info("Starting EdgePOS API", "port", cfg.Port)

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		// TODO: Start the HTTP server.
		if err := server.ListenAndServe(); err != nil {
			// TODO: Send the returned error into serverErrors.
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		slog.Error("HTTP server failed", "error", err)

	case <-ctx.Done():
		slog.Info("Shutdown signal received")

		// Create a timeout context for graceful draining
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			slog.Error("Server forced to shutdown:", "error", err)

			if err := server.Close(); err != nil {
				slog.Error("Server forced to close:", "error", err)
			}
		}
		slog.Info("HTTP server shutdown completed")

	}
}
