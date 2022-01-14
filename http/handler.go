package http

import (
	"errors"
	"fmt"
	"strings"

	"gateway/configuration"
	"gateway/utils"

	"github.com/gofiber/fiber/v2"
)

// Registring Handlers to HTTP (fast http)
func registerHandlers(config *configuration.Configuration, endpoint *configuration.Endpoint, handlers ...fiber.Handler) error {
	// Empty Target
	if len(endpoint.Targets) == 0 {
		panic("target route should not be empty")
	}

	// Pre-Calculation for gain performance when there are request
	for i := 0; i < len(endpoint.Targets); i++ {
		target := &endpoint.Targets[i]
		x := strings.HasSuffix(endpoint.Endpoint, "/*")
		y := strings.HasSuffix(target.URLTarget, "/*")

		// Check if wild route or star route exist
		if (x || y) && !(x && y) {
			panic("wild route not set correctly on route " + endpoint.Endpoint)
		}
		target.IsStar = x && y
	}

	// Endpoint Segment
	// Check Route Segments
	endpoint.Segments = strings.Split(endpoint.Endpoint, "/")
	for i, val := range endpoint.Segments {
		if strings.HasPrefix(val, ":") {
			endpoint.ParamsIndex = append(endpoint.ParamsIndex, i)
		}
	}

	// Registring Route
	if endpoint.Method != "" && endpoint.Method != "ALL" {
		if !utils.ArrayContain(config.Header.Methods, endpoint.Method) {
			panic(fmt.Sprintf("Endpoint %s have method %s that doesn't exist in allowed methods!.", endpoint.Endpoint, endpoint.Method))
		}

		return createRoute(endpoint.Method, endpoint.Endpoint, handlers...)
	}

	for _, val := range config.Header.Methods {
		err := createRoute(val, endpoint.Endpoint, handlers...)
		if err != nil {
			return err
		}
	}

	return nil
}

func createRoute(method string, endpoint string, handlers ...fiber.Handler) error {
	switch method {
	case fiber.MethodPut:
		router.Put(endpoint, handlers...)
	case fiber.MethodDelete:
		router.Delete(endpoint, handlers...)
	case fiber.MethodGet:
		router.Get(endpoint, handlers...)
	case fiber.MethodPatch:
		router.Patch(endpoint, handlers...)
	case fiber.MethodHead:
		router.Head(endpoint, handlers...)
	case fiber.MethodOptions:
		router.Options(endpoint, handlers...)
	case fiber.MethodPost:
		router.Post(endpoint, handlers...)
	default:
		return errors.New("Method " + method + " not yet defined!.")
	}

	return nil
}
