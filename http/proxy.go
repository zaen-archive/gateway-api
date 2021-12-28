package http

import (
	"gateway/configuration"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/utils"
)

// createProxyStrategy
func createProxyStrategy(endpoint *configuration.Endpoint) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqUrl := string(c.Request().URI().Path())
		reqQuery := c.Request().URI().QueryString()

		urlTarget := endpoint.URLTarget
		urlDest := endpoint.Endpoint

		// Check Params
		reqSegments := strings.Split(reqUrl, "/")
		for _, val := range endpoint.ParamsIndex {
			oldSegment := endpoint.Segments[val]
			newSegment := reqSegments[val]

			urlTarget = strings.ReplaceAll(urlTarget, oldSegment, newSegment)
			urlDest = strings.Replace(urlDest, oldSegment, newSegment, 1)
		}

		// Check Star
		if endpoint.IsStar {
			urlTarget = utils.TrimRight(urlTarget, '*')

			if len(reqUrl) >= len(urlDest) {
				urlTarget += string(reqUrl[len(urlDest)-1:])
			}
		}

		err := proxy.Do(c, endpoint.HostTarget+urlTarget+"?"+string(reqQuery))
		if err != nil {
			return err
		}

		c.Response().Header.Del(fiber.HeaderServer)

		return nil
	}
}
