package httptransport

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

// TestRequestIDMiddleware verifies that RequestIDMiddleware sets the X-Request-ID header.
func TestRequestIDMiddleware(t *testing.T) {

	app := fiber.New()
	app.Use(RequestIDMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		requestID := c.GetRespHeader(fiber.HeaderXRequestID)
		return c.SendString(requestID)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	buf := make([]byte, 64)
	n, _ := resp.Body.Read(buf)
	require.NotEmpty(t, string(buf[:n]))

}

// TestCORSMiddleware verifies that CORSMiddleware adds the Access-Control-Allow-Origin header.
func TestCORSMiddleware(t *testing.T) {
	app := fiber.New()
	app.Use(CORSMiddleware())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Origin", "http://localhost") // Simulates navigator to actually receive header

	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
}

// TestRecoveryMiddleware verifies that RecoveryMiddleware catches panics and responds with 500.
func TestRecoveryMiddleware(t *testing.T) {

	app := fiber.New()
	app.Use(RecoveryMiddleware())
	app.Get("/", func(c *fiber.Ctx) error {
		panic("something went wrong")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 500, resp.StatusCode)

}
