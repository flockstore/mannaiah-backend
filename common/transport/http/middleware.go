package httptransport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// RequestIDMiddleware attaches a unique X-Request-ID to each request.
func RequestIDMiddleware() fiber.Handler {
	return requestid.New()
}

// CORSMiddleware sets up Cross-Origin Resource Sharing with default options.
func CORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	})
}

// RecoveryMiddleware recovers from panics and returns a 500 error.
func RecoveryMiddleware() fiber.Handler {
	return recover.New()
}
