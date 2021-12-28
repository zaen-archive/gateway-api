package middleware

import (
	"strings"

	"gateway/configuration"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Cors
func Cors(config *configuration.Configuration) fiber.Handler {
	return cors.New(cors.Config{
		AllowHeaders:     strings.Join(config.Header.AllowHeaders, ", "),
		AllowOrigins:     strings.Join(config.Header.Origins, ", "),
		AllowMethods:     strings.Join(config.Header.Methods, ", "),
		AllowCredentials: config.Header.Credentials,
	})
}
