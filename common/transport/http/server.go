package httptransport

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Server holds the HTTP server and configuration for graceful shutdown.
type Server struct {
	// app is the Fiber application instance.
	app *fiber.App

	// logger is the reusable zap.SugaredLogger instance.
	logger *zap.SugaredLogger

	// port is the port number on which the server will listen.
	port int
}

// App returns the internal Fiber app instance (useful for testing).
func (s *Server) App() *fiber.App {
	return s.app
}

// Options configures the Server.
type Options struct {
	// Port is the port number where the server should listen.
	Port int

	// Logger is the logger instance to use for logging.
	Logger *zap.SugaredLogger

	// Routes is a function to register application-specific routes.
	Routes func(router fiber.Router)
}

// New creates a new Server with the provided options.
func New(opts Options) *Server {

	app := fiber.New(fiber.Config{
		ErrorHandler: defaultErrorHandler,
	})

	registerMiddlewares(app)

	if opts.Routes != nil {
		opts.Routes(app)
	}

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return &Server{
		app:    app,
		logger: opts.Logger,
		port:   opts.Port,
	}
}

// Start runs the HTTP server and handles graceful shutdown on SIGINT/SIGTERM.
func (s *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf(":%d", s.port)
	s.logger.Infof("Starting server on %s", addr)

	// Start server asynchronously
	go func() {
		if err := s.app.Listen(addr); err != nil {
			s.logger.Errorf("Server error: %v", err)
		}
	}()

	// Wait for interrupt or context cancellation
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigChan:
		s.logger.Info("Received shutdown signal, stopping server...")
	case <-ctx.Done():
		s.logger.Info("Context cancelled, stopping server...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.app.ShutdownWithContext(shutdownCtx)
}

// defaultErrorHandler uses the centralized WriteError to send standardized error responses.
func defaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return WriteError(c, code, err.Error(), nil)
}

// registerMiddlewares sets up standard middlewares on the Fiber app.
func registerMiddlewares(app *fiber.App) {
	app.Use(RequestIDMiddleware())
	app.Use(CORSMiddleware())
	app.Use(RecoveryMiddleware())
}
