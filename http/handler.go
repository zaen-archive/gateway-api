package http

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	. "gitlab.com/zaen/gateway/configuration"
	"gitlab.com/zaen/gateway/utils"
)

func registerHandlers(config *Configuration, endpoint *Endpoint, handlers ...func(*fiber.Ctx) error) {
	if endpoint.Method != "" {
		if !utils.ArrayContain(config.Header.Methods, endpoint.Method) {
			panic(fmt.Sprintf("Endpoint %s have method %s that doesn't exist in allowed methods!.", endpoint.Endpoint, endpoint.Method))
		}

		createRoute(endpoint.Method, endpoint.Endpoint, handlers...)
		return
	}

	for _, val := range config.Header.Methods {
		createRoute(val, endpoint.Endpoint, handlers...)
	}
}

func createRoute(method string, endpoint string, handlers ...func(*fiber.Ctx) error) error {
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
