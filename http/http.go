package http

import (
	"github.com/gofiber/fiber/v2"
	. "gitlab.com/zaen/gateway/configuration"
	"gitlab.com/zaen/gateway/middleware"
)

var router *fiber.App
var config *Configuration

// Config :
func Config(conf *Configuration) {
	config = conf
	router = fiber.New()
}

// Use :
func Use(middleware ...func(*fiber.Ctx) error) {
	for _, val := range middleware {
		router.Use(val)
	}
}

// Run :
func Run() error {

	// Create Statics Web
	for _, val := range config.Statics {
		router.Static(val.Alias, val.Path)
	}

	for i := 0; i < len(config.Endpoints); i++ {
		endpoint := config.Endpoints[i]
		registerHandlers(
			config, &endpoint,
			middleware.CreateJwtStrategy(&endpoint),
			middleware.CreateProxyStrategy(&endpoint),
		)
	}

	return router.Listen("0.0.0.0:" + config.Port)
}
