package httptransport

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flockstore/mannaiah-backend/common/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

// TestServer_Healthz verifies that the healthcheck endpoint returns 200.
func TestServer_Healthz(t *testing.T) {
	log := logger.New("debug", nil)

	srv := New(Options{
		Port:   0, // port 0 means random port; we won't actually bind
		Logger: log,
	})

	// Directly access the app via the public App() method
	app := srv.App()

	req := httptest.NewRequest("GET", "/healthz", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}

// TestServer_ErrorHandler verifies that the error handler returns standardized JSON.
func TestServer_ErrorHandler(t *testing.T) {
	log := logger.New("debug", nil)

	srv := New(Options{
		Port:   0,
		Logger: log,
		Routes: func(r fiber.Router) {
			r.Get("/error", func(c *fiber.Ctx) error {
				return fiber.ErrBadRequest
			})
		},
	})

	app := srv.App()

	req := httptest.NewRequest("GET", "/error", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)

	var body map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	require.Contains(t, body, "error")
	require.Equal(t, "Bad Request", body["error"].(map[string]interface{})["message"])
}

func TestServer_StartAndShutdown(t *testing.T) {
	log := logger.New("debug", nil)

	srv := New(Options{
		Port:   8081,
		Logger: log,
		Routes: func(r fiber.Router) {
			r.Get("/", func(c *fiber.Ctx) error { return c.SendString("ok") })
		},
	})

	ctx, cancel := context.WithCancel(context.Background())

	// Run Start() in a goroutine
	go func() {
		err := srv.Start(ctx)
		require.NoError(t, err)
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Cancel context to trigger shutdown
	cancel()

	// Give time for shutdown to complete
	time.Sleep(100 * time.Millisecond)
}
