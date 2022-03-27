package http

import (
	"encoding/json"
	"fmt"
	"gateway/configuration"
	"gateway/middleware"
	"gateway/proxy"

	"github.com/gofiber/fiber/v2"
)

var router *fiber.App
var config *configuration.Configuration

func makeError(err string) []byte {
	errBytes, _ := json.Marshal(
		map[string]interface{}{
			"ok":      false,
			"message": err,
		},
	)

	return errBytes
}

// Handle Error
func errorHandler(c *fiber.Ctx, e error) error {
	return c.Send(makeError(e.Error()))
}

// Config :
func Config(conf *configuration.Configuration) {
	config = conf
	router = fiber.New(
		fiber.Config{
			ErrorHandler: errorHandler,
			ServerHeader: "kano_gateway",
			// Prefork:      true,
		},
	)
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

	// TODO: This is just for testing purpose
	router.Get(
		"/test",
		middleware.CreateLimitStrategy(&configuration.Endpoint{
			RateLimiter:  1,
			RateDuration: 1,
		}),
		func(c *fiber.Ctx) error {
			fmt.Println(c.Hostname())
			fmt.Println(c.Protocol())

			c.WriteString("Hello WOrld")
			return nil
		},
	)

	// Creating Route
	for i := 0; i < len(config.Endpoints); i++ {
		endpoint := config.Endpoints[i]
		err := registerHandlers(
			config, &endpoint,
			middleware.CreateJwtStrategy(&endpoint),
			proxy.CreateProxyStrategy(&endpoint),
		)

		if err != nil {
			panic(err)
		}
	}

	newRoute := fiber.New()
	router.Use(newRoute)

	// certManager := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("dev.local"),
	// 	Cache:      autocert.DirCache("./certs"),
	// 	Email:      "gooner0709@gmail.com",
	// }
	// tlsConf := &tls.Config{
	// 	GetCertificate: certManager.GetCertificate,
	// 	NextProtos: []string{
	// 		"http/1.1", acme.ALPNProto,
	// 	},
	// }

	// ln, err := net.Listen("tcp4", ":443")
	// if err != nil {
	// 	panic(err)
	// }

	// lnTls := tls.NewListener(ln, certManager.TLSConfig())

	return router.Listen("0.0.0.0:" + config.Port)
	// return router.Listener(lnTls)
}
