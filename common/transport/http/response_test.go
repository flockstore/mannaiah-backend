package httptransport

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

// TestWriteSuccess verifies that WriteSuccess returns status 200 with the correct response format.
func TestWriteSuccess(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderXRequestID, "test-id")
		return WriteSuccess(c, map[string]string{"message": "ok"})
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)

	var body SuccessResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	require.Equal(t, "test-id", body.RequestID)
	require.Equal(t, "ok", body.Data.(map[string]interface{})["message"])
}

// TestWriteCreated verifies that WriteCreated returns status 201 with the correct response format.
func TestWriteCreated(t *testing.T) {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderXRequestID, "created-id")
		return WriteCreated(c, map[string]string{"id": "123"})
	})

	req := httptest.NewRequest("POST", "/", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)

	var body SuccessResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	require.Equal(t, "created-id", body.RequestID)
	require.Equal(t, "123", body.Data.(map[string]interface{})["id"])
}

// TestWriteError verifies that WriteError returns the correct error structure and status code.
func TestWriteError(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderXRequestID, "error-id")
		return WriteError(c, 400, "bad request", map[string]string{"field": "name"})
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)

	var body ErrorResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	require.Equal(t, "error-id", body.RequestID)
	require.Equal(t, 400, body.Error.Code)
	require.Equal(t, "bad request", body.Error.Message)
	require.Equal(t, "name", body.Error.Details.(map[string]interface{})["field"])
}
