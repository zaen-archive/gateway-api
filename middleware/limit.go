package middleware

import (
	"gateway/configuration"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// CreateLimitStrategy
func CreateLimitStrategy(endpoint *configuration.Endpoint) fiber.Handler {
	rateLimiter := endpoint.RateLimiter
	rateDuration := endpoint.RateDuration
	if rateLimiter > 0 && rateDuration == 0 {
		rateDuration = 1
	}

	return limiter.New(
		limiter.Config{
			Max:        rateLimiter,
			Expiration: time.Duration(rateDuration) * time.Second,
		},
	)
}
