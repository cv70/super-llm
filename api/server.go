package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

	"super-llm/api/chat"
	"super-llm/domain/committee"
)

// Server holds the API server configuration
type Server struct {
	router     *gin.Engine
	committee  *committee.CommitteeDomain
}

// NewServer creates a new API server
func NewServer(committee *committee.CommitteeDomain) *Server {
	// Set gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.New()
	
	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	server := &Server{
		router:     router,
		committee:  committee,
	}
	
	// Register routes
	server.registerRoutes()
	
	return server
}

// registerRoutes registers all API routes
func (s *Server) registerRoutes() {
	// Create chat handler
	chatHandler := chat.NewHandler(s.committee)
	s.router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
        AllowCredentials: true,
        MaxAge:           24 * time.Hour, // 缓存预检结果的时间
    }))

	// API routes
	api := s.router.Group("/v1")
	{
		// Chat completions endpoint
		api.POST("/chat/completions", chatHandler.ChatCompletions)

		// completions endpoint
		api.POST("/completions", chatHandler.ChatCompletions)
	}
}

// Start starts the HTTP server
func (s *Server) Start(ctx context.Context, port string) error {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: s.router,
	}
	
	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", slog.Any("err", err))
		}
	}()
	
	slog.Info("Server started", slog.Any("port", port))
	
	// Wait for context cancellation
	<-ctx.Done()
	
	// Shutdown server gracefully
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Server shutdown error", slog.Any("err", err))
		return err
	}
	
	slog.Info("Server stopped")
	return nil
}

// GetRouter returns the gin router for testing or other purposes
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}