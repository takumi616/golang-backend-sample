package web

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	Port    string
	Handler http.Handler
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		Port:    port,
		Handler: handler,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// Generate a context with stopping signals
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Create a http listener
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create a http listener", "error", err)
		return err
	}

	server := &http.Server{
		Handler: s.Handler,
	}

	// Start the http server
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			slog.ErrorContext(ctx, "failed to serve http", "error", err)
			return err
		}

		return nil
	})

	// Shutdown the http server
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(shutdownCtx); err != nil {
		slog.ErrorContext(ctx, "failed to shut down the http server", "error", err)
	}

	// Wait for returned value from the goroutine
	return eg.Wait()
}
