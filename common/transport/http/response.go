package httptransport

import (
	"github.com/gofiber/fiber/v2"
)

// SuccessResponse is the standard format for successful responses.
type SuccessResponse struct {
	// Data contains the actual payload returned to the client.
	Data any `json:"data"`

	// RequestID is the unique ID assigned to the request, used for tracing.
	RequestID string `json:"requestId"`
}

// ErrorResponse represents a standardized error payload.
type ErrorResponse struct {
	// Error describes the error structure.
	Error ErrorBody `json:"error"`

	// RequestID is the unique ID assigned to the request, used for tracing.
	RequestID string `json:"requestId"`
}

// ErrorBody contains the core details of the error.
type ErrorBody struct {
	// Code is the HTTP status code of the error.
	Code int `json:"code"`

	// Message is a human-readable explanation of the error.
	Message string `json:"message"`

	// Details can contain additional context or be nil.
	Details any `json:"details,omitempty"`
}

// WriteSuccess sends a standardized success response with the provided data.
func WriteSuccess(c *fiber.Ctx, data any) error {
	requestID := c.GetRespHeader(fiber.HeaderXRequestID)
	resp := SuccessResponse{
		Data:      data,
		RequestID: requestID,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

// WriteCreated sends a standardized 201 response for resource creation.
func WriteCreated(c *fiber.Ctx, data any) error {
	requestID := c.GetRespHeader(fiber.HeaderXRequestID)
	resp := SuccessResponse{
		Data:      data,
		RequestID: requestID,
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// WriteError sends a standardized error response with the given status code and message.
func WriteError(c *fiber.Ctx, code int, message string, details any) error {
	requestID := c.GetRespHeader(fiber.HeaderXRequestID)
	errResp := ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
		RequestID: requestID,
	}
	return c.Status(code).JSON(errResp)
}
