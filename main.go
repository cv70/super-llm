package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"super-llm/api"
	"super-llm/config"
	"super-llm/domain/committee"
	"syscall"

	"github.com/cv70/pkgo/mistake"
)

func main() {
	// Load configuration
	cfg := config.GetConfig()
	if cfg == nil {
		log.Fatal("Failed to load configuration")
	}
	ctx := context.Background()

	// Create committee domaincommittee
	committee, err := committee.BuildCommitteeDomain(ctx, cfg)
	mistake.Unwrap(err)

	// Create API server
	server := api.NewServer(committee)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Handle OS signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		slog.Info("Shutting down server...")
		cancel()
	}()

	// Start server
	if err := server.Start(ctx, port); err != nil {
		slog.Error("Server error", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("Server stopped")
}
